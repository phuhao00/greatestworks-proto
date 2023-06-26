// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: gm.proto

package gm

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

type CSPlayerGMCmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op     uint32   `protobuf:"varint,1,opt,name=op,proto3" json:"op,omitempty"`
	Params []uint32 `protobuf:"varint,2,rep,packed,name=params,proto3" json:"params,omitempty"`
}

func (x *CSPlayerGMCmd) Reset() {
	*x = CSPlayerGMCmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSPlayerGMCmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSPlayerGMCmd) ProtoMessage() {}

func (x *CSPlayerGMCmd) ProtoReflect() protoreflect.Message {
	mi := &file_gm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSPlayerGMCmd.ProtoReflect.Descriptor instead.
func (*CSPlayerGMCmd) Descriptor() ([]byte, []int) {
	return file_gm_proto_rawDescGZIP(), []int{0}
}

func (x *CSPlayerGMCmd) GetOp() uint32 {
	if x != nil {
		return x.Op
	}
	return 0
}

func (x *CSPlayerGMCmd) GetParams() []uint32 {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_gm_proto protoreflect.FileDescriptor

var file_gm_proto_rawDesc = []byte{
	0x0a, 0x08, 0x67, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x67, 0x6d, 0x22, 0x37,
	0x0a, 0x0d, 0x43, 0x53, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x47, 0x4d, 0x43, 0x6d, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x6f, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x68, 0x75, 0x68, 0x61, 0x6f, 0x30, 0x30, 0x2f, 0x67,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x73, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6d, 0x3b, 0x67, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gm_proto_rawDescOnce sync.Once
	file_gm_proto_rawDescData = file_gm_proto_rawDesc
)

func file_gm_proto_rawDescGZIP() []byte {
	file_gm_proto_rawDescOnce.Do(func() {
		file_gm_proto_rawDescData = protoimpl.X.CompressGZIP(file_gm_proto_rawDescData)
	})
	return file_gm_proto_rawDescData
}

var file_gm_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gm_proto_goTypes = []interface{}{
	(*CSPlayerGMCmd)(nil), // 0: gm.CSPlayerGMCmd
}
var file_gm_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gm_proto_init() }
func file_gm_proto_init() {
	if File_gm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSPlayerGMCmd); i {
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
			RawDescriptor: file_gm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gm_proto_goTypes,
		DependencyIndexes: file_gm_proto_depIdxs,
		MessageInfos:      file_gm_proto_msgTypes,
	}.Build()
	File_gm_proto = out.File
	file_gm_proto_rawDesc = nil
	file_gm_proto_goTypes = nil
	file_gm_proto_depIdxs = nil
}
