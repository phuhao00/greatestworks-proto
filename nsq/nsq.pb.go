// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: nsq.proto

package nsq

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

type NsqCommand int32

const (
	NsqCommand_None NsqCommand = 0
)

// Enum value maps for NsqCommand.
var (
	NsqCommand_name = map[int32]string{
		0: "None",
	}
	NsqCommand_value = map[string]int32{
		"None": 0,
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
		mi := &file_nsq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplexMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexMessage) ProtoMessage() {}

func (x *ComplexMessage) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ComplexMessage.ProtoReflect.Descriptor instead.
func (*ComplexMessage) Descriptor() ([]byte, []int) {
	return file_nsq_proto_rawDescGZIP(), []int{0}
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

type Mail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Mail) Reset() {
	*x = Mail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nsq_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mail) ProtoMessage() {}

func (x *Mail) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Mail.ProtoReflect.Descriptor instead.
func (*Mail) Descriptor() ([]byte, []int) {
	return file_nsq_proto_rawDescGZIP(), []int{1}
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

var File_nsq_proto protoreflect.FileDescriptor

var file_nsq_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6e, 0x73, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e, 0x73, 0x71,
	0x22, 0x5b, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x21, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0f, 0x2e, 0x6e, 0x73, 0x71, 0x2e, 0x4e, 0x73, 0x71, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x06, 0x0a,
	0x04, 0x4d, 0x61, 0x69, 0x6c, 0x22, 0x66, 0x0a, 0x0d, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x03, 0x63, 0x6d,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6e, 0x73, 0x71, 0x2e, 0x4e, 0x73,
	0x71, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x2a, 0x16, 0x0a,
	0x0a, 0x4e, 0x73, 0x71, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x6f, 0x6e, 0x65, 0x10, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x68, 0x75, 0x68, 0x61, 0x6f, 0x30, 0x30, 0x2f, 0x67, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6e, 0x73, 0x71, 0x3b, 0x6e, 0x73, 0x71, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_nsq_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_nsq_proto_goTypes = []interface{}{
	(NsqCommand)(0),        // 0: nsq.NsqCommand
	(*ComplexMessage)(nil), // 1: nsq.ComplexMessage
	(*Mail)(nil),           // 2: nsq.Mail
	(*FriendMessage)(nil),  // 3: nsq.FriendMessage
}
var file_nsq_proto_depIdxs = []int32{
	0, // 0: nsq.ComplexMessage.cmd:type_name -> nsq.NsqCommand
	0, // 1: nsq.FriendMessage.cmd:type_name -> nsq.NsqCommand
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_nsq_proto_init() }
func file_nsq_proto_init() {
	if File_nsq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nsq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_nsq_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mail); i {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nsq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
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
