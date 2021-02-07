package tailf

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/hpcloud/tail"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/pipeline"
)

const (
	pipelineTimeField = "time"
	statusField       = "status"
)

type tailer struct {
	tf *Tailf

	filename string
	source   string
	tags     map[string]string

	tail *tail.Tail
	pipe *pipeline.Pipeline

	textLine    bytes.Buffer
	tailerOpen  bool
	channelOpen bool
}

func newTailer(tl *Tailf, filename string) *tailer {
	t := tailer{tf: tl, filename: filename, source: tl.Source}

	t.tags = func() map[string]string {
		var m = make(map[string]string)

		for k, v := range tl.Tags {
			m[k] = v
		}

		if _, ok := m["filename"]; !ok {
			m["filename"] = filename
		}
		return m
	}()

	t.tailerOpen = true
	t.channelOpen = true

	return &t
}

func (t *tailer) run() {
	var err error

	t.tail, err = tail.TailFile(t.filename, t.tf.tailerConf)
	if err != nil {
		t.tf.log.Error("failed of build tailer, err:%s", err)
		return
	}
	defer t.tail.Cleanup()

	if t.tf.Pipeline != "" {
		t.pipe, err = pipeline.NewPipelineFromFile(t.tf.Pipeline)
		if err != nil {
			t.tf.log.Error("failed of pipeline, err:%s", err)
			return
		}
	}

	t.receiver()
}

func (t *tailer) receiver() {
	ticker := time.NewTicker(checkFileExistInterval)
	defer ticker.Stop()

	var line *tail.Line

	for {
		line = nil

		select {
		case <-datakit.Exit.Wait():
			t.tf.log.Debugf("Tailing source:%s, file %s is ending", t.source, t.filename)
			return

		case line, t.tailerOpen = <-t.tail.Lines:
			if t.tailerOpen {
				t.tf.log.Debugf("get %d bytes from %s.%s", len(line.Text), t.source, t.filename)
			} else {
				t.channelOpen = false
			}

			t.tf.log.Debugf("get %d bytes from %s.%s", len(line.Text), t.source, t.filename)

		case <-ticker.C:
			_, statErr := os.Lstat(t.filename)
			if os.IsNotExist(statErr) {
				t.tf.log.Warnf("check file %s is not exist", t.filename)
				return
			}
		}

		text, status := t.multiline(line)
		switch status {
		case _return:
			return
		case _continue:
			continue
		case _next:
			//pass
		}

		// 此处不需要输出 error log，已在函数执行过程中就近输出
		// 只是为了控制顺序流
		text, err := t.decode(text)
		if err != nil {
			t.tf.log.Errorf("decode error, %s", err) // only print err
			continue
		}

		data, err := t.pipeline(text)
		if err != nil {
			t.tf.log.Errorf("run pipeline error, %s", err)
			continue
		}

		if err := io.NamedFeed(data, io.Logging, t.source); err != nil {
			t.tf.log.Error(err)
		}
	}
}

type multilineStatus int

const (
	// tail channel 关闭，执行 return
	_return multilineStatus = iota
	// multiline 判断数据为多行，将数据存入缓存，继续读取下一行
	_continue
	// multiline 判断多行数据结束，将缓存中的数据放出，继续执行后续处理
	_next
)

func (t *tailer) multiline(line *tail.Line) (text string, status multilineStatus) {
	if line != nil {
		text = strings.TrimRight(line.Text, "\r")

		if t.tf.multiline.IsEnabled() {
			if text = t.tf.multiline.ProcessLine(text, &t.textLine); text == "" {
				status = _continue
				return
			}
		}
	}

	if line == nil || !t.channelOpen || !t.tailerOpen {
		if text += t.tf.multiline.Flush(&t.textLine); text == "" {
			if !t.channelOpen {
				status = _return
				t.tf.log.Warnf("Tailing %s data channel is closed", t.filename)
				return
			}

			status = _continue
			return
		}
	}

	if line != nil && line.Err != nil {
		t.tf.log.Errorf("Tailing %q: %s", t.filename, line.Err.Error())
		status = _continue
		return
	}

	status = _next
	return
}

func (t *tailer) decode(text string) (str string, err error) {
	return t.tf.decoder.String(text)
}

func (t *tailer) pipeline(text string) (data []byte, err error) {
	var fields = make(map[string]interface{})

	if t.pipe != nil {
		fields, err = t.pipe.Run(text).Result()
		if err != nil {
			return
		}
	} else {
		fields["message"] = text
	}

	// 不使用pipeline功能，也会取time和stauts字段（使用默认值）

	ts, err := t.takeTime(fields)
	if err != nil {
		return
	}

	t.addStatus(fields)

	data, err = io.MakeMetric(t.source, t.tags, fields, ts)
	if err != nil {
		t.tf.log.Error(err)
	}

	return
}

func (t *tailer) takeTime(fields map[string]interface{}) (ts time.Time, err error) {
	// time should be nano-second
	if v, ok := fields[pipelineTimeField]; ok {
		nanots, ok := v.(int64)
		if !ok {
			t.tf.log.Warnf("filed `%s' should be nano-second, but got `%s'",
				pipelineTimeField, reflect.TypeOf(v).String())
			err = fmt.Errorf("invalid filed `%s: %v'", pipelineTimeField, v)
			return
		}

		ts = time.Unix(nanots/int64(time.Second), nanots%int64(time.Second))
		delete(fields, pipelineTimeField)
	} else {
		ts = time.Now()
	}

	return
}

func (t *tailer) addStatus(fields map[string]interface{}) {
	if v, ok := fields[statusField]; ok {
		// "status" type should be string
		if str, ok := v.(string); ok && str != "" {
			return
		}
	}
	fields["status"] = "info"
}
