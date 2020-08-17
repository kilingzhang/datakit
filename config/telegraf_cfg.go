package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/influxdata/toml"
	"github.com/influxdata/toml/ast"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

//用于支持在datakit.conf中加入telegraf的agent配置
type agent struct {
	Interval                   string `toml:interval`
	RoundInterval              bool   `toml:"round_interval"`
	Precision                  string `toml:"precision"`
	CollectionJitter           string `toml:"collection_jitter"`
	FlushInterval              string `toml:"flush_interval"`
	FlushJitter                string `toml:"flush_jitter"`
	MetricBatchSize            int    `toml:"metric_batch_size"`
	MetricBufferLimit          int    `toml:"metric_buffer_limit"`
	FlushBufferWhenFull        bool   `toml:"-"`
	UTC                        bool   `toml:"utc"`
	Debug                      bool   `toml:"debug"`
	Quiet                      bool   `toml:"quiet"`
	LogTarget                  string `toml:"logtarget"`
	Logfile                    string `toml:"logfile"`
	LogfileRotationInterval    string `toml:"logfile_rotation_interval"`
	LogfileRotationMaxSize     string `toml:"logfile_rotation_max_size"`
	LogfileRotationMaxArchives int    `toml:"logfile_rotation_max_archives"`
	OmitHostname               bool   `toml:"omit_hostname"`
}

type fileoutCfg struct {
	OutputFiles string
}

type httpoutCfg struct {
	HTTPServer string
}

type telegrafcfg struct {
	Agent *agent `toml:"agent"`
}

func defaultTelegrafAgentCfg() *agent {
	c := &agent{
		Interval:                   "10s",
		RoundInterval:              true,
		MetricBatchSize:            1000,
		MetricBufferLimit:          100000,
		CollectionJitter:           "0s",
		FlushInterval:              "10s",
		FlushJitter:                "0s",
		Precision:                  "ns",
		Debug:                      false,
		Quiet:                      false,
		LogTarget:                  "file",
		Logfile:                    filepath.Join(datakit.TelegrafDir, "agent.log"),
		LogfileRotationMaxArchives: 5,
		LogfileRotationMaxSize:     "32MB",
		OmitHostname:               true, // do not append host tag
	}
	return c
}

func (c *Config) loadTelegrafInputsConfigs(inputcfgs map[string]*ast.Table, filters []string) (string, error) {

	telegrafCfgFiles := map[string]interface{}{}

	for fp, tbl := range inputcfgs {

		for field, node := range tbl.Fields {
			switch field {
			case "inputs":
				tbl_, ok := node.(*ast.Table)
				if !ok {
					l.Warnf("ignore bad toml node within %s", fp)
				} else {
					for inputName, _ := range tbl_.Fields {
						l.Debugf("check if telegraf input name(%s)?", inputName)

						if _, ok := inputs.TelegrafInputs[inputName]; ok {
							inputs.AddTelegrafInput(inputName, fp)

							l.Infof("enable telegraf input %s, config: %s", inputName, fp)
							telegrafCfgFiles[fp] = nil
						}
					}
				}
			default:
				l.Warnf("ignore bad toml node within %s", fp)
				// pass: all telegraf input should be the format: inputs.xxx
			}
		}
	}

	l.Info("generating telegraf conf...")
	return c.generateTelegrafConfig(telegrafCfgFiles)
}

const (
	fileOutputTemplate = `
[[outputs.file]]
## Files to write to, "stdout" is a specially handled file.
files = ['{{.OutputFiles}}']
`

	httpOutputTemplate = `
[[outputs.http]]
url = "{{.HTTPServer}}"
method = "POST"
data_format = "influx"
## Additional HTTP headers
#[outputs.http.headers]

`

	warning = `
###################################################################################
# Do not edit this file, it was generated and overrided on each datakit restart.
###################################################################################
`
)

func marshalAgentCfg(cfg *agent) (string, error) {
	agdata, err := toml.Marshal(cfg)
	if err != nil {
		return "", err
	}
	return string(agdata), nil
}

func (c *Config) generateTelegrafConfig(files map[string]interface{}) (string, error) {

	agentcfg, err := marshalAgentCfg(c.MainCfg.TelegrafAgentCfg)
	if err != nil {
		l.Errorf("marshal agent faled: %s", err.Error())
		return "", err
	}

	agentcfg = "\n[agent]\n" + agentcfg
	agentcfg += "\n"

	globalTags := "[global_tags]\n"
	for k, v := range c.MainCfg.GlobalTags {
		tag := fmt.Sprintf("%s='%s'\n", k, v)
		globalTags += tag
	}

	fileoutstr := ""
	httpoutstr := ""

	if c.MainCfg.OutputFile != "" {
		fileCfg := fileoutCfg{
			OutputFiles: c.MainCfg.OutputFile,
		}

		tpl := template.New("")
		tpl, err = tpl.Parse(fileOutputTemplate)
		if err != nil {
			l.Errorf("%s", err.Error())
			return "", err
		}

		buf := bytes.NewBuffer([]byte{})
		if err = tpl.Execute(buf, &fileCfg); err != nil {
			l.Errorf("%s", err.Error())
			return "", err
		}
		fileoutstr = string(buf.Bytes())
	}

	if c.MainCfg.DataWay != nil {
		httpCfg := httpoutCfg{
			HTTPServer: fmt.Sprintf("http://%s/telegraf", c.MainCfg.HTTPBind),
		}

		tpl := template.New("")
		tpl, err = tpl.Parse(httpOutputTemplate)
		if err != nil {
			l.Errorf("%s", err.Error())
			return "", err
		}

		buf := bytes.NewBuffer([]byte{})
		if err = tpl.Execute(buf, &httpCfg); err != nil {
			l.Errorf("%s", err.Error())
			return "", err
		}

		httpoutstr = string(buf.Bytes())
	}

	tlegrafConfig := warning + globalTags + agentcfg + fileoutstr + httpoutstr

	parts := []string{}

	for f, _ := range files {

		l.Infof("try merge %s as telegraf config...", f)

		d, err := ioutil.ReadFile(f)
		if err != nil {
			l.Errorf("%s", err.Error())
			continue
		}

		prt, err := c.BuildInputCfg(d)
		if err != nil {
			continue
		}

		l.Debugf("append telegraf config: %s", prt)

		parts = append(parts, prt)
	}

	if len(parts) == 0 {
		return tlegrafConfig, nil
	}

	inputscfgs := strings.Join(parts, "\n")

	// check if @parts include any datakit input
	tbl, err := toml.Parse([]byte(inputscfgs))
	if err != nil {
		l.Error(err)
		return "", err
	}

	for field, node := range tbl.Fields {
		switch field {
		case "inputs":
			tbl_, ok := node.(*ast.Table)
			if !ok {
				l.Warnf("ignore bad toml node: %s", tbl.Source())
			} else {
				for inputName, _ := range tbl_.Fields {

					// NOTE: if telegraf found any unknown inputs, telegraf will exit,
					// so if any xxx.conf with datakit input and telegraf input mixed, telegraf will exit
					if _, ok := inputs.Inputs[inputName]; ok {
						l.Errorf("found datakit input `%s' within merged telegraf conf:\n%s", inputName, tbl.Source())
						l.Warnf("disable all telegraf inputs")
						for _, v := range TelegrafInputs {
							v.enabled = false
						}
						return "", fmt.Errorf("invalid datakit config")
					}
				}
			}
		default:
			l.Warn("invalid inputs format, ignored")
		}
	}

	tlegrafConfig += inputscfgs

	return tlegrafConfig, err
}

func (c *Config) BuildInputCfg(d []byte) (string, error) {

	var err error

	t := template.New("")
	t, err = t.Parse(string(d))
	if err != nil {
		l.Errorf("%s", err.Error())
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, c.MainCfg); err != nil {
		l.Errorf("%s", err.Error())
		return "", err
	}

	return buf.String(), nil
}
