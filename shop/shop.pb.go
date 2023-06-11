// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: shop.proto

package shop

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

type ShopType int32

const (
	ShopType_Min           ShopType = 0
	ShopType_NpcShop       ShopType = 1
	ShopType_DressShop     ShopType = 2
	ShopType_HouseShop     ShopType = 3
	ShopType_HomeShop      ShopType = 4
	ShopType_AnimalShop    ShopType = 5
	ShopType_PublicMall    ShopType = 6
	ShopType_Activity      ShopType = 7
	ShopType_PackageShop   ShopType = 8 // 礼包
	ShopType_FurnitureShop ShopType = 9 // 家具
	ShopType_Max           ShopType = 100
)

// Enum value maps for ShopType.
var (
	ShopType_name = map[int32]string{
		0:   "Min",
		1:   "NpcShop",
		2:   "DressShop",
		3:   "HouseShop",
		4:   "HomeShop",
		5:   "AnimalShop",
		6:   "PublicMall",
		7:   "Activity",
		8:   "PackageShop",
		9:   "FurnitureShop",
		100: "Max",
	}
	ShopType_value = map[string]int32{
		"Min":           0,
		"NpcShop":       1,
		"DressShop":     2,
		"HouseShop":     3,
		"HomeShop":      4,
		"AnimalShop":    5,
		"PublicMall":    6,
		"Activity":      7,
		"PackageShop":   8,
		"FurnitureShop": 9,
		"Max":           100,
	}
)

func (x ShopType) Enum() *ShopType {
	p := new(ShopType)
	*p = x
	return p
}

func (x ShopType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ShopType) Descriptor() protoreflect.EnumDescriptor {
	return file_shop_proto_enumTypes[0].Descriptor()
}

func (ShopType) Type() protoreflect.EnumType {
	return &file_shop_proto_enumTypes[0]
}

func (x ShopType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ShopType.Descriptor instead.
func (ShopType) EnumDescriptor() ([]byte, []int) {
	return file_shop_proto_rawDescGZIP(), []int{0}
}

type NpcShopData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NpcType    uint32   `protobuf:"varint,1,opt,name=npcType,proto3" json:"npcType,omitempty"`                      // npc类型
	NpcID      uint64   `protobuf:"varint,2,opt,name=npcID,proto3" json:"npcID,omitempty"`                          // npc唯一ID
	ShopType   ShopType `protobuf:"varint,3,opt,name=shopType,proto3,enum=shop.ShopType" json:"shopType,omitempty"` // 商店类型
	PositionID uint32   `protobuf:"varint,4,opt,name=positionID,proto3" json:"positionID,omitempty"`                // 位置ID
	ShopID     []uint32 `protobuf:"varint,5,rep,packed,name=shopID,proto3" json:"shopID,omitempty"`                 // 商品ID 复数
	AppearTM   int64    `protobuf:"varint,6,opt,name=appearTM,proto3" json:"appearTM,omitempty"`                    // 出现日期(当天的起始时间)
}

func (x *NpcShopData) Reset() {
	*x = NpcShopData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NpcShopData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NpcShopData) ProtoMessage() {}

func (x *NpcShopData) ProtoReflect() protoreflect.Message {
	mi := &file_shop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NpcShopData.ProtoReflect.Descriptor instead.
func (*NpcShopData) Descriptor() ([]byte, []int) {
	return file_shop_proto_rawDescGZIP(), []int{0}
}

func (x *NpcShopData) GetNpcType() uint32 {
	if x != nil {
		return x.NpcType
	}
	return 0
}

func (x *NpcShopData) GetNpcID() uint64 {
	if x != nil {
		return x.NpcID
	}
	return 0
}

func (x *NpcShopData) GetShopType() ShopType {
	if x != nil {
		return x.ShopType
	}
	return ShopType_Min
}

func (x *NpcShopData) GetPositionID() uint32 {
	if x != nil {
		return x.PositionID
	}
	return 0
}

func (x *NpcShopData) GetShopID() []uint32 {
	if x != nil {
		return x.ShopID
	}
	return nil
}

func (x *NpcShopData) GetAppearTM() int64 {
	if x != nil {
		return x.AppearTM
	}
	return 0
}

var File_shop_proto protoreflect.FileDescriptor

var file_shop_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x68,
	0x6f, 0x70, 0x22, 0xbd, 0x01, 0x0a, 0x0b, 0x4e, 0x70, 0x63, 0x53, 0x68, 0x6f, 0x70, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x70, 0x63, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x6e, 0x70, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x70, 0x63, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x70, 0x63,
	0x49, 0x44, 0x12, 0x2a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x53, 0x68, 0x6f, 0x70,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x44, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06,
	0x73, 0x68, 0x6f, 0x70, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x65, 0x61, 0x72,
	0x54, 0x4d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x70, 0x70, 0x65, 0x61, 0x72,
	0x54, 0x4d, 0x2a, 0xa7, 0x01, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x07, 0x0a, 0x03, 0x4d, 0x69, 0x6e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x70, 0x63, 0x53,
	0x68, 0x6f, 0x70, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x72, 0x65, 0x73, 0x73, 0x53, 0x68,
	0x6f, 0x70, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x68, 0x6f,
	0x70, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x10,
	0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x53, 0x68, 0x6f, 0x70, 0x10,
	0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4d, 0x61, 0x6c, 0x6c, 0x10,
	0x06, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x10, 0x07, 0x12,
	0x0f, 0x0a, 0x0b, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x10, 0x08,
	0x12, 0x11, 0x0a, 0x0d, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x53, 0x68, 0x6f,
	0x70, 0x10, 0x09, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x61, 0x78, 0x10, 0x64, 0x42, 0x33, 0x5a, 0x31,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x68, 0x75, 0x68, 0x61,
	0x6f, 0x30, 0x30, 0x2f, 0x67, 0x72, 0x65, 0x61, 0x74, 0x65, 0x73, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x3b, 0x73, 0x68, 0x6f,
	0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shop_proto_rawDescOnce sync.Once
	file_shop_proto_rawDescData = file_shop_proto_rawDesc
)

func file_shop_proto_rawDescGZIP() []byte {
	file_shop_proto_rawDescOnce.Do(func() {
		file_shop_proto_rawDescData = protoimpl.X.CompressGZIP(file_shop_proto_rawDescData)
	})
	return file_shop_proto_rawDescData
}

var file_shop_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_shop_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_shop_proto_goTypes = []interface{}{
	(ShopType)(0),       // 0: shop.ShopType
	(*NpcShopData)(nil), // 1: shop.NpcShopData
}
var file_shop_proto_depIdxs = []int32{
	0, // 0: shop.NpcShopData.shopType:type_name -> shop.ShopType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_shop_proto_init() }
func file_shop_proto_init() {
	if File_shop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NpcShopData); i {
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
			RawDescriptor: file_shop_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shop_proto_goTypes,
		DependencyIndexes: file_shop_proto_depIdxs,
		EnumInfos:         file_shop_proto_enumTypes,
		MessageInfos:      file_shop_proto_msgTypes,
	}.Build()
	File_shop_proto = out.File
	file_shop_proto_rawDesc = nil
	file_shop_proto_goTypes = nil
	file_shop_proto_depIdxs = nil
}