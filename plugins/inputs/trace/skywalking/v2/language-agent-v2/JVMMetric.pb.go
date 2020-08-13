//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.2
// source: language-agent-v2/JVMMetric.proto

package language_agent_v2

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	common "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/trace/skywalking/v2/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type JVMMetricCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics           []*common.JVMMetric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	ServiceInstanceId int32               `protobuf:"varint,2,opt,name=serviceInstanceId,proto3" json:"serviceInstanceId,omitempty"`
}

func (x *JVMMetricCollection) Reset() {
	*x = JVMMetricCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_language_agent_v2_JVMMetric_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JVMMetricCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JVMMetricCollection) ProtoMessage() {}

func (x *JVMMetricCollection) ProtoReflect() protoreflect.Message {
	mi := &file_language_agent_v2_JVMMetric_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JVMMetricCollection.ProtoReflect.Descriptor instead.
func (*JVMMetricCollection) Descriptor() ([]byte, []int) {
	return file_language_agent_v2_JVMMetric_proto_rawDescGZIP(), []int{0}
}

func (x *JVMMetricCollection) GetMetrics() []*common.JVMMetric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *JVMMetricCollection) GetServiceInstanceId() int32 {
	if x != nil {
		return x.ServiceInstanceId
	}
	return 0
}

var File_language_agent_v2_JVMMetric_proto protoreflect.FileDescriptor

var file_language_agent_v2_JVMMetric_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2d, 0x76, 0x32, 0x2f, 0x4a, 0x56, 0x4d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x4a, 0x56, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x13, 0x4a, 0x56,
	0x4d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x24, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4a, 0x56, 0x4d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x07,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x32, 0x46, 0x0a, 0x16, 0x4a, 0x56, 0x4d, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2c, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x14, 0x2e, 0x4a, 0x56, 0x4d,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x09, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x00, 0x42, 0xb7, 0x01,
	0x0a, 0x33, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x6b, 0x79,
	0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x32, 0x50, 0x01, 0x5a, 0x61, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e,
	0x6a, 0x69, 0x61, 0x67, 0x6f, 0x75, 0x79, 0x75, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x63, 0x61, 0x72, 0x65, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61,
	0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2d, 0x76, 0x32, 0xaa, 0x02, 0x1a, 0x53, 0x6b, 0x79,
	0x57, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_language_agent_v2_JVMMetric_proto_rawDescOnce sync.Once
	file_language_agent_v2_JVMMetric_proto_rawDescData = file_language_agent_v2_JVMMetric_proto_rawDesc
)

func file_language_agent_v2_JVMMetric_proto_rawDescGZIP() []byte {
	file_language_agent_v2_JVMMetric_proto_rawDescOnce.Do(func() {
		file_language_agent_v2_JVMMetric_proto_rawDescData = protoimpl.X.CompressGZIP(file_language_agent_v2_JVMMetric_proto_rawDescData)
	})
	return file_language_agent_v2_JVMMetric_proto_rawDescData
}

var file_language_agent_v2_JVMMetric_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_language_agent_v2_JVMMetric_proto_goTypes = []interface{}{
	(*JVMMetricCollection)(nil), // 0: JVMMetricCollection
	(*common.JVMMetric)(nil),    // 1: JVMMetric
	(*common.Commands)(nil),     // 2: Commands
}
var file_language_agent_v2_JVMMetric_proto_depIdxs = []int32{
	1, // 0: JVMMetricCollection.metrics:type_name -> JVMMetric
	0, // 1: JVMMetricReportService.collect:input_type -> JVMMetricCollection
	2, // 2: JVMMetricReportService.collect:output_type -> Commands
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_language_agent_v2_JVMMetric_proto_init() }
func file_language_agent_v2_JVMMetric_proto_init() {
	if File_language_agent_v2_JVMMetric_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_language_agent_v2_JVMMetric_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JVMMetricCollection); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_language_agent_v2_JVMMetric_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_language_agent_v2_JVMMetric_proto_goTypes,
		DependencyIndexes: file_language_agent_v2_JVMMetric_proto_depIdxs,
		MessageInfos:      file_language_agent_v2_JVMMetric_proto_msgTypes,
	}.Build()
	File_language_agent_v2_JVMMetric_proto = out.File
	file_language_agent_v2_JVMMetric_proto_rawDesc = nil
	file_language_agent_v2_JVMMetric_proto_goTypes = nil
	file_language_agent_v2_JVMMetric_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// JVMMetricReportServiceClient is the client API for JVMMetricReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JVMMetricReportServiceClient interface {
	Collect(ctx context.Context, in *JVMMetricCollection, opts ...grpc.CallOption) (*common.Commands, error)
}

type jVMMetricReportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJVMMetricReportServiceClient(cc grpc.ClientConnInterface) JVMMetricReportServiceClient {
	return &jVMMetricReportServiceClient{cc}
}

func (c *jVMMetricReportServiceClient) Collect(ctx context.Context, in *JVMMetricCollection, opts ...grpc.CallOption) (*common.Commands, error) {
	out := new(common.Commands)
	err := c.cc.Invoke(ctx, "/JVMMetricReportService/collect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JVMMetricReportServiceServer is the server API for JVMMetricReportService service.
type JVMMetricReportServiceServer interface {
	Collect(context.Context, *JVMMetricCollection) (*common.Commands, error)
}

// UnimplementedJVMMetricReportServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJVMMetricReportServiceServer struct {
}

func (*UnimplementedJVMMetricReportServiceServer) Collect(context.Context, *JVMMetricCollection) (*common.Commands, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collect not implemented")
}

func RegisterJVMMetricReportServiceServer(s *grpc.Server, srv JVMMetricReportServiceServer) {
	s.RegisterService(&_JVMMetricReportService_serviceDesc, srv)
}

func _JVMMetricReportService_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JVMMetricCollection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JVMMetricReportServiceServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JVMMetricReportService/Collect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JVMMetricReportServiceServer).Collect(ctx, req.(*JVMMetricCollection))
	}
	return interceptor(ctx, in, info, handler)
}

var _JVMMetricReportService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "JVMMetricReportService",
	HandlerType: (*JVMMetricReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "collect",
			Handler:    _JVMMetricReportService_Collect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "language-agent-v2/JVMMetric.proto",
}
