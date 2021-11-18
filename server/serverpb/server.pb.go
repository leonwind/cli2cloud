// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: serverpb/server.proto

package serverpb

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

type OutputLine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientID uint32 `protobuf:"varint,1,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	Line     string `protobuf:"bytes,2,opt,name=Line,proto3" json:"Line,omitempty"`
}

func (x *OutputLine) Reset() {
	*x = OutputLine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serverpb_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutputLine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputLine) ProtoMessage() {}

func (x *OutputLine) ProtoReflect() protoreflect.Message {
	mi := &file_serverpb_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutputLine.ProtoReflect.Descriptor instead.
func (*OutputLine) Descriptor() ([]byte, []int) {
	return file_serverpb_server_proto_rawDescGZIP(), []int{0}
}

func (x *OutputLine) GetClientID() uint32 {
	if x != nil {
		return x.ClientID
	}
	return 0
}

func (x *OutputLine) GetLine() string {
	if x != nil {
		return x.Line
	}
	return ""
}

type RegisterMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientID uint32 `protobuf:"varint,1,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
}

func (x *RegisterMessage) Reset() {
	*x = RegisterMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serverpb_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterMessage) ProtoMessage() {}

func (x *RegisterMessage) ProtoReflect() protoreflect.Message {
	mi := &file_serverpb_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterMessage.ProtoReflect.Descriptor instead.
func (*RegisterMessage) Descriptor() ([]byte, []int) {
	return file_serverpb_server_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterMessage) GetClientID() uint32 {
	if x != nil {
		return x.ClientID
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serverpb_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_serverpb_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_serverpb_server_proto_rawDescGZIP(), []int{2}
}

var File_serverpb_server_proto protoreflect.FileDescriptor

var file_serverpb_server_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x65, 0x63, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x70, 0x62, 0x22, 0x3c, 0x0a, 0x0a, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4c, 0x69, 0x6e,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x4c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x69, 0x6e,
	0x65, 0x22, 0x2d, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xc3, 0x01, 0x0a, 0x09, 0x43, 0x6c,
	0x69, 0x32, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x3a, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x72, 0x65, 0x63, 0x73, 0x68, 0x61, 0x72, 0x64, 0x70, 0x62,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x72, 0x65, 0x63, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x16,
	0x2e, 0x72, 0x65, 0x63, 0x73, 0x68, 0x61, 0x72, 0x64, 0x70, 0x62, 0x2e, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x1a, 0x11, 0x2e, 0x72, 0x65, 0x63, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x28, 0x01, 0x12, 0x42, 0x0a, 0x09, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x63, 0x73, 0x68,
	0x61, 0x72, 0x64, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x63, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x70, 0x62, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x30, 0x01, 0x42,
	0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65,
	0x6f, 0x6e, 0x77, 0x69, 0x6e, 0x64, 0x2f, 0x63, 0x6c, 0x69, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_serverpb_server_proto_rawDescOnce sync.Once
	file_serverpb_server_proto_rawDescData = file_serverpb_server_proto_rawDesc
)

func file_serverpb_server_proto_rawDescGZIP() []byte {
	file_serverpb_server_proto_rawDescOnce.Do(func() {
		file_serverpb_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_serverpb_server_proto_rawDescData)
	})
	return file_serverpb_server_proto_rawDescData
}

var file_serverpb_server_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_serverpb_server_proto_goTypes = []interface{}{
	(*OutputLine)(nil),      // 0: recshardpb.OutputLine
	(*RegisterMessage)(nil), // 1: recshardpb.RegisterMessage
	(*Empty)(nil),           // 2: recshardpb.Empty
}
var file_serverpb_server_proto_depIdxs = []int32{
	2, // 0: recshardpb.Cli2Cloud.Register:input_type -> recshardpb.Empty
	0, // 1: recshardpb.Cli2Cloud.Publish:input_type -> recshardpb.OutputLine
	1, // 2: recshardpb.Cli2Cloud.Subscribe:input_type -> recshardpb.RegisterMessage
	1, // 3: recshardpb.Cli2Cloud.Register:output_type -> recshardpb.RegisterMessage
	2, // 4: recshardpb.Cli2Cloud.Publish:output_type -> recshardpb.Empty
	0, // 5: recshardpb.Cli2Cloud.Subscribe:output_type -> recshardpb.OutputLine
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_serverpb_server_proto_init() }
func file_serverpb_server_proto_init() {
	if File_serverpb_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_serverpb_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutputLine); i {
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
		file_serverpb_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterMessage); i {
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
		file_serverpb_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_serverpb_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_serverpb_server_proto_goTypes,
		DependencyIndexes: file_serverpb_server_proto_depIdxs,
		MessageInfos:      file_serverpb_server_proto_msgTypes,
	}.Build()
	File_serverpb_server_proto = out.File
	file_serverpb_server_proto_rawDesc = nil
	file_serverpb_server_proto_goTypes = nil
	file_serverpb_server_proto_depIdxs = nil
}
