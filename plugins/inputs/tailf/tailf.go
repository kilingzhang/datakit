package tailf

import (
	"fmt"
	"path/filepath"
	"sync"
	"time"

	"github.com/hpcloud/tail"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
)

type Tailf struct {
	LogFiles          []string          `toml:"logfiles"`
	Ignore            []string          `toml:"ignore"`
	Source            string            `toml:"source"`
	PipelinePath      string            `toml:"pipeline_path"`
	FromBeginning     bool              `toml:"from_beginning"`
	CharacterEncoding string            `toml:"character_encoding"`
	Tags              map[string]string `toml:"tags"`

	MultilineConfig MultilineConfig `toml:"multiline"`

	InputName   string            `toml:"-"`
	CatalogStr  string            `toml:"-"`
	SampleCfg   string            `toml:"-"`
	PipelineCfg map[string]string `toml:"-"`

	multiline       *Multiline
	decoder         decoder
	tailerConf      tail.Config
	runningFileList sync.Map

	wg  sync.WaitGroup
	log *logger.Logger
}

func NewTailf(inputName, catalogStr, sampleCfg string, pipelineCfg map[string]string) *Tailf {
	return &Tailf{
		InputName:       inputName,
		CatalogStr:      catalogStr,
		SampleCfg:       sampleCfg,
		PipelineCfg:     pipelineCfg,
		runningFileList: sync.Map{},
		wg:              sync.WaitGroup{},
		Tags:            make(map[string]string),
	}
}

func (t *Tailf) PipelineConfig() map[string]string {
	return t.PipelineCfg
}

func (t *Tailf) Catalog() string {
	return t.CatalogStr
}

func (t *Tailf) SampleConfig() string {
	return t.SampleCfg
}

func (*Tailf) Test() (result *inputs.TestResult, err error) {
	// 监听文件变更，无法进行测试
	result.Desc = "success"
	return
}

func (t *Tailf) Run() {
	t.log = logger.SLogger(t.InputName)

	if t.loadcfg() {
		return
	}

	t.log.Infof("tailf input started.")

	ticker := time.NewTicker(defaultDruation)
	defer ticker.Stop()

	for {
		select {
		case <-datakit.Exit.Wait():
			t.log.Infof("waiting for all tailers to exit")
			t.wg.Wait()
			t.log.Info("exit")
			return

		case <-ticker.C:
			fileList := t.getFileList(t.LogFiles, t.Ignore)

			for _, file := range fileList {
				t.tailNewFiles(file)
			}

			if t.FromBeginning {
				// disable auto-discovery, ticker was unreachable
				ticker.Stop()
			}
		}
	}
}

func (t *Tailf) loadcfg() bool {
	var err error

	if t.PipelinePath == "" {
		t.PipelinePath = filepath.Join(datakit.PipelineDir, t.Source+".p")
	} else {
		t.PipelinePath = filepath.Join(datakit.PipelineDir, t.PipelinePath)
	}

	if isExist(t.PipelinePath) {
		t.log.Debugf("use pipeline %s", t.PipelinePath)
	} else {
		t.PipelinePath = ""
		t.log.Warn("no pipeline applied")
	}

	for {
		select {
		case <-datakit.Exit.Wait():
			t.log.Info("exit")
			return true
		default:
			// nil
		}

		if t.Source == "" {
			err = fmt.Errorf("tailf source cannot be empty")
			goto label
		}

		if t.decoder, err = NewDecoder(t.CharacterEncoding); err != nil {
			goto label
		}

		if t.multiline, err = t.MultilineConfig.NewMultiline(); err != nil {
			goto label
		}

		if err = checkPipeLine(t.PipelinePath); err != nil {
			goto label
		} else {
			break
		}

	label:
		t.log.Error(err)
		time.Sleep(time.Second)
	}

	var seek *tail.SeekInfo
	if !t.FromBeginning {
		seek = &tail.SeekInfo{
			Whence: 2, // seek is 2
			Offset: 0,
		}
	}

	t.tailerConf = tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  seek,
		MustExist: true,
		Poll:      false, // default watch method is "inotify"
		Pipe:      false,
		Logger:    tail.DiscardingLogger,
	}

	return false
}

func (t *Tailf) tailNewFiles(file string) {
	if _, ok := t.runningFileList.Load(file); ok {
		return
	}

	t.runningFileList.Store(file, nil)

	t.log.Debugf("start tail, %s", file)

	t.wg.Add(1)
	go func() {
		defer t.wg.Done()

		t.tailStart(file)
		t.runningFileList.Delete(file)
		t.log.Debugf("remove file %s from the list", file)
	}()
}

func (t *Tailf) tailStart(filename string) {
	newTailer(t, filename).run()
}
