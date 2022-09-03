// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: attribute.proto

package attribute

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

//daily update attribute
type AttributeDaily int32

const (
	AttributeDaily_NoneDaily AttributeDaily = 0
)

// Enum value maps for AttributeDaily.
var (
	AttributeDaily_name = map[int32]string{
		0: "NoneDaily",
	}
	AttributeDaily_value = map[string]int32{
		"NoneDaily": 0,
	}
)

func (x AttributeDaily) Enum() *AttributeDaily {
	p := new(AttributeDaily)
	*p = x
	return p
}

func (x AttributeDaily) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AttributeDaily) Descriptor() protoreflect.EnumDescriptor {
	return file_attribute_proto_enumTypes[0].Descriptor()
}

func (AttributeDaily) Type() protoreflect.EnumType {
	return &file_attribute_proto_enumTypes[0]
}

func (x AttributeDaily) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AttributeDaily.Descriptor instead.
func (AttributeDaily) EnumDescriptor() ([]byte, []int) {
	return file_attribute_proto_rawDescGZIP(), []int{0}
}

//weekly update attribute
type AttributeWeekly int32

const (
	AttributeWeekly_NoneWeekly AttributeWeekly = 0
)

// Enum value maps for AttributeWeekly.
var (
	AttributeWeekly_name = map[int32]string{
		0: "NoneWeekly",
	}
	AttributeWeekly_value = map[string]int32{
		"NoneWeekly": 0,
	}
)

func (x AttributeWeekly) Enum() *AttributeWeekly {
	p := new(AttributeWeekly)
	*p = x
	return p
}

func (x AttributeWeekly) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AttributeWeekly) Descriptor() protoreflect.EnumDescriptor {
	return file_attribute_proto_enumTypes[1].Descriptor()
}

func (AttributeWeekly) Type() protoreflect.EnumType {
	return &file_attribute_proto_enumTypes[1]
}

func (x AttributeWeekly) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AttributeWeekly.Descriptor instead.
func (AttributeWeekly) EnumDescriptor() ([]byte, []int) {
	return file_attribute_proto_rawDescGZIP(), []int{1}
}

//monthly update attribute
type AttributeMonthly int32

const (
	AttributeMonthly_NoneMonthly AttributeMonthly = 0
)

// Enum value maps for AttributeMonthly.
var (
	AttributeMonthly_name = map[int32]string{
		0: "NoneMonthly",
	}
	AttributeMonthly_value = map[string]int32{
		"NoneMonthly": 0,
	}
)

func (x AttributeMonthly) Enum() *AttributeMonthly {
	p := new(AttributeMonthly)
	*p = x
	return p
}

func (x AttributeMonthly) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AttributeMonthly) Descriptor() protoreflect.EnumDescriptor {
	return file_attribute_proto_enumTypes[2].Descriptor()
}

func (AttributeMonthly) Type() protoreflect.EnumType {
	return &file_attribute_proto_enumTypes[2]
}

func (x AttributeMonthly) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AttributeMonthly.Descriptor instead.
func (AttributeMonthly) EnumDescriptor() ([]byte, []int) {
	return file_attribute_proto_rawDescGZIP(), []int{2}
}

//once update attribute
type AttributeOnce int32

const (
	AttributeOnce_NoneOnce AttributeOnce = 0
)

// Enum value maps for AttributeOnce.
var (
	AttributeOnce_name = map[int32]string{
		0: "NoneOnce",
	}
	AttributeOnce_value = map[string]int32{
		"NoneOnce": 0,
	}
)

func (x AttributeOnce) Enum() *AttributeOnce {
	p := new(AttributeOnce)
	*p = x
	return p
}

func (x AttributeOnce) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AttributeOnce) Descriptor() protoreflect.EnumDescriptor {
	return file_attribute_proto_enumTypes[3].Descriptor()
}

func (AttributeOnce) Type() protoreflect.EnumType {
	return &file_attribute_proto_enumTypes[3]
}

func (x AttributeOnce) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AttributeOnce.Descriptor instead.
func (AttributeOnce) EnumDescriptor() ([]byte, []int) {
	return file_attribute_proto_rawDescGZIP(), []int{3}
}

type TestImport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TestImport) Reset() {
	*x = TestImport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attribute_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestImport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestImport) ProtoMessage() {}

func (x *TestImport) ProtoReflect() protoreflect.Message {
	mi := &file_attribute_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestImport.ProtoReflect.Descriptor instead.
func (*TestImport) Descriptor() ([]byte, []int) {
	return file_attribute_proto_rawDescGZIP(), []int{0}
}

var File_attribute_proto protoreflect.FileDescriptor

var file_attribute_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x22, 0x0c, 0x0a, 0x0a,
	0x54, 0x65, 0x73, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2a, 0x1f, 0x0a, 0x0e, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x12, 0x0d, 0x0a, 0x09,
	0x4e, 0x6f, 0x6e, 0x65, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x10, 0x00, 0x2a, 0x21, 0x0a, 0x0f, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x12, 0x0e,
	0x0a, 0x0a, 0x4e, 0x6f, 0x6e, 0x65, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x10, 0x00, 0x2a, 0x23,
	0x0a, 0x10, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x74, 0x68,
	0x6c, 0x79, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x6f, 0x6e, 0x65, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c,
	0x79, 0x10, 0x00, 0x2a, 0x1d, 0x0a, 0x0d, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x4f, 0x6e, 0x63, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x6f, 0x6e, 0x65, 0x4f, 0x6e, 0x63, 0x65,
	0x10, 0x00, 0x42, 0x4e, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x68, 0x75, 0x68, 0x61, 0x6f, 0x30, 0x30, 0x2f, 0x67, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x3b, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0xaa, 0x02, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x47, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_attribute_proto_rawDescOnce sync.Once
	file_attribute_proto_rawDescData = file_attribute_proto_rawDesc
)

func file_attribute_proto_rawDescGZIP() []byte {
	file_attribute_proto_rawDescOnce.Do(func() {
		file_attribute_proto_rawDescData = protoimpl.X.CompressGZIP(file_attribute_proto_rawDescData)
	})
	return file_attribute_proto_rawDescData
}

var file_attribute_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_attribute_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_attribute_proto_goTypes = []interface{}{
	(AttributeDaily)(0),   // 0: attribute.AttributeDaily
	(AttributeWeekly)(0),  // 1: attribute.AttributeWeekly
	(AttributeMonthly)(0), // 2: attribute.AttributeMonthly
	(AttributeOnce)(0),    // 3: attribute.AttributeOnce
	(*TestImport)(nil),    // 4: attribute.TestImport
}
var file_attribute_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_attribute_proto_init() }
func file_attribute_proto_init() {
	if File_attribute_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_attribute_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestImport); i {
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
			RawDescriptor: file_attribute_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_attribute_proto_goTypes,
		DependencyIndexes: file_attribute_proto_depIdxs,
		EnumInfos:         file_attribute_proto_enumTypes,
		MessageInfos:      file_attribute_proto_msgTypes,
	}.Build()
	File_attribute_proto = out.File
	file_attribute_proto_rawDesc = nil
	file_attribute_proto_goTypes = nil
	file_attribute_proto_depIdxs = nil
}
