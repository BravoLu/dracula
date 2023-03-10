// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: dracula.proto

package dracula

import (
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

type ExportHealthDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg int32 `protobuf:"varint,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ExportHealthDataRequest) Reset() {
	*x = ExportHealthDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dracula_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportHealthDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportHealthDataRequest) ProtoMessage() {}

func (x *ExportHealthDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dracula_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportHealthDataRequest.ProtoReflect.Descriptor instead.
func (*ExportHealthDataRequest) Descriptor() ([]byte, []int) {
	return file_dracula_proto_rawDescGZIP(), []int{0}
}

func (x *ExportHealthDataRequest) GetMsg() int32 {
	if x != nil {
		return x.Msg
	}
	return 0
}

type ExportHealthDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ExportHealthDataResponse) Reset() {
	*x = ExportHealthDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dracula_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportHealthDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportHealthDataResponse) ProtoMessage() {}

func (x *ExportHealthDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dracula_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportHealthDataResponse.ProtoReflect.Descriptor instead.
func (*ExportHealthDataResponse) Descriptor() ([]byte, []int) {
	return file_dracula_proto_rawDescGZIP(), []int{1}
}

func (x *ExportHealthDataResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_dracula_proto protoreflect.FileDescriptor

var file_dracula_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x72, 0x61, 0x63, 0x75, 0x6c, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x72, 0x61, 0x63, 0x75,
	0x6c, 0x61, 0x22, 0x2b, 0x0a, 0x17, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22,
	0x2c, 0x0a, 0x18, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xa3, 0x01,
	0x0a, 0x17, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x48, 0x74,
	0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x87, 0x01, 0x0a, 0x0c, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x3a, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x72, 0x61, 0x63, 0x75, 0x6c, 0x61, 0x2e, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x70,
	0x70, 0x6c, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x2e, 0x64, 0x72, 0x61, 0x63, 0x75, 0x6c, 0x61, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x42, 0x72, 0x61, 0x76, 0x6f, 0x4c, 0x75, 0x2f, 0x64, 0x72, 0x61, 0x63, 0x75, 0x6c,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dracula_proto_rawDescOnce sync.Once
	file_dracula_proto_rawDescData = file_dracula_proto_rawDesc
)

func file_dracula_proto_rawDescGZIP() []byte {
	file_dracula_proto_rawDescOnce.Do(func() {
		file_dracula_proto_rawDescData = protoimpl.X.CompressGZIP(file_dracula_proto_rawDescData)
	})
	return file_dracula_proto_rawDescData
}

var file_dracula_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dracula_proto_goTypes = []interface{}{
	(*ExportHealthDataRequest)(nil),  // 0: grpc.apple_health_monitor.dracula.ExportHealthDataRequest
	(*ExportHealthDataResponse)(nil), // 1: grpc.apple_health_monitor.dracula.ExportHealthDataResponse
}
var file_dracula_proto_depIdxs = []int32{
	0, // 0: grpc.apple_health_monitor.dracula.ExportHealthHttpService.ExportHealth:input_type -> grpc.apple_health_monitor.dracula.ExportHealthDataRequest
	1, // 1: grpc.apple_health_monitor.dracula.ExportHealthHttpService.ExportHealth:output_type -> grpc.apple_health_monitor.dracula.ExportHealthDataResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dracula_proto_init() }
func file_dracula_proto_init() {
	if File_dracula_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dracula_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportHealthDataRequest); i {
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
		file_dracula_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportHealthDataResponse); i {
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
			RawDescriptor: file_dracula_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dracula_proto_goTypes,
		DependencyIndexes: file_dracula_proto_depIdxs,
		MessageInfos:      file_dracula_proto_msgTypes,
	}.Build()
	File_dracula_proto = out.File
	file_dracula_proto_rawDesc = nil
	file_dracula_proto_goTypes = nil
	file_dracula_proto_depIdxs = nil
}
