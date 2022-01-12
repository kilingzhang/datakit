package zipkin

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/trace"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	zipkincore "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/zipkin/corev1"
)

// const (
// 	statusErr    = "error"
// 	sourceZipkin = "zipkin"
// )

type Endpoint struct {
	ServiceName string `json:"serviceName"`
	Ipv4        string `json:"ipv4"`
	Ipv6        string `json:"ipv6,omitempty"`
	Port        int16  `json:"port"`
}

type Annotation struct {
	Timestamp int64     `json:"timestamp"`
	Value     string    `json:"value"`
	Host      *Endpoint `json:"endpoint,omitempty"`
}

type BinaryAnnotation struct {
	Key   string    `json:"key"`
	Value string    `json:"value"`
	Host  *Endpoint `json:"endpoint,omitempty"`
}

type ZipkinSpanV1 struct {
	TraceID           string              `thrift:"traceId,1" db:"traceId" json:"traceId"`
	Name              string              `thrift:"name,3" db:"name" json:"name"`
	ParentID          string              `thrift:"parentId,5" db:"parentId" json:"parentId,omitempty"`
	ID                string              `thrift:"id,4" db:"id" json:"id"`
	Timestamp         int64               `thrift:"timestamp,10" db:"timestamp" json:"timestamp,omitempty"`
	Duration          int64               `thrift:"duration,11" db:"duration" json:"duration,omitempty"`
	Debug             bool                `thrift:"debug,9" db:"debug" json:"debug,omitempty"`
	Annotations       []*Annotation       `thrift:"annotations,6" db:"annotations" json:"annotations"`
	BinaryAnnotations []*BinaryAnnotation `thrift:"binary_annotations,8" db:"binary_annotations" json:"binaryAnnotations"`
}

func zipkinConvThriftToJSON(z *zipkincore.Span) *zipkincore.SpanJsonApater {
	zc := &zipkincore.SpanJsonApater{}
	zc.TraceID = uint64(z.TraceID)
	zc.Name = z.Name
	zc.ID = uint64(z.ID)
	if z.ParentID != nil {
		zc.ParentID = uint64(*z.ParentID)
	}

	for _, ano := range z.Annotations {
		jAno := zipkincore.AnnotationJsonApater{}
		jAno.Timestamp = uint64(ano.Timestamp)
		jAno.Value = ano.Value
		if ano.Host != nil {
			ep := &zipkincore.EndpointJsonApater{}
			ep.ServiceName = ano.Host.ServiceName
			ep.Port = ano.Host.Port
			ep.Ipv6 = append(ep.Ipv6, ano.Host.Ipv6...)

			ipbytes := int2ip(uint32(ano.Host.Ipv4))
			ep.Ipv4 = net.IP(ipbytes)
			jAno.Host = ep
		}
		zc.Annotations = append(zc.Annotations, jAno)
	}

	for _, bno := range z.BinaryAnnotations {
		jBno := zipkincore.BinaryAnnotationJsonApater{}
		jBno.Key = bno.Key
		jBno.Value = append(jBno.Value, bno.Value...)
		jBno.AnnotationType = bno.AnnotationType
		if bno.Host != nil {
			ep := &zipkincore.EndpointJsonApater{}
			ep.ServiceName = bno.Host.ServiceName
			ep.Port = bno.Host.Port
			ep.Ipv6 = append(ep.Ipv6, bno.Host.Ipv6...)

			ipbytes := int2ip(uint32(bno.Host.Ipv4))
			ep.Ipv4 = net.IP(ipbytes)

			jBno.Host = ep
		}
		zc.BinaryAnnotations = append(zc.BinaryAnnotations, jBno)
	}
	zc.Debug = z.Debug
	if z.Timestamp != nil {
		zc.Timestamp = uint64(*z.Timestamp)
	}

	if z.Duration != nil {
		zc.Duration = uint64(*z.Duration)
	}

	if z.TraceIDHigh != nil {
		zc.TraceIDHigh = uint64(*z.TraceIDHigh)
	}

	return zc
}

func unmarshalZipkinThriftV1(octets []byte) ([]*zipkincore.Span, error) {
	buffer := thrift.NewTMemoryBuffer()
	if _, err := buffer.Write(octets); err != nil {
		return nil, err
	}

	transport := thrift.NewTBinaryProtocolTransport(buffer)
	_, size, err := transport.ReadListBegin()
	if err != nil {
		return nil, err
	}

	spans := make([]*zipkincore.Span, 0)
	for i := 0; i < size; i++ {
		zs := &zipkincore.Span{}
		if err = zs.Read(transport); err != nil {
			return nil, err
		}
		spans = append(spans, zs)
	}

	if err = transport.ReadListEnd(); err != nil {
		return nil, err
	}

	return spans, nil
}

func thriftSpansToAdapters(zspans []*zipkincore.Span) ([]*trace.TraceAdapter, error) {
	var (
		adapterGroup       []*trace.TraceAdapter
		spanIDs, parentIDs = getSpanIDsAndParentIDs(zspans)
	)
	for _, span := range zspans {
		if span == nil {
			continue
		}

		tAdapter := &trace.TraceAdapter{
			OperationName: span.Name,
			Source:        inputName,
			SpanID:        fmt.Sprintf("%d", uint64(span.ID)),
			Tags:          zipkinTags,
			TraceID:       fmt.Sprintf("%d", uint64(span.TraceID)),
		}

		spanInfo := &io.SpanInfo{
			Toolkit: inputName,
		}

		if span.Duration != nil {
			tAdapter.Duration = (*span.Duration) * int64(time.Microsecond)
		}

		if span.Timestamp != nil {
			tAdapter.Start = (*span.Timestamp) * int64(time.Microsecond)
		} else {
			tAdapter.Start = getStartTimestamp(span)
		}

		js, err := json.Marshal(zipkinConvThriftToJSON(span))
		if err != nil {
			return nil, err
		}
		tAdapter.Content = string(js)

		if span.ParentID != nil {
			tAdapter.ParentID = fmt.Sprintf("%d", uint64(*span.ParentID))
		}

		for _, ano := range span.Annotations {
			if ano.Host != nil && ano.Host.ServiceName != "" {
				tAdapter.ServiceName = ano.Host.ServiceName
				break
			}
		}
		if tAdapter.ServiceName == "" {
			for _, bano := range span.BinaryAnnotations {
				if bano.Host != nil && bano.Host.ServiceName != "" {
					tAdapter.ServiceName = bano.Host.ServiceName
					break
				}
			}
		}
		spanInfo.Service = tAdapter.ServiceName

		tAdapter.Status = trace.STATUS_OK
		for _, bano := range span.BinaryAnnotations {
			if bano != nil && bano.Key == trace.STATUS_ERR {
				tAdapter.Status = trace.STATUS_ERR
				break
			}
		}

		if tAdapter.Duration == 0 {
			tAdapter.Duration = getDurationThriftAno(span.Annotations)
		}

		adapterGroup = append(adapterGroup, tAdapter)
	}

	return adapterGroup, nil
}

func jsonV1SpansToAdapters(zspans []*ZipkinSpanV1) ([]*trace.TraceAdapter, error) {
	var adapterGroup []*trace.TraceAdapter
	for _, span := range zspans {
		tAdapter := &trace.TraceAdapter{}
		tAdapter.Source = sourceZipkin
		tAdapter.Duration = span.Duration * int64(time.Microsecond)
		tAdapter.Start = span.Timestamp * int64(time.Microsecond)
		if tAdapter.Start == 0 {
			tAdapter.Start = getFirstTimestamp(span)
		}

		js, err := json.Marshal(span)
		if err != nil {
			return nil, err
		}
		tAdapter.Content = string(js)

		tAdapter.OperationName = span.Name
		tAdapter.ParentID = span.ParentID
		tAdapter.TraceID = span.TraceID
		tAdapter.SpanID = span.ID

		for _, ano := range span.Annotations {
			if ano.Host != nil && ano.Host.ServiceName != "" {
				tAdapter.ServiceName = ano.Host.ServiceName
				break
			}
		}

		if tAdapter.ServiceName == "" {
			for _, bno := range span.BinaryAnnotations {
				if bno.Host != nil && bno.Host.ServiceName != "" {
					tAdapter.ServiceName = bno.Host.ServiceName
					break
				}
			}
		}

		tAdapter.Status = trace.STATUS_OK
		for _, bno := range span.BinaryAnnotations {
			if bno != nil && bno.Key == statusErr {
				tAdapter.Status = trace.STATUS_ERR
				break
			}
		}

		if tAdapter.Duration == 0 {
			tAdapter.Duration = getDurationByAno(span.Annotations)
		}

		adapterGroup = append(adapterGroup, tAdapter)
	}

	return adapterGroup, nil
}

func getSpanIDsAndParentIDs(trace []*zipkincore.Span) (map[int64]bool, map[int64]bool) {
	var (
		spanIDs   = make(map[int64]bool)
		parentIDs = make(map[int64]bool)
	)
	for _, span := range trace {
		if span == nil {
			continue
		}
		spanIDs[span.ID] = true
		if span.ParentID != nil && *span.ParentID != 0 {
			parentIDs[*span.ParentID] = true
		}
	}

	return spanIDs, parentIDs
}

func getFirstTimestamp(zs *ZipkinSpanV1) int64 {
	var (
		ts      int64 = 0x7FFFFFFFFFFFFFFF
		isFound bool
	)
	for _, ano := range zs.Annotations {
		if ano.Timestamp == 0 {
			continue
		}
		if ano.Timestamp < ts {
			isFound = true
			ts = ano.Timestamp
		}
	}

	if isFound {
		return ts * 1000
	}

	return time.Now().UnixNano()
}

func getDurationByAno(anos []*Annotation) int64 {
	if len(anos) < 2 {
		return 0
	}

	var (
		startTS int64 = 0x7FFFFFFFFFFFFFFF
		stopTS  int64
	)
	for _, ano := range anos {
		if ano.Timestamp == 0 {
			continue
		}
		if ano.Timestamp > stopTS {
			stopTS = ano.Timestamp
		}

		if ano.Timestamp < startTS {
			startTS = ano.Timestamp
		}
	}
	if stopTS > startTS {
		return (stopTS - startTS) * 1000
	}

	return 0
}

func getStartTimestamp(zs *zipkincore.Span) int64 {
	var (
		ts      int64 = 0x7FFFFFFFFFFFFFFF
		isFound bool
	)
	for _, ano := range zs.Annotations {
		if ano.Timestamp == 0 {
			continue
		}
		if ano.Timestamp < ts {
			isFound = true
			ts = ano.Timestamp
		}
	}

	if isFound {
		return ts * 1000
	}

	return time.Now().UnixNano()
}

func getDurationThriftAno(anos []*zipkincore.Annotation) int64 {
	if len(anos) < 2 {
		return 0
	}

	var (
		start int64 = 0x7FFFFFFFFFFFFFFF
		stop  int64
	)
	for _, ano := range anos {
		if ano.Timestamp == 0 {
			continue
		}

		if ano.Timestamp > stop {
			stop = ano.Timestamp
		}
		if ano.Timestamp < start {
			start = ano.Timestamp
		}
	}

	if stop > start {
		return (stop - start) * int64(time.Microsecond)
	}

	return 0
}

func int2ip(i uint32) []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i)
	return bs
}
