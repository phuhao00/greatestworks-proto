// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: messageId.proto

package messageId

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

type MessageId int32

const (
	MessageId_None MessageId = 0
	//gateway--------
	MessageId_CSGatewayLogin      MessageId = 10001
	MessageId_SCGatewayLogin      MessageId = 10002
	MessageId_CSGatewayLogout     MessageId = 10003
	MessageId_SCGatewayLogout     MessageId = 10004
	MessageId_CSGatewayOnlineList MessageId = 10005
	MessageId_SCGatewayOnlineList MessageId = 10006
	MessageId_CSGatewayJoinOnline MessageId = 10007
	MessageId_SCGatewayJoinOnline MessageId = 10008
	MessageId_CSReconnection      MessageId = 10009
	MessageId_SCReconnection      MessageId = 10010
	MessageId_SceneHeartbeat      MessageId = 10011
	//
	MessageId_CSCreatePlayer MessageId = 50001
	MessageId_SCCreatePlayer MessageId = 50002
	MessageId_CSLogin        MessageId = 50003
	MessageId_SCLogin        MessageId = 50004
	//player
	MessageId_CSAddFriend   MessageId = 100001
	MessageId_SCAddFriend   MessageId = 100002
	MessageId_CSDelFriend   MessageId = 100003
	MessageId_SCDelFriend   MessageId = 100004
	MessageId_CSSendChatMsg MessageId = 100005
	MessageId_SCSendChatMsg MessageId = 100006
	//
	MessageId_ClientRequestObjectSync         MessageId = 200010
	MessageId_ServerRequestObjectSync         MessageId = 200011
	MessageId_ServerRequestObjectSyncComplete MessageId = 200012
	MessageId_Instantiate                     MessageId = 200020
	MessageId_Destroy                         MessageId = 200021
	MessageId_DestroyNetworkObjects           MessageId = 200022
	MessageId_SyncTransform                   MessageId = 200030
	//card
	MessageId_CSCardAction MessageId = 300100
	MessageId_SCCardAction MessageId = 300101
)

// Enum value maps for MessageId.
var (
	MessageId_name = map[int32]string{
		0:      "None",
		10001:  "CSGatewayLogin",
		10002:  "SCGatewayLogin",
		10003:  "CSGatewayLogout",
		10004:  "SCGatewayLogout",
		10005:  "CSGatewayOnlineList",
		10006:  "SCGatewayOnlineList",
		10007:  "CSGatewayJoinOnline",
		10008:  "SCGatewayJoinOnline",
		10009:  "CSReconnection",
		10010:  "SCReconnection",
		10011:  "SceneHeartbeat",
		50001:  "CSCreatePlayer",
		50002:  "SCCreatePlayer",
		50003:  "CSLogin",
		50004:  "SCLogin",
		100001: "CSAddFriend",
		100002: "SCAddFriend",
		100003: "CSDelFriend",
		100004: "SCDelFriend",
		100005: "CSSendChatMsg",
		100006: "SCSendChatMsg",
		200010: "ClientRequestObjectSync",
		200011: "ServerRequestObjectSync",
		200012: "ServerRequestObjectSyncComplete",
		200020: "Instantiate",
		200021: "Destroy",
		200022: "DestroyNetworkObjects",
		200030: "SyncTransform",
		300100: "CSCardAction",
		300101: "SCCardAction",
	}
	MessageId_value = map[string]int32{
		"None":                            0,
		"CSGatewayLogin":                  10001,
		"SCGatewayLogin":                  10002,
		"CSGatewayLogout":                 10003,
		"SCGatewayLogout":                 10004,
		"CSGatewayOnlineList":             10005,
		"SCGatewayOnlineList":             10006,
		"CSGatewayJoinOnline":             10007,
		"SCGatewayJoinOnline":             10008,
		"CSReconnection":                  10009,
		"SCReconnection":                  10010,
		"SceneHeartbeat":                  10011,
		"CSCreatePlayer":                  50001,
		"SCCreatePlayer":                  50002,
		"CSLogin":                         50003,
		"SCLogin":                         50004,
		"CSAddFriend":                     100001,
		"SCAddFriend":                     100002,
		"CSDelFriend":                     100003,
		"SCDelFriend":                     100004,
		"CSSendChatMsg":                   100005,
		"SCSendChatMsg":                   100006,
		"ClientRequestObjectSync":         200010,
		"ServerRequestObjectSync":         200011,
		"ServerRequestObjectSyncComplete": 200012,
		"Instantiate":                     200020,
		"Destroy":                         200021,
		"DestroyNetworkObjects":           200022,
		"SyncTransform":                   200030,
		"CSCardAction":                    300100,
		"SCCardAction":                    300101,
	}
)

func (x MessageId) Enum() *MessageId {
	p := new(MessageId)
	*p = x
	return p
}

func (x MessageId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageId) Descriptor() protoreflect.EnumDescriptor {
	return file_messageId_proto_enumTypes[0].Descriptor()
}

func (MessageId) Type() protoreflect.EnumType {
	return &file_messageId_proto_enumTypes[0]
}

func (x MessageId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageId.Descriptor instead.
func (MessageId) EnumDescriptor() ([]byte, []int) {
	return file_messageId_proto_rawDescGZIP(), []int{0}
}

var File_messageId_proto protoreflect.FileDescriptor

var file_messageId_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x2a, 0xb3, 0x05, 0x0a,
	0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f,
	0x6e, 0x65, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0e, 0x43, 0x53, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x91, 0x4e, 0x12, 0x13, 0x0a, 0x0e, 0x53, 0x43, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x92, 0x4e, 0x12, 0x14,
	0x0a, 0x0f, 0x43, 0x53, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x6f, 0x75,
	0x74, 0x10, 0x93, 0x4e, 0x12, 0x14, 0x0a, 0x0f, 0x53, 0x43, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x10, 0x94, 0x4e, 0x12, 0x18, 0x0a, 0x13, 0x43, 0x53,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x10, 0x95, 0x4e, 0x12, 0x18, 0x0a, 0x13, 0x53, 0x43, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x10, 0x96, 0x4e, 0x12, 0x18,
	0x0a, 0x13, 0x43, 0x53, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4a, 0x6f, 0x69, 0x6e, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x97, 0x4e, 0x12, 0x18, 0x0a, 0x13, 0x53, 0x43, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x4a, 0x6f, 0x69, 0x6e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x10,
	0x98, 0x4e, 0x12, 0x13, 0x0a, 0x0e, 0x43, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x10, 0x99, 0x4e, 0x12, 0x13, 0x0a, 0x0e, 0x53, 0x43, 0x52, 0x65, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x9a, 0x4e, 0x12, 0x13, 0x0a, 0x0e,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x10, 0x9b,
	0x4e, 0x12, 0x14, 0x0a, 0x0e, 0x43, 0x53, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x10, 0xd1, 0x86, 0x03, 0x12, 0x14, 0x0a, 0x0e, 0x53, 0x43, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x10, 0xd2, 0x86, 0x03, 0x12, 0x0d, 0x0a,
	0x07, 0x43, 0x53, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0xd3, 0x86, 0x03, 0x12, 0x0d, 0x0a, 0x07,
	0x53, 0x43, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0xd4, 0x86, 0x03, 0x12, 0x11, 0x0a, 0x0b, 0x43,
	0x53, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x10, 0xa1, 0x8d, 0x06, 0x12, 0x11,
	0x0a, 0x0b, 0x53, 0x43, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x10, 0xa2, 0x8d,
	0x06, 0x12, 0x11, 0x0a, 0x0b, 0x43, 0x53, 0x44, 0x65, 0x6c, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x10, 0xa3, 0x8d, 0x06, 0x12, 0x11, 0x0a, 0x0b, 0x53, 0x43, 0x44, 0x65, 0x6c, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x10, 0xa4, 0x8d, 0x06, 0x12, 0x13, 0x0a, 0x0d, 0x43, 0x53, 0x53, 0x65, 0x6e,
	0x64, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x10, 0xa5, 0x8d, 0x06, 0x12, 0x13, 0x0a, 0x0d,
	0x53, 0x43, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x10, 0xa6, 0x8d,
	0x06, 0x12, 0x1d, 0x0a, 0x17, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x10, 0xca, 0x9a, 0x0c,
	0x12, 0x1d, 0x0a, 0x17, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x10, 0xcb, 0x9a, 0x0c, 0x12,
	0x25, 0x0a, 0x1f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x10, 0xcc, 0x9a, 0x0c, 0x12, 0x11, 0x0a, 0x0b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x69, 0x61, 0x74, 0x65, 0x10, 0xd4, 0x9a, 0x0c, 0x12, 0x0d, 0x0a, 0x07, 0x44, 0x65, 0x73,
	0x74, 0x72, 0x6f, 0x79, 0x10, 0xd5, 0x9a, 0x0c, 0x12, 0x1b, 0x0a, 0x15, 0x44, 0x65, 0x73, 0x74,
	0x72, 0x6f, 0x79, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x10, 0xd6, 0x9a, 0x0c, 0x12, 0x13, 0x0a, 0x0d, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x10, 0xde, 0x9a, 0x0c, 0x12, 0x12, 0x0a, 0x0c, 0x43, 0x53,
	0x43, 0x61, 0x72, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0xc4, 0xa8, 0x12, 0x12, 0x12,
	0x0a, 0x0c, 0x53, 0x43, 0x43, 0x61, 0x72, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0xc5,
	0xa8, 0x12, 0x42, 0x4a, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x68, 0x75, 0x68, 0x61, 0x6f, 0x30, 0x30, 0x2f, 0x67, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x64, 0xaa, 0x02, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x47, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messageId_proto_rawDescOnce sync.Once
	file_messageId_proto_rawDescData = file_messageId_proto_rawDesc
)

func file_messageId_proto_rawDescGZIP() []byte {
	file_messageId_proto_rawDescOnce.Do(func() {
		file_messageId_proto_rawDescData = protoimpl.X.CompressGZIP(file_messageId_proto_rawDescData)
	})
	return file_messageId_proto_rawDescData
}

var file_messageId_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_messageId_proto_goTypes = []interface{}{
	(MessageId)(0), // 0: messageId.MessageId
}
var file_messageId_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_messageId_proto_init() }
func file_messageId_proto_init() {
	if File_messageId_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messageId_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messageId_proto_goTypes,
		DependencyIndexes: file_messageId_proto_depIdxs,
		EnumInfos:         file_messageId_proto_enumTypes,
	}.Build()
	File_messageId_proto = out.File
	file_messageId_proto_rawDesc = nil
	file_messageId_proto_goTypes = nil
	file_messageId_proto_depIdxs = nil
}
