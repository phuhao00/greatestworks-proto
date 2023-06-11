// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: purchase.proto

package purchase

import (
	ErrCode "github.com/phuhao00/greatestworks-proto/ErrCode"
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

type CardAction int32

const (
	CardAction_NOAction      CardAction = 0
	CardAction_DailyReceived CardAction = 1
	CardAction_Buy           CardAction = 2
	CardAction_Renew         CardAction = 3
)

// Enum value maps for CardAction.
var (
	CardAction_name = map[int32]string{
		0: "NOAction",
		1: "DailyReceived",
		2: "Buy",
		3: "Renew",
	}
	CardAction_value = map[string]int32{
		"NOAction":      0,
		"DailyReceived": 1,
		"Buy":           2,
		"Renew":         3,
	}
)

func (x CardAction) Enum() *CardAction {
	p := new(CardAction)
	*p = x
	return p
}

func (x CardAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CardAction) Descriptor() protoreflect.EnumDescriptor {
	return file_purchase_proto_enumTypes[0].Descriptor()
}

func (CardAction) Type() protoreflect.EnumType {
	return &file_purchase_proto_enumTypes[0]
}

func (x CardAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CardAction.Descriptor instead.
func (CardAction) EnumDescriptor() ([]byte, []int) {
	return file_purchase_proto_rawDescGZIP(), []int{0}
}

type CSCardAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardId uint32     `protobuf:"varint,1,opt,name=CardId,proto3" json:"CardId,omitempty"`
	Action CardAction `protobuf:"varint,2,opt,name=action,proto3,enum=purchase.CardAction" json:"action,omitempty"`
}

func (x *CSCardAction) Reset() {
	*x = CSCardAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchase_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSCardAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSCardAction) ProtoMessage() {}

func (x *CSCardAction) ProtoReflect() protoreflect.Message {
	mi := &file_purchase_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSCardAction.ProtoReflect.Descriptor instead.
func (*CSCardAction) Descriptor() ([]byte, []int) {
	return file_purchase_proto_rawDescGZIP(), []int{0}
}

func (x *CSCardAction) GetCardId() uint32 {
	if x != nil {
		return x.CardId
	}
	return 0
}

func (x *CSCardAction) GetAction() CardAction {
	if x != nil {
		return x.Action
	}
	return CardAction_NOAction
}

type SCCardAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardId uint32          `protobuf:"varint,1,opt,name=CardId,proto3" json:"CardId,omitempty"`
	Code   ErrCode.ErrCode `protobuf:"varint,2,opt,name=Code,proto3,enum=ErrCode.ErrCode" json:"Code,omitempty"`
	Action CardAction      `protobuf:"varint,3,opt,name=action,proto3,enum=purchase.CardAction" json:"action,omitempty"`
}

func (x *SCCardAction) Reset() {
	*x = SCCardAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchase_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SCCardAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SCCardAction) ProtoMessage() {}

func (x *SCCardAction) ProtoReflect() protoreflect.Message {
	mi := &file_purchase_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SCCardAction.ProtoReflect.Descriptor instead.
func (*SCCardAction) Descriptor() ([]byte, []int) {
	return file_purchase_proto_rawDescGZIP(), []int{1}
}

func (x *SCCardAction) GetCardId() uint32 {
	if x != nil {
		return x.CardId
	}
	return 0
}

func (x *SCCardAction) GetCode() ErrCode.ErrCode {
	if x != nil {
		return x.Code
	}
	return ErrCode.ErrCode_Success
}

func (x *SCCardAction) GetAction() CardAction {
	if x != nil {
		return x.Action
	}
	return CardAction_NOAction
}

type SCCards struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*CardInfo `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
}

func (x *SCCards) Reset() {
	*x = SCCards{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchase_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SCCards) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SCCards) ProtoMessage() {}

func (x *SCCards) ProtoReflect() protoreflect.Message {
	mi := &file_purchase_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SCCards.ProtoReflect.Descriptor instead.
func (*SCCards) Descriptor() ([]byte, []int) {
	return file_purchase_proto_rawDescGZIP(), []int{2}
}

func (x *SCCards) GetInfos() []*CardInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

type CardInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExpireTime       int64 `protobuf:"varint,1,opt,name=ExpireTime,proto3" json:"ExpireTime,omitempty"`
	CanReceivedTimes int32 `protobuf:"varint,2,opt,name=CanReceivedTimes,proto3" json:"CanReceivedTimes,omitempty"`
}

func (x *CardInfo) Reset() {
	*x = CardInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchase_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardInfo) ProtoMessage() {}

func (x *CardInfo) ProtoReflect() protoreflect.Message {
	mi := &file_purchase_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardInfo.ProtoReflect.Descriptor instead.
func (*CardInfo) Descriptor() ([]byte, []int) {
	return file_purchase_proto_rawDescGZIP(), []int{3}
}

func (x *CardInfo) GetExpireTime() int64 {
	if x != nil {
		return x.ExpireTime
	}
	return 0
}

func (x *CardInfo) GetCanReceivedTimes() int32 {
	if x != nil {
		return x.CanReceivedTimes
	}
	return 0
}

var File_purchase_proto protoreflect.FileDescriptor

var file_purchase_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x1a, 0x0f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0c, 0x43,
	0x53, 0x43, 0x61, 0x72, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x43,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x43, 0x61, 0x72,
	0x64, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x43,
	0x61, 0x72, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x7a, 0x0a, 0x0c, 0x53, 0x43, 0x43, 0x61, 0x72, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x61, 0x72, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x43, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x2e, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x2c, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x33, 0x0a,
	0x07, 0x53, 0x43, 0x43, 0x61, 0x72, 0x64, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x6e, 0x66,
	0x6f, 0x73, 0x22, 0x56, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e,
	0x0a, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a,
	0x0a, 0x10, 0x43, 0x61, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x43, 0x61, 0x6e, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x2a, 0x41, 0x0a, 0x0a, 0x43, 0x61,
	0x72, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x75, 0x79,
	0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x10, 0x03, 0x42, 0x46, 0x5a,
	0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x68, 0x75, 0x68,
	0x61, 0x6f, 0x30, 0x30, 0x2f, 0x67, 0x72, 0x65, 0x61, 0x74, 0x65, 0x73, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73,
	0x65, 0x3b, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0xaa, 0x02, 0x08, 0x70, 0x75, 0x72,
	0x63, 0x68, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_purchase_proto_rawDescOnce sync.Once
	file_purchase_proto_rawDescData = file_purchase_proto_rawDesc
)

func file_purchase_proto_rawDescGZIP() []byte {
	file_purchase_proto_rawDescOnce.Do(func() {
		file_purchase_proto_rawDescData = protoimpl.X.CompressGZIP(file_purchase_proto_rawDescData)
	})
	return file_purchase_proto_rawDescData
}

var file_purchase_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_purchase_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_purchase_proto_goTypes = []interface{}{
	(CardAction)(0),      // 0: purchase.CardAction
	(*CSCardAction)(nil), // 1: purchase.CSCardAction
	(*SCCardAction)(nil), // 2: purchase.SCCardAction
	(*SCCards)(nil),      // 3: purchase.SCCards
	(*CardInfo)(nil),     // 4: purchase.CardInfo
	(ErrCode.ErrCode)(0), // 5: ErrCode.ErrCode
}
var file_purchase_proto_depIdxs = []int32{
	0, // 0: purchase.CSCardAction.action:type_name -> purchase.CardAction
	5, // 1: purchase.SCCardAction.Code:type_name -> ErrCode.ErrCode
	0, // 2: purchase.SCCardAction.action:type_name -> purchase.CardAction
	4, // 3: purchase.SCCards.infos:type_name -> purchase.CardInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_purchase_proto_init() }
func file_purchase_proto_init() {
	if File_purchase_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_purchase_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSCardAction); i {
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
		file_purchase_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SCCardAction); i {
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
		file_purchase_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SCCards); i {
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
		file_purchase_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardInfo); i {
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
			RawDescriptor: file_purchase_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_purchase_proto_goTypes,
		DependencyIndexes: file_purchase_proto_depIdxs,
		EnumInfos:         file_purchase_proto_enumTypes,
		MessageInfos:      file_purchase_proto_msgTypes,
	}.Build()
	File_purchase_proto = out.File
	file_purchase_proto_rawDesc = nil
	file_purchase_proto_goTypes = nil
	file_purchase_proto_depIdxs = nil
}