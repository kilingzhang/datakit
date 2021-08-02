package inputs

import (
	"fmt"
	"sort"
	"strings"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
)

var (
	MonofontOnTagFieldName = true
	TODO                   = "TODO" // global todo string
)

const (
	Int    = "int"
	Float  = "float"
	String = "string"
	Bool   = "bool"

	// TODO:
	// Prometheus metric types: https://prometheus.io/docs/concepts/metric_types/
	// DataDog metricc types: https://docs.datadoghq.com/developers/metrics/types/?tab=count#metric-types
	Gauge = "gauge"
	Count = "count"
	Rate  = "rate"
	// add more...
)

// units
const (
	UnknownUnit = "-"

	SizeByte  = "Byte"
	SizeIByte = "Byte" // deprecated

	SizeMiB = "MB"

	NCount = "count"

	// time units
	DurationNS     = "nsec"
	DurationUS     = "usec"
	DurationMS     = "msec"
	DurationSecond = "second"
	DurationMinute = "minute"
	DurationHour   = "hour"
	DurationDay    = "day"

	Percent = "%"

	// TODO: add more...
	BytesPerSec    = "B/s"
	RequestsPerSec = "reqs/s"
	Celsius        = "°C"

	Peta = "P" // 10^15
	Tera = "T" // 10^12
	Giga = "G" // 10^9
	Mega = "M" // 10^6
	Kilo = "k" // 10^3
)

type Measurement interface {
	LineProto() (*io.Point, error)
	Info() *MeasurementInfo
}

type FieldInfo struct {
	Type     string // gauge/count/...
	DataType string // int/float/bool/...
	Unit     string
	Desc     string // markdown string
	Disabled bool
}

type TagInfo struct {
	Desc string
}

type MeasurementInfo struct {
	Name   string
	Desc   string
	Fields map[string]interface{}
	Tags   map[string]interface{}
}

func (m *MeasurementInfo) FieldsMarkdownTable() string {
	const tableHeader = `
| 指标 | 描述| 数据类型 | 单位   |
| ---- |---- | :---:    | :----: |`
	const monoRowfmt = "|`%s`|%s|%s|%s|" // 指标/标签列等宽字体展示
	const normalRowfmt = "|%s|%s|%s|%s|"

	rowfmt := monoRowfmt
	if !MonofontOnTagFieldName {
		rowfmt = normalRowfmt
	}

	rows := []string{tableHeader}
	keys := sortMapKey(m.Fields)
	for _, key := range keys { // XXX: f.Type not used
		f, ok := m.Fields[key].(*FieldInfo)
		if !ok {
			continue
		}

		unit := f.Unit
		if unit == "" {
			unit = UnknownUnit
		}

		rows = append(rows, fmt.Sprintf(rowfmt, key, f.Desc, f.DataType, unit))
	}
	return strings.Join(rows, "\n")
}

func (m *MeasurementInfo) TagsMarkdownTable() string {

	if len(m.Tags) == 0 {
		return "暂无"
	}

	tableHeader := `
| 标签名 | 描述    |
|  ----  | --------|`

	rows := []string{tableHeader}
	keys := sortMapKey(m.Tags)
	for _, key := range keys {
		desc := ""
		switch t := m.Tags[key].(type) {
		case *TagInfo:
			desc = t.Desc
		case TagInfo:
			desc = t.Desc
		default:
		}

		if MonofontOnTagFieldName {
			rows = append(rows, fmt.Sprintf("|`%s`|%s|", key, desc))
		} else {
			rows = append(rows, fmt.Sprintf("|%s|%s|", key, desc))
		}
	}
	return strings.Join(rows, "\n")
}

func FeedMeasurement(name, category string, measurements []Measurement, opt *io.Option) error {
	if len(measurements) == 0 {
		return fmt.Errorf("no points")
	}

	var pts []*io.Point
	for _, m := range measurements {
		if pt, err := m.LineProto(); err != nil {
			return err
		} else {
			pts = append(pts, pt)
		}
	}
	return io.Feed(name, category, pts, opt)
}

func NewTagInfo(desc string) *TagInfo {
	return &TagInfo{
		Desc: desc,
	}
}

func sortMapKey(m map[string]interface{}) (res []string) {
	for k := range m {
		res = append(res, k)
	}
	sort.Strings(res)
	return
}
