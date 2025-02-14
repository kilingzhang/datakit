// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package cmds

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit/config"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/pipeline"
)

func runPLFlags() error {
	var txt string

	if *flagPLTxtFile != "" {
		txtBytes, err := ioutil.ReadFile(*flagPLTxtFile)
		if err != nil {
			return fmt.Errorf("ioutil.ReadFile: %w", err)
		}
		txt = string(txtBytes)
		txt = strings.TrimSuffix(txt, "\n")
	}

	if txt == "" {
		if *flagPLTxtData != "" {
			txt = *flagPLTxtData
		}
	}

	if txt == "" {
		return fmt.Errorf("no testing string")
	}

	if strings.HasSuffix(txt, "\n") {
		warnf("[E] txt has suffix EOL\n")
	}

	return pipelineDebugger(debugPipelineName, txt)
}

func pipelineDebugger(plname, txt string) error {
	if err := pipeline.Init(config.Cfg.Pipeline); err != nil {
		return err
	}

	plPath, err := config.GetPipelinePath(plname)
	if err != nil {
		return fmt.Errorf("get pipeline failed: %w", err)
	}
	pl, err := pipeline.NewPipelineFromFile(plPath)
	if err != nil {
		return fmt.Errorf("new pipeline failed: %w", err)
	}

	start := time.Now()
	res, err := pl.Run(txt, "")
	if err != nil {
		return fmt.Errorf("run pipeline failed: %w", err)
	}
	cost := time.Since(start)

	if res == nil || (len(res.Output.Fields) == 0 && len(res.Output.Tags) == 0) {
		errorf("[E] No data extracted from pipeline\n")
		return nil
	}

	result := map[string]interface{}{}
	maxWidth := 0

	if plTime, err := res.GetTime(); err == nil {
		if *flagPLDate {
			result["time"] = plTime
		} else {
			result["time"] = plTime.UnixNano()
		}
	}

	for k, v := range res.Output.Fields {
		if len(k) > maxWidth {
			maxWidth = len(k)
		}
		result[k] = v
	}

	for k, v := range res.Output.Tags {
		result[k+"#"] = v
		if len(k)+1 > maxWidth {
			maxWidth = len(k) + 1
		}
	}

	if res.Output.DataMeasurement != "" {
		result["source#"] = res.Output.DataMeasurement
	}

	if *flagPLTable {
		fmtStr := fmt.Sprintf("%% %ds: %%v", maxWidth)
		lines := []string{}
		for k, v := range result {
			lines = append(lines, fmt.Sprintf(fmtStr, k, v))
		}

		sort.Strings(lines)
		for _, l := range lines {
			fmt.Println(l)
		}
	} else {
		j, err := json.MarshalIndent(result, "", defaultJSONIndent)
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", string(j))
	}

	infof("---------------\n")
	infof("Extracted %d fields, %d tags; drop: %v, cost: %v\n",
		len(res.Output.Fields), len(res.Output.Tags), res.Output.Dropped, cost)

	return nil
}
