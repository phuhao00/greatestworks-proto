// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: activity.proto

package activity

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

// 活动类型
type ActivityType int32

const (
	ActivityType_otherNone          ActivityType = 0 // 不是活动
	ActivityType_newCitizen         ActivityType = 1 // 新公民报到
	ActivityType_firstRecharge      ActivityType = 2 // 萌新首冲礼
	ActivityType_fourteenSign       ActivityType = 3 // 14日签到
	ActivityType_celebrityRollDraw  ActivityType = 4 // 庆典大转盘
	ActivityType_dragonBoatFestival ActivityType = 5 // 端午节活动
)

// Enum value maps for ActivityType.
var (
	ActivityType_name = map[int32]string{
		0: "otherNone",
		1: "newCitizen",
		2: "firstRecharge",
		3: "fourteenSign",
		4: "celebrityRollDraw",
		5: "dragonBoatFestival",
	}
	ActivityType_value = map[string]int32{
		"otherNone":          0,
		"newCitizen":         1,
		"firstRecharge":      2,
		"fourteenSign":       3,
		"celebrityRollDraw":  4,
		"dragonBoatFestival": 5,
	}
)

func (x ActivityType) Enum() *ActivityType {
	p := new(ActivityType)
	*p = x
	return p
}

func (x ActivityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActivityType) Descriptor() protoreflect.EnumDescriptor {
	return file_activity_proto_enumTypes[0].Descriptor()
}

func (ActivityType) Type() protoreflect.EnumType {
	return &file_activity_proto_enumTypes[0]
}

func (x ActivityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActivityType.Descriptor instead.
func (ActivityType) EnumDescriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{0}
}

type ActivityItemInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActId uint32       `protobuf:"varint,1,opt,name=actId,proto3" json:"actId,omitempty"` // 活动id
	Info  []*GoodsInfo `protobuf:"bytes,2,rep,name=info,proto3" json:"info,omitempty"`    // 物品信息
}

func (x *ActivityItemInfo) Reset() {
	*x = ActivityItemInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityItemInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityItemInfo) ProtoMessage() {}

func (x *ActivityItemInfo) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityItemInfo.ProtoReflect.Descriptor instead.
func (*ActivityItemInfo) Descriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{0}
}

func (x *ActivityItemInfo) GetActId() uint32 {
	if x != nil {
		return x.ActId
	}
	return 0
}

func (x *ActivityItemInfo) GetInfo() []*GoodsInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type GoodsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId uint32 `protobuf:"varint,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	Num    int32  `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
	Expire int64  `protobuf:"varint,3,opt,name=Expire,proto3" json:"Expire,omitempty"` // 过期时间
}

func (x *GoodsInfo) Reset() {
	*x = GoodsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsInfo) ProtoMessage() {}

func (x *GoodsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsInfo.ProtoReflect.Descriptor instead.
func (*GoodsInfo) Descriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{1}
}

func (x *GoodsInfo) GetItemId() uint32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *GoodsInfo) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *GoodsInfo) GetExpire() int64 {
	if x != nil {
		return x.Expire
	}
	return 0
}

var File_activity_proto protoreflect.FileDescriptor

var file_activity_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x22, 0x51, 0x0a, 0x10, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x4d, 0x0a,
	0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x74,
	0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x49, 0x74, 0x65, 0x6d,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x4e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x2a, 0x81, 0x01, 0x0a,
	0x0c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a,
	0x09, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x6e, 0x65, 0x77, 0x43, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x52, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x10, 0x02, 0x12,
	0x10, 0x0a, 0x0c, 0x66, 0x6f, 0x75, 0x72, 0x74, 0x65, 0x65, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x10,
	0x03, 0x12, 0x15, 0x0a, 0x11, 0x63, 0x65, 0x6c, 0x65, 0x62, 0x72, 0x69, 0x74, 0x79, 0x52, 0x6f,
	0x6c, 0x6c, 0x44, 0x72, 0x61, 0x77, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x64, 0x72, 0x61, 0x67,
	0x6f, 0x6e, 0x42, 0x6f, 0x61, 0x74, 0x46, 0x65, 0x73, 0x74, 0x69, 0x76, 0x61, 0x6c, 0x10, 0x05,
	0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x68, 0x75, 0x68, 0x61, 0x6f, 0x30, 0x30, 0x2f, 0x67, 0x72, 0x65, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x3b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_activity_proto_rawDescOnce sync.Once
	file_activity_proto_rawDescData = file_activity_proto_rawDesc
)

func file_activity_proto_rawDescGZIP() []byte {
	file_activity_proto_rawDescOnce.Do(func() {
		file_activity_proto_rawDescData = protoimpl.X.CompressGZIP(file_activity_proto_rawDescData)
	})
	return file_activity_proto_rawDescData
}

var file_activity_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_activity_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_activity_proto_goTypes = []interface{}{
	(ActivityType)(0),        // 0: activity.ActivityType
	(*ActivityItemInfo)(nil), // 1: activity.ActivityItemInfo
	(*GoodsInfo)(nil),        // 2: activity.GoodsInfo
}
var file_activity_proto_depIdxs = []int32{
	2, // 0: activity.ActivityItemInfo.info:type_name -> activity.GoodsInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_activity_proto_init() }
func file_activity_proto_init() {
	if File_activity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_activity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityItemInfo); i {
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
		file_activity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsInfo); i {
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
			RawDescriptor: file_activity_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_activity_proto_goTypes,
		DependencyIndexes: file_activity_proto_depIdxs,
		EnumInfos:         file_activity_proto_enumTypes,
		MessageInfos:      file_activity_proto_msgTypes,
	}.Build()
	File_activity_proto = out.File
	file_activity_proto_rawDesc = nil
	file_activity_proto_goTypes = nil
	file_activity_proto_depIdxs = nil
}
