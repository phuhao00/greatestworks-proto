// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: complain.proto

package complain

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

type Complain int32

const (
	Complain_Others      Complain = 0   // 其它
	Complain_Abuse       Complain = 1   // 辱骂
	Complain_Plugin      Complain = 2   // 外挂
	Complain_Seduce      Complain = 3   // 诱导
	Complain_AD          Complain = 4   // 广告
	Complain_IllegalMsg  Complain = 5   // 违法消息
	Complain_IllegalName Complain = 6   // 违法名字
	Complain_None        Complain = 100 // 不选
)

// Enum value maps for Complain.
var (
	Complain_name = map[int32]string{
		0:   "Others",
		1:   "Abuse",
		2:   "Plugin",
		3:   "Seduce",
		4:   "AD",
		5:   "IllegalMsg",
		6:   "IllegalName",
		100: "None",
	}
	Complain_value = map[string]int32{
		"Others":      0,
		"Abuse":       1,
		"Plugin":      2,
		"Seduce":      3,
		"AD":          4,
		"IllegalMsg":  5,
		"IllegalName": 6,
		"None":        100,
	}
)

func (x Complain) Enum() *Complain {
	p := new(Complain)
	*p = x
	return p
}

func (x Complain) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Complain) Descriptor() protoreflect.EnumDescriptor {
	return file_complain_proto_enumTypes[0].Descriptor()
}

func (Complain) Type() protoreflect.EnumType {
	return &file_complain_proto_enumTypes[0]
}

func (x Complain) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Complain.Descriptor instead.
func (Complain) EnumDescriptor() ([]byte, []int) {
	return file_complain_proto_rawDescGZIP(), []int{0}
}

var File_complain_proto protoreflect.FileDescriptor

var file_complain_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x2a, 0x6c, 0x0a, 0x08, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x73,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x62, 0x75, 0x73, 0x65, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x65, 0x64,
	0x75, 0x63, 0x65, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x41, 0x44, 0x10, 0x04, 0x12, 0x0e, 0x0a,
	0x0a, 0x49, 0x6c, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x4d, 0x73, 0x67, 0x10, 0x05, 0x12, 0x0f, 0x0a,
	0x0b, 0x49, 0x6c, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x10, 0x06, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x64, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x68, 0x75, 0x68, 0x61, 0x6f, 0x30, 0x30, 0x2f,
	0x67, 0x72, 0x65, 0x61, 0x74, 0x65, 0x73, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x3b, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_complain_proto_rawDescOnce sync.Once
	file_complain_proto_rawDescData = file_complain_proto_rawDesc
)

func file_complain_proto_rawDescGZIP() []byte {
	file_complain_proto_rawDescOnce.Do(func() {
		file_complain_proto_rawDescData = protoimpl.X.CompressGZIP(file_complain_proto_rawDescData)
	})
	return file_complain_proto_rawDescData
}

var file_complain_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_complain_proto_goTypes = []interface{}{
	(Complain)(0), // 0: complain.Complain
}
var file_complain_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_complain_proto_init() }
func file_complain_proto_init() {
	if File_complain_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_complain_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_complain_proto_goTypes,
		DependencyIndexes: file_complain_proto_depIdxs,
		EnumInfos:         file_complain_proto_enumTypes,
	}.Build()
	File_complain_proto = out.File
	file_complain_proto_rawDesc = nil
	file_complain_proto_goTypes = nil
	file_complain_proto_depIdxs = nil
}
