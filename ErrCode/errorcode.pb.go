// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: errorcode.proto

package ErrCode

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

type ErrCode int32

const (
	ErrCode_Success                 ErrCode = 0
	ErrCode_gameplay_max            ErrCode = 100000
	ErrCode_communicate_max         ErrCode = 200000
	ErrCode_purchase_max            ErrCode = 300000
	ErrCode_gateway_verify          ErrCode = 399998
	ErrCode_gateway_repeated_verify ErrCode = 399999
	ErrCode_gateway_max             ErrCode = 400000
)

// Enum value maps for ErrCode.
var (
	ErrCode_name = map[int32]string{
		0:      "Success",
		100000: "gameplay_max",
		200000: "communicate_max",
		300000: "purchase_max",
		399998: "gateway_verify",
		399999: "gateway_repeated_verify",
		400000: "gateway_max",
	}
	ErrCode_value = map[string]int32{
		"Success":                 0,
		"gameplay_max":            100000,
		"communicate_max":         200000,
		"purchase_max":            300000,
		"gateway_verify":          399998,
		"gateway_repeated_verify": 399999,
		"gateway_max":             400000,
	}
)

func (x ErrCode) Enum() *ErrCode {
	p := new(ErrCode)
	*p = x
	return p
}

func (x ErrCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrCode) Descriptor() protoreflect.EnumDescriptor {
	return file_errorcode_proto_enumTypes[0].Descriptor()
}

func (ErrCode) Type() protoreflect.EnumType {
	return &file_errorcode_proto_enumTypes[0]
}

func (x ErrCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrCode.Descriptor instead.
func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return file_errorcode_proto_rawDescGZIP(), []int{0}
}

var File_errorcode_proto protoreflect.FileDescriptor

var file_errorcode_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x2a, 0x9d, 0x01, 0x0a, 0x07, 0x45,
	0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x6d, 0x61, 0x78, 0x10, 0xa0, 0x8d, 0x06, 0x12, 0x15, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x78, 0x10, 0xc0, 0x9a, 0x0c, 0x12, 0x12,
	0x0a, 0x0c, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x78, 0x10, 0xe0,
	0xa7, 0x12, 0x12, 0x14, 0x0a, 0x0e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x10, 0xfe, 0xb4, 0x18, 0x12, 0x1d, 0x0a, 0x17, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x10, 0xff, 0xb4, 0x18, 0x12, 0x11, 0x0a, 0x0b, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x5f, 0x6d, 0x61, 0x78, 0x10, 0x80, 0xb5, 0x18, 0x42, 0x43, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x68, 0x75, 0x68, 0x61, 0x6f, 0x30,
	0x30, 0x2f, 0x67, 0x72, 0x65, 0x61, 0x74, 0x65, 0x73, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x3b, 0x45, 0x72,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0xaa, 0x02, 0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errorcode_proto_rawDescOnce sync.Once
	file_errorcode_proto_rawDescData = file_errorcode_proto_rawDesc
)

func file_errorcode_proto_rawDescGZIP() []byte {
	file_errorcode_proto_rawDescOnce.Do(func() {
		file_errorcode_proto_rawDescData = protoimpl.X.CompressGZIP(file_errorcode_proto_rawDescData)
	})
	return file_errorcode_proto_rawDescData
}

var file_errorcode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errorcode_proto_goTypes = []interface{}{
	(ErrCode)(0), // 0: ErrCode.ErrCode
}
var file_errorcode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errorcode_proto_init() }
func file_errorcode_proto_init() {
	if File_errorcode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_errorcode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errorcode_proto_goTypes,
		DependencyIndexes: file_errorcode_proto_depIdxs,
		EnumInfos:         file_errorcode_proto_enumTypes,
	}.Build()
	File_errorcode_proto = out.File
	file_errorcode_proto_rawDesc = nil
	file_errorcode_proto_goTypes = nil
	file_errorcode_proto_depIdxs = nil
}
