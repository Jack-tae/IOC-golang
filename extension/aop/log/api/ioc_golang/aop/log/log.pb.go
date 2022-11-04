// EDIT IT, change to your package, service and message

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: extension/aop/log/api/ioc_golang/aop/log/log.proto

package log

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AutowireType string     `protobuf:"bytes,1,opt,name=autowireType,proto3" json:"autowireType,omitempty"`
	Sdid         string     `protobuf:"bytes,2,opt,name=sdid,proto3" json:"sdid,omitempty"`
	MethodName   string     `protobuf:"bytes,3,opt,name=methodName,proto3" json:"methodName,omitempty"`
	Machers      []*Matcher `protobuf:"bytes,4,rep,name=machers,proto3" json:"machers,omitempty"`
	Invocation   bool       `protobuf:"varint,5,opt,name=invocation,proto3" json:"invocation,omitempty"`
	Level        int64      `protobuf:"varint,6,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *LogRequest) Reset() {
	*x = LogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extension_aop_log_api_ioc_golang_aop_log_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRequest) ProtoMessage() {}

func (x *LogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_extension_aop_log_api_ioc_golang_aop_log_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRequest.ProtoReflect.Descriptor instead.
func (*LogRequest) Descriptor() ([]byte, []int) {
	return file_extension_aop_log_api_ioc_golang_aop_log_log_proto_rawDescGZIP(), []int{0}
}

func (x *LogRequest) GetAutowireType() string {
	if x != nil {
		return x.AutowireType
	}
	return ""
}

func (x *LogRequest) GetSdid() string {
	if x != nil {
		return x.Sdid
	}
	return ""
}

func (x *LogRequest) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

func (x *LogRequest) GetMachers() []*Matcher {
	if x != nil {
		return x.Machers
	}
	return nil
}

func (x *LogRequest) GetInvocation() bool {
	if x != nil {
		return x.Invocation
	}
	return false
}

func (x *LogRequest) GetLevel() int64 {
	if x != nil {
		return x.Level
	}
	return 0
}

type LogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *LogResponse) Reset() {
	*x = LogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extension_aop_log_api_ioc_golang_aop_log_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogResponse) ProtoMessage() {}

func (x *LogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_extension_aop_log_api_ioc_golang_aop_log_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogResponse.ProtoReflect.Descriptor instead.
func (*LogResponse) Descriptor() ([]byte, []int) {
	return file_extension_aop_log_api_ioc_golang_aop_log_log_proto_rawDescGZIP(), []int{1}
}

func (x *LogResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type Matcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index      int64  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	MatchPath  string `protobuf:"bytes,2,opt,name=matchPath,proto3" json:"matchPath,omitempty"`
	MatchValue string `protobuf:"bytes,3,opt,name=matchValue,proto3" json:"matchValue,omitempty"`
}

func (x *Matcher) Reset() {
	*x = Matcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extension_aop_log_api_ioc_golang_aop_log_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Matcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Matcher) ProtoMessage() {}

func (x *Matcher) ProtoReflect() protoreflect.Message {
	mi := &file_extension_aop_log_api_ioc_golang_aop_log_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Matcher.ProtoReflect.Descriptor instead.
func (*Matcher) Descriptor() ([]byte, []int) {
	return file_extension_aop_log_api_ioc_golang_aop_log_log_proto_rawDescGZIP(), []int{2}
}

func (x *Matcher) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Matcher) GetMatchPath() string {
	if x != nil {
		return x.MatchPath
	}
	return ""
}

func (x *Matcher) GetMatchValue() string {
	if x != nil {
		return x.MatchValue
	}
	return ""
}

var File_extension_aop_log_api_ioc_golang_aop_log_log_proto protoreflect.FileDescriptor

var file_extension_aop_log_api_ioc_golang_aop_log_log_proto_rawDesc = []byte{
	0x0a, 0x32, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x6f, 0x70, 0x2f,
	0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2f, 0x61, 0x6f, 0x70, 0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x61, 0x6f, 0x70, 0x2e, 0x6c, 0x6f, 0x67, 0x22, 0xd1, 0x01, 0x0a, 0x0a, 0x4c, 0x6f, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x77,
	0x69, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61,
	0x75, 0x74, 0x6f, 0x77, 0x69, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x64, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x64, 0x69, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x35, 0x0a, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x61, 0x6f,
	0x70, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x07, 0x6d,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x6e, 0x76, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x27, 0x0a, 0x0b,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x5d, 0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x50,
	0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x32, 0x58, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4a, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x1e, 0x2e, 0x69, 0x6f, 0x63, 0x5f,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x61, 0x6f, 0x70, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6f, 0x63, 0x5f,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x61, 0x6f, 0x70, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x14,
	0x5a, 0x12, 0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x61, 0x6f, 0x70,
	0x2f, 0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_extension_aop_log_api_ioc_golang_aop_log_log_proto_rawDescOnce sync.Once
	file_extension_aop_log_api_ioc_golang_aop_log_log_proto_rawDescData = file_extension_aop_log_api_ioc_golang_aop_log_log_proto_rawDesc
)

func file_extension_aop_log_api_ioc_golang_aop_log_log_proto_rawDescGZIP() []byte {
	file_extension_aop_log_api_ioc_golang_aop_log_log_proto_rawDescOnce.Do(func() {
		file_extension_aop_log_api_ioc_golang_aop_log_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_extension_aop_log_api_ioc_golang_aop_log_log_proto_rawDescData)
	})
	return file_extension_aop_log_api_ioc_golang_aop_log_log_proto_rawDescData
}

var file_extension_aop_log_api_ioc_golang_aop_log_log_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_extension_aop_log_api_ioc_golang_aop_log_log_proto_goTypes = []interface{}{
	(*LogRequest)(nil),  // 0: ioc_golang.aop.log.LogRequest
	(*LogResponse)(nil), // 1: ioc_golang.aop.log.LogResponse
	(*Matcher)(nil),     // 2: ioc_golang.aop.log.Matcher
}
var file_extension_aop_log_api_ioc_golang_aop_log_log_proto_depIdxs = []int32{
	2, // 0: ioc_golang.aop.log.LogRequest.machers:type_name -> ioc_golang.aop.log.Matcher
	0, // 1: ioc_golang.aop.log.LogService.Log:input_type -> ioc_golang.aop.log.LogRequest
	1, // 2: ioc_golang.aop.log.LogService.Log:output_type -> ioc_golang.aop.log.LogResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_extension_aop_log_api_ioc_golang_aop_log_log_proto_init() }
func file_extension_aop_log_api_ioc_golang_aop_log_log_proto_init() {
	if File_extension_aop_log_api_ioc_golang_aop_log_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_extension_aop_log_api_ioc_golang_aop_log_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogRequest); i {
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
		file_extension_aop_log_api_ioc_golang_aop_log_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogResponse); i {
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
		file_extension_aop_log_api_ioc_golang_aop_log_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Matcher); i {
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
			RawDescriptor: file_extension_aop_log_api_ioc_golang_aop_log_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_extension_aop_log_api_ioc_golang_aop_log_log_proto_goTypes,
		DependencyIndexes: file_extension_aop_log_api_ioc_golang_aop_log_log_proto_depIdxs,
		MessageInfos:      file_extension_aop_log_api_ioc_golang_aop_log_log_proto_msgTypes,
	}.Build()
	File_extension_aop_log_api_ioc_golang_aop_log_log_proto = out.File
	file_extension_aop_log_api_ioc_golang_aop_log_log_proto_rawDesc = nil
	file_extension_aop_log_api_ioc_golang_aop_log_log_proto_goTypes = nil
	file_extension_aop_log_api_ioc_golang_aop_log_log_proto_depIdxs = nil
}