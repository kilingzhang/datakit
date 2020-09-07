package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/influxdata/toml"
	"github.com/influxdata/toml/ast"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
	tgi "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/telegraf_inputs"
)

type fileoutCfg struct {
	OutputFiles string
}

type httpoutCfg struct {
	HTTPServer string
}

func loadTelegrafInputsConfigs(c *datakit.Config, inputcfgs map[string]*ast.Table, filters []string) (string, error) {

	// TODO: filters maybe removed
	_ = filters

	telegrafCfgFiles := map[string]interface{}{}

	for fp, tbl := range inputcfgs {

		for field, node := range tbl.Fields {
			switch field {
			case "inputs":
				stbl, ok := node.(*ast.Table)
				if !ok {
					l.Warnf("ignore bad toml node within %s", fp)
				} else {
					for inputName := range stbl.Fields {
						l.Debugf("check if telegraf input name(%s)?", inputName)

						if _, ok := tgi.TelegrafInputs[inputName]; ok {
							l.Infof("find telegraf input %s, config: %s", inputName, fp)
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
	return generateTelegrafConfig(c, telegrafCfgFiles)
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

func marshalAgentCfg(cfg *datakit.TelegrafCfg) (string, error) {
	agdata, err := toml.Marshal(cfg)
	if err != nil {
		return "", err
	}
	return string(agdata), nil
}

func generateTelegrafConfig(c *datakit.Config, files map[string]interface{}) (string, error) {
	telegrafConfig := warning

	globalTags := "[global_tags]\n"
	for k, v := range c.MainCfg.GlobalTags {
		tag := fmt.Sprintf("%s='%s'\n", k, v)
		globalTags += tag
	}

	telegrafConfig += globalTags

	var buf string
	var err error

	if buf, err = marshalAgentCfg(c.MainCfg.TelegrafAgentCfg); err != nil {
		l.Errorf("marshal agent faled: %s", err.Error())
		return "", err
	}

	telegrafConfig += ("\n[agent]\n" + buf + "\n")

	if c.MainCfg.OutputFile != "" {
		if buf, err = applyTelegrafFileOutput(c.MainCfg.OutputFile); err != nil {
			return "", err
		}
		telegrafConfig += buf
	}

	// NOTE: telegraf can also POST to dataway directly, but we redirect the POST
	// to datakit HTTP server to collecting all input's statistics.
	// HTTP server on datakit should be open if any telegraf input enabled.
	if c.MainCfg.DataWay != nil {
		if buf, err = applyTelegrafHTTPOutput(c.MainCfg.HTTPBind); err != nil {
			return "", err
		}
		telegrafConfig += buf
	}

	if buf, err = mergeTelegrafInputsCfgs(files, c.MainCfg); err != nil {
		return "", err
	}

	telegrafConfig += buf

	return telegrafConfig, nil
}

func mergeTelegrafInputsCfgs(files map[string]interface{}, mc *datakit.MainConfig) (string, error) {
	parts := []string{}

	for f := range files {

		l.Infof("try merge %s as telegraf config...", f)

		if fdata, err := ioutil.ReadFile(f); err != nil {
			l.Errorf("%s", err.Error())
			continue
		} else {

			prt, err := BuildInputCfg(fdata, mc)
			if err != nil {
				continue
			}

			if err := addTelegrafCfg(prt, f); err != nil {
				l.Warnf("ignore telegraf input cfg file %s", f)
				continue
			}

			l.Debugf("append telegraf config: %s", prt)

			parts = append(parts, prt)
		}
	}

	merged := strings.Join(parts, "\n")

	// check if merged config parsing ok
	if _, err := toml.Parse([]byte(merged)); err != nil {
		l.Error(err)
		return "", err
	}

	return merged, nil
}

func applyTelegrafFileOutput(fp string) (string, error) {
	fileCfg := fileoutCfg{
		OutputFiles: fp,
	}

	var err error
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

	return buf.String(), nil
}

func applyTelegrafHTTPOutput(server string) (string, error) {
	httpCfg := httpoutCfg{
		HTTPServer: fmt.Sprintf("http://%s/telegraf", server),
	}

	var err error
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

	return buf.String(), nil
}

func addTelegrafCfg(cfgdata, fp string) error {

	tbl, err := toml.Parse([]byte(cfgdata))
	if err != nil {
		l.Warnf("parse failed: %s, ignored, cfgdata:\n%s", err.Error(), cfgdata)
		return err
	}

	inputNames := []string{}

	// test if all inputs.xxx ok
	for field, node := range tbl.Fields {
		switch field {
		case "inputs":
			stbl, ok := node.(*ast.Table)
			if !ok {
				l.Warnf("ignore bad toml node: %s", tbl.Source())
			} else {
				for inputName := range stbl.Fields {

					// NOTE: if telegraf found any unknown inputs(usually it's a datakit input), telegraf
					// will exit, so if any xxx.conf both contains datakit & telegraf inputs, just disable
					// applying xxx.conf on telegraf
					if _, ok := inputs.Inputs[inputName]; ok {
						l.Warnf("found datakit input `%s' while parsing telegraf conf:\n%s", inputName, tbl.Source())
						return fmt.Errorf("mixed datakit inputs %s", inputName)
					}
					inputNames = append(inputNames, inputName)
				}
			}
		default:
			l.Warn("invalid inputs format, ignored")
		}
	}

	for _, name := range inputNames {
		inputs.AddTelegrafInput(name, fp)
	}
	return nil
}

// Telegraf input sample config may contains some template filed from main-config like {{.Hostname}}.
// After importing telegraf source code, most telegraf input sample config comes from telegraf source
// code (no datakit main-config template filed added). But we still keep the settings for some config
// samples that added manually by datakit.
func BuildInputCfg(d []byte, mc *datakit.MainConfig) (string, error) {

	var err error

	t := template.New("")
	t, err = t.Parse(string(d))
	if err != nil {
		l.Errorf("%s", err.Error())
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, mc); err != nil {
		l.Errorf("%s", err.Error())
		return "", err
	}

	return buf.String(), nil
}
