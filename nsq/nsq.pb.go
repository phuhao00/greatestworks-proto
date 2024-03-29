// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: nsq.proto

package nsq

import (
	server_common "github.com/phuhao00/greatestworks-proto/server_common"
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

type NsqCommand int32

const (
	NsqCommand_None                NsqCommand = 0
	NsqCommand_KickOutPlayer       NsqCommand = 1
	NsqCommand_GmAddItem           NsqCommand = 2
	NsqCommand_ServerRestartNotice NsqCommand = 3
	NsqCommand_BanUser             NsqCommand = 4
	NsqCommand_ReloadConfigs       NsqCommand = 5
	NsqCommand_SetMaxOnlineCount   NsqCommand = 6
	NsqCommand_LimitUserAction     NsqCommand = 7
	NsqCommand_SystemClose         NsqCommand = 8
	//friend
	NsqCommand_Friend_Request   NsqCommand = 1001
	NsqCommand_Friend_Accept    NsqCommand = 1002
	NsqCommand_Friend_REFUSE    NsqCommand = 1003 // 拒绝
	NsqCommand_Friend_DELETE    NsqCommand = 1004 // 删除
	NsqCommand_Friend_REFUSE_CD NsqCommand = 1005
	//mail
	NsqCommand_MailSend   NsqCommand = 1101
	NsqCommand_DeleteMail NsqCommand = 1102
)

// Enum value maps for NsqCommand.
var (
	NsqCommand_name = map[int32]string{
		0:    "None",
		1:    "KickOutPlayer",
		2:    "GmAddItem",
		3:    "ServerRestartNotice",
		4:    "BanUser",
		5:    "ReloadConfigs",
		6:    "SetMaxOnlineCount",
		7:    "LimitUserAction",
		8:    "SystemClose",
		1001: "Friend_Request",
		1002: "Friend_Accept",
		1003: "Friend_REFUSE",
		1004: "Friend_DELETE",
		1005: "Friend_REFUSE_CD",
		1101: "MailSend",
		1102: "DeleteMail",
	}
	NsqCommand_value = map[string]int32{
		"None":                0,
		"KickOutPlayer":       1,
		"GmAddItem":           2,
		"ServerRestartNotice": 3,
		"BanUser":             4,
		"ReloadConfigs":       5,
		"SetMaxOnlineCount":   6,
		"LimitUserAction":     7,
		"SystemClose":         8,
		"Friend_Request":      1001,
		"Friend_Accept":       1002,
		"Friend_REFUSE":       1003,
		"Friend_DELETE":       1004,
		"Friend_REFUSE_CD":    1005,
		"MailSend":            1101,
		"DeleteMail":          1102,
	}
)

func (x NsqCommand) Enum() *NsqCommand {
	p := new(NsqCommand)
	*p = x
	return p
}

func (x NsqCommand) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NsqCommand) Descriptor() protoreflect.EnumDescriptor {
	return file_nsq_proto_enumTypes[0].Descriptor()
}

func (NsqCommand) Type() protoreflect.EnumType {
	return &file_nsq_proto_enumTypes[0]
}

func (x NsqCommand) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NsqCommand.Descriptor instead.
func (NsqCommand) EnumDescriptor() ([]byte, []int) {
	return file_nsq_proto_rawDescGZIP(), []int{0}
}

type BanUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID     uint64                   `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	KickReason server_common.KickReason `protobuf:"varint,2,opt,name=kickReason,proto3,enum=ServerCommon.KickReason" json:"kickReason,omitempty"`
	Reason     string                   `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
	SerId      string                   `protobuf:"bytes,4,opt,name=serId,proto3" json:"serId,omitempty"` // 非0 表示此服务器玩家不踢
}

func (x *BanUserInfo) Reset() {
	*x = BanUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nsq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BanUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BanUserInfo) ProtoMessage() {}

func (x *BanUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_nsq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BanUserInfo.ProtoReflect.Descriptor instead.
func (*BanUserInfo) Descriptor() ([]byte, []int) {
	return file_nsq_proto_rawDescGZIP(), []int{0}
}

func (x *BanUserInfo) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *BanUserInfo) GetKickReason() server_common.KickReason {
	if x != nil {
		return x.KickReason
	}
	return server_common.KickReason_UNKNOWN
}

func (x *BanUserInfo) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *BanUserInfo) GetSerId() string {
	if x != nil {
		return x.SerId
	}
	return ""
}

type ComplexMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd  NsqCommand `protobuf:"varint,1,opt,name=cmd,proto3,enum=nsq.NsqCommand" json:"cmd,omitempty"`
	Data []byte     `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Time int64      `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *ComplexMessage) Reset() {
	*x = ComplexMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nsq_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplexMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexMessage) ProtoMessage() {}

func (x *ComplexMessage) ProtoReflect() protoreflect.Message {
	mi := &file_nsq_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplexMessage.ProtoReflect.Descriptor instead.
func (*ComplexMessage) Descriptor() ([]byte, []int) {
	return file_nsq_proto_rawDescGZIP(), []int{1}
}

func (x *ComplexMessage) GetCmd() NsqCommand {
	if x != nil {
		return x.Cmd
	}
	return NsqCommand_None
}

func (x *ComplexMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ComplexMessage) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type FriendMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender   uint64     `protobuf:"varint,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver uint64     `protobuf:"varint,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Cmd      NsqCommand `protobuf:"varint,3,opt,name=cmd,proto3,enum=nsq.NsqCommand" json:"cmd,omitempty"`
}

func (x *FriendMessage) Reset() {
	*x = FriendMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nsq_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendMessage) ProtoMessage() {}

func (x *FriendMessage) ProtoReflect() protoreflect.Message {
	mi := &file_nsq_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendMessage.ProtoReflect.Descriptor instead.
func (*FriendMessage) Descriptor() ([]byte, []int) {
	return file_nsq_proto_rawDescGZIP(), []int{2}
}

func (x *FriendMessage) GetSender() uint64 {
	if x != nil {
		return x.Sender
	}
	return 0
}

func (x *FriendMessage) GetReceiver() uint64 {
	if x != nil {
		return x.Receiver
	}
	return 0
}

func (x *FriendMessage) GetCmd() NsqCommand {
	if x != nil {
		return x.Cmd
	}
	return NsqCommand_None
}

type ReloadConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target   string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	ServerID string `protobuf:"bytes,2,opt,name=serverID,proto3" json:"serverID,omitempty"`
	FileName string `protobuf:"bytes,3,opt,name=fileName,proto3" json:"fileName,omitempty"`
	SendTime int64  `protobuf:"varint,4,opt,name=sendTime,proto3" json:"sendTime,omitempty"`
}

func (x *ReloadConfig) Reset() {
	*x = ReloadConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nsq_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReloadConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReloadConfig) ProtoMessage() {}

func (x *ReloadConfig) ProtoReflect() protoreflect.Message {
	mi := &file_nsq_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReloadConfig.ProtoReflect.Descriptor instead.
func (*ReloadConfig) Descriptor() ([]byte, []int) {
	return file_nsq_proto_rawDescGZIP(), []int{3}
}

func (x *ReloadConfig) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *ReloadConfig) GetServerID() string {
	if x != nil {
		return x.ServerID
	}
	return ""
}

func (x *ReloadConfig) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *ReloadConfig) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

type SetMaxPlayerNum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxPlayerNum int32 `protobuf:"varint,1,opt,name=maxPlayerNum,proto3" json:"maxPlayerNum,omitempty"`
}

func (x *SetMaxPlayerNum) Reset() {
	*x = SetMaxPlayerNum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nsq_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetMaxPlayerNum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMaxPlayerNum) ProtoMessage() {}

func (x *SetMaxPlayerNum) ProtoReflect() protoreflect.Message {
	mi := &file_nsq_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMaxPlayerNum.ProtoReflect.Descriptor instead.
func (*SetMaxPlayerNum) Descriptor() ([]byte, []int) {
	return file_nsq_proto_rawDescGZIP(), []int{4}
}

func (x *SetMaxPlayerNum) GetMaxPlayerNum() int32 {
	if x != nil {
		return x.MaxPlayerNum
	}
	return 0
}

type SetLogLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level uint32 `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *SetLogLevel) Reset() {
	*x = SetLogLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nsq_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetLogLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLogLevel) ProtoMessage() {}

func (x *SetLogLevel) ProtoReflect() protoreflect.Message {
	mi := &file_nsq_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetLogLevel.ProtoReflect.Descriptor instead.
func (*SetLogLevel) Descriptor() ([]byte, []int) {
	return file_nsq_proto_rawDescGZIP(), []int{5}
}

func (x *SetLogLevel) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type CloseGameSystem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SysID uint32 `protobuf:"varint,1,opt,name=sysID,proto3" json:"sysID,omitempty"`
	Close bool   `protobuf:"varint,2,opt,name=close,proto3" json:"close,omitempty"`
}

func (x *CloseGameSystem) Reset() {
	*x = CloseGameSystem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nsq_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseGameSystem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseGameSystem) ProtoMessage() {}

func (x *CloseGameSystem) ProtoReflect() protoreflect.Message {
	mi := &file_nsq_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseGameSystem.ProtoReflect.Descriptor instead.
func (*CloseGameSystem) Descriptor() ([]byte, []int) {
	return file_nsq_proto_rawDescGZIP(), []int{6}
}

func (x *CloseGameSystem) GetSysID() uint32 {
	if x != nil {
		return x.SysID
	}
	return 0
}

func (x *CloseGameSystem) GetClose() bool {
	if x != nil {
		return x.Close
	}
	return false
}

type LimitUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID uint64                     `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Infos  []*server_common.LimitInfo `protobuf:"bytes,2,rep,name=infos,proto3" json:"infos,omitempty"`
}

func (x *LimitUser) Reset() {
	*x = LimitUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nsq_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitUser) ProtoMessage() {}

func (x *LimitUser) ProtoReflect() protoreflect.Message {
	mi := &file_nsq_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitUser.ProtoReflect.Descriptor instead.
func (*LimitUser) Descriptor() ([]byte, []int) {
	return file_nsq_proto_rawDescGZIP(), []int{7}
}

func (x *LimitUser) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *LimitUser) GetInfos() []*server_common.LimitInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

var File_nsq_proto protoreflect.FileDescriptor

var file_nsq_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6e, 0x73, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e, 0x73, 0x71,
	0x1a, 0x12, 0x73, 0x65, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x0b, 0x42, 0x61, 0x6e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x38, 0x0a, 0x0a,
	0x6b, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x18, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x4b, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x0a, 0x6b, 0x69, 0x63, 0x6b,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6e, 0x73, 0x71, 0x2e, 0x4e, 0x73, 0x71, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x22, 0x66, 0x0a, 0x0d, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6e, 0x73, 0x71, 0x2e, 0x4e, 0x73, 0x71, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x22, 0x7a, 0x0a, 0x0c, 0x52, 0x65, 0x6c,
	0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x6d, 0x61, 0x78, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x22, 0x23, 0x0a, 0x0b,
	0x53, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x22, 0x3d, 0x0a, 0x0f, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x79, 0x73, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x79, 0x73, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x22, 0x52, 0x0a, 0x09, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69,
	0x6e, 0x66, 0x6f, 0x73, 0x2a, 0xb6, 0x02, 0x0a, 0x0a, 0x4e, 0x73, 0x71, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x11, 0x0a,
	0x0d, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x10, 0x01,
	0x12, 0x0d, 0x0a, 0x09, 0x47, 0x6d, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x10, 0x02, 0x12,
	0x17, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x61, 0x6e, 0x55,
	0x73, 0x65, 0x72, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x4d,
	0x61, 0x78, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0x06, 0x12,
	0x13, 0x0a, 0x0f, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x10, 0x08, 0x12, 0x13, 0x0a, 0x0e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0xe9, 0x07, 0x12, 0x12, 0x0a, 0x0d, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x5f, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x10, 0xea, 0x07, 0x12, 0x12,
	0x0a, 0x0d, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x52, 0x45, 0x46, 0x55, 0x53, 0x45, 0x10,
	0xeb, 0x07, 0x12, 0x12, 0x0a, 0x0d, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x44, 0x45, 0x4c,
	0x45, 0x54, 0x45, 0x10, 0xec, 0x07, 0x12, 0x15, 0x0a, 0x10, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x5f, 0x52, 0x45, 0x46, 0x55, 0x53, 0x45, 0x5f, 0x43, 0x44, 0x10, 0xed, 0x07, 0x12, 0x0d, 0x0a,
	0x08, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x6e, 0x64, 0x10, 0xcd, 0x08, 0x12, 0x0f, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x10, 0xce, 0x08, 0x42, 0x31, 0x5a,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x68, 0x75, 0x68,
	0x61, 0x6f, 0x30, 0x30, 0x2f, 0x67, 0x72, 0x65, 0x61, 0x74, 0x65, 0x73, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x73, 0x71, 0x3b, 0x6e, 0x73, 0x71,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nsq_proto_rawDescOnce sync.Once
	file_nsq_proto_rawDescData = file_nsq_proto_rawDesc
)

func file_nsq_proto_rawDescGZIP() []byte {
	file_nsq_proto_rawDescOnce.Do(func() {
		file_nsq_proto_rawDescData = protoimpl.X.CompressGZIP(file_nsq_proto_rawDescData)
	})
	return file_nsq_proto_rawDescData
}

var file_nsq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nsq_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_nsq_proto_goTypes = []interface{}{
	(NsqCommand)(0),                 // 0: nsq.NsqCommand
	(*BanUserInfo)(nil),             // 1: nsq.BanUserInfo
	(*ComplexMessage)(nil),          // 2: nsq.ComplexMessage
	(*FriendMessage)(nil),           // 3: nsq.FriendMessage
	(*ReloadConfig)(nil),            // 4: nsq.ReloadConfig
	(*SetMaxPlayerNum)(nil),         // 5: nsq.SetMaxPlayerNum
	(*SetLogLevel)(nil),             // 6: nsq.SetLogLevel
	(*CloseGameSystem)(nil),         // 7: nsq.CloseGameSystem
	(*LimitUser)(nil),               // 8: nsq.LimitUser
	(server_common.KickReason)(0),   // 9: ServerCommon.KickReason
	(*server_common.LimitInfo)(nil), // 10: ServerCommon.LimitInfo
}
var file_nsq_proto_depIdxs = []int32{
	9,  // 0: nsq.BanUserInfo.kickReason:type_name -> ServerCommon.KickReason
	0,  // 1: nsq.ComplexMessage.cmd:type_name -> nsq.NsqCommand
	0,  // 2: nsq.FriendMessage.cmd:type_name -> nsq.NsqCommand
	10, // 3: nsq.LimitUser.infos:type_name -> ServerCommon.LimitInfo
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_nsq_proto_init() }
func file_nsq_proto_init() {
	if File_nsq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nsq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BanUserInfo); i {
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
		file_nsq_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplexMessage); i {
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
		file_nsq_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendMessage); i {
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
		file_nsq_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReloadConfig); i {
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
		file_nsq_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetMaxPlayerNum); i {
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
		file_nsq_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetLogLevel); i {
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
		file_nsq_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseGameSystem); i {
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
		file_nsq_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitUser); i {
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
			RawDescriptor: file_nsq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nsq_proto_goTypes,
		DependencyIndexes: file_nsq_proto_depIdxs,
		EnumInfos:         file_nsq_proto_enumTypes,
		MessageInfos:      file_nsq_proto_msgTypes,
	}.Build()
	File_nsq_proto = out.File
	file_nsq_proto_rawDesc = nil
	file_nsq_proto_goTypes = nil
	file_nsq_proto_depIdxs = nil
}
