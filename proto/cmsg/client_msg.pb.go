// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.0
// 	protoc        v3.13.0
// source: cmsg/client_msg.proto

// 客户端与服务端交互协议

package cmsg

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Code int32

const (
	Code_OK      Code = 0
	Code_Unknown Code = 1
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0: "OK",
		1: "Unknown",
	}
	Code_value = map[string]int32{
		"OK":      0,
		"Unknown": 1,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_cmsg_client_msg_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_cmsg_client_msg_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_cmsg_client_msg_proto_rawDescGZIP(), []int{0}
}

// 登录服务器并进入聊天室
type CReqLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *CReqLogin) Reset() {
	*x = CReqLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmsg_client_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CReqLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CReqLogin) ProtoMessage() {}

func (x *CReqLogin) ProtoReflect() protoreflect.Message {
	mi := &file_cmsg_client_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CReqLogin.ProtoReflect.Descriptor instead.
func (*CReqLogin) Descriptor() ([]byte, []int) {
	return file_cmsg_client_msg_proto_rawDescGZIP(), []int{0}
}

func (x *CReqLogin) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type SRespLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code Code `protobuf:"varint,1,opt,name=code,proto3,enum=cmsg.Code" json:"code,omitempty"`
}

func (x *SRespLogin) Reset() {
	*x = SRespLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmsg_client_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SRespLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SRespLogin) ProtoMessage() {}

func (x *SRespLogin) ProtoReflect() protoreflect.Message {
	mi := &file_cmsg_client_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SRespLogin.ProtoReflect.Descriptor instead.
func (*SRespLogin) Descriptor() ([]byte, []int) {
	return file_cmsg_client_msg_proto_rawDescGZIP(), []int{1}
}

func (x *SRespLogin) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_OK
}

// 发送聊天消息
type CReqSendChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CReqSendChatMessage) Reset() {
	*x = CReqSendChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmsg_client_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CReqSendChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CReqSendChatMessage) ProtoMessage() {}

func (x *CReqSendChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cmsg_client_msg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CReqSendChatMessage.ProtoReflect.Descriptor instead.
func (*CReqSendChatMessage) Descriptor() ([]byte, []int) {
	return file_cmsg_client_msg_proto_rawDescGZIP(), []int{2}
}

func (x *CReqSendChatMessage) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type SRespSendChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code Code `protobuf:"varint,1,opt,name=code,proto3,enum=cmsg.Code" json:"code,omitempty"`
}

func (x *SRespSendChatMessage) Reset() {
	*x = SRespSendChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmsg_client_msg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SRespSendChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SRespSendChatMessage) ProtoMessage() {}

func (x *SRespSendChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cmsg_client_msg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SRespSendChatMessage.ProtoReflect.Descriptor instead.
func (*SRespSendChatMessage) Descriptor() ([]byte, []int) {
	return file_cmsg_client_msg_proto_rawDescGZIP(), []int{3}
}

func (x *SRespSendChatMessage) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_OK
}

// 有其他玩家进入聊天室
type SNotifyUserEnter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *SNotifyUserEnter) Reset() {
	*x = SNotifyUserEnter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmsg_client_msg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SNotifyUserEnter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SNotifyUserEnter) ProtoMessage() {}

func (x *SNotifyUserEnter) ProtoReflect() protoreflect.Message {
	mi := &file_cmsg_client_msg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SNotifyUserEnter.ProtoReflect.Descriptor instead.
func (*SNotifyUserEnter) Descriptor() ([]byte, []int) {
	return file_cmsg_client_msg_proto_rawDescGZIP(), []int{4}
}

func (x *SNotifyUserEnter) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

// 有其他玩家离开聊天室
type SNotifyUserLeave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *SNotifyUserLeave) Reset() {
	*x = SNotifyUserLeave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmsg_client_msg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SNotifyUserLeave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SNotifyUserLeave) ProtoMessage() {}

func (x *SNotifyUserLeave) ProtoReflect() protoreflect.Message {
	mi := &file_cmsg_client_msg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SNotifyUserLeave.ProtoReflect.Descriptor instead.
func (*SNotifyUserLeave) Descriptor() ([]byte, []int) {
	return file_cmsg_client_msg_proto_rawDescGZIP(), []int{5}
}

func (x *SNotifyUserLeave) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

// 同步其他玩家的聊天消息
type SNotifyUserChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Text    string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *SNotifyUserChatMessage) Reset() {
	*x = SNotifyUserChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmsg_client_msg_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SNotifyUserChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SNotifyUserChatMessage) ProtoMessage() {}

func (x *SNotifyUserChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cmsg_client_msg_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SNotifyUserChatMessage.ProtoReflect.Descriptor instead.
func (*SNotifyUserChatMessage) Descriptor() ([]byte, []int) {
	return file_cmsg_client_msg_proto_rawDescGZIP(), []int{6}
}

func (x *SNotifyUserChatMessage) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *SNotifyUserChatMessage) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type SNotifyUserList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accounts []string `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *SNotifyUserList) Reset() {
	*x = SNotifyUserList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmsg_client_msg_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SNotifyUserList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SNotifyUserList) ProtoMessage() {}

func (x *SNotifyUserList) ProtoReflect() protoreflect.Message {
	mi := &file_cmsg_client_msg_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SNotifyUserList.ProtoReflect.Descriptor instead.
func (*SNotifyUserList) Descriptor() ([]byte, []int) {
	return file_cmsg_client_msg_proto_rawDescGZIP(), []int{7}
}

func (x *SNotifyUserList) GetAccounts() []string {
	if x != nil {
		return x.Accounts
	}
	return nil
}

var File_cmsg_client_msg_proto protoreflect.FileDescriptor

var file_cmsg_client_msg_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6d, 0x73, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x73,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6d, 0x73, 0x67, 0x22, 0x25, 0x0a,
	0x09, 0x43, 0x52, 0x65, 0x71, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x0a, 0x53, 0x52, 0x65, 0x73, 0x70, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x1e, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0a, 0x2e, 0x63, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x22, 0x29, 0x0a, 0x13, 0x43, 0x52, 0x65, 0x71, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68,
	0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x36, 0x0a,
	0x14, 0x53, 0x52, 0x65, 0x73, 0x70, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x63, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x2c, 0x0a, 0x10, 0x53, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x10, 0x53, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x46, 0x0a, 0x16, 0x53, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x2d, 0x0a, 0x0f, 0x53, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2a, 0x1b, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmsg_client_msg_proto_rawDescOnce sync.Once
	file_cmsg_client_msg_proto_rawDescData = file_cmsg_client_msg_proto_rawDesc
)

func file_cmsg_client_msg_proto_rawDescGZIP() []byte {
	file_cmsg_client_msg_proto_rawDescOnce.Do(func() {
		file_cmsg_client_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmsg_client_msg_proto_rawDescData)
	})
	return file_cmsg_client_msg_proto_rawDescData
}

var file_cmsg_client_msg_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cmsg_client_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cmsg_client_msg_proto_goTypes = []interface{}{
	(Code)(0),                      // 0: cmsg.Code
	(*CReqLogin)(nil),              // 1: cmsg.CReqLogin
	(*SRespLogin)(nil),             // 2: cmsg.SRespLogin
	(*CReqSendChatMessage)(nil),    // 3: cmsg.CReqSendChatMessage
	(*SRespSendChatMessage)(nil),   // 4: cmsg.SRespSendChatMessage
	(*SNotifyUserEnter)(nil),       // 5: cmsg.SNotifyUserEnter
	(*SNotifyUserLeave)(nil),       // 6: cmsg.SNotifyUserLeave
	(*SNotifyUserChatMessage)(nil), // 7: cmsg.SNotifyUserChatMessage
	(*SNotifyUserList)(nil),        // 8: cmsg.SNotifyUserList
}
var file_cmsg_client_msg_proto_depIdxs = []int32{
	0, // 0: cmsg.SRespLogin.code:type_name -> cmsg.Code
	0, // 1: cmsg.SRespSendChatMessage.code:type_name -> cmsg.Code
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cmsg_client_msg_proto_init() }
func file_cmsg_client_msg_proto_init() {
	if File_cmsg_client_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmsg_client_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CReqLogin); i {
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
		file_cmsg_client_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SRespLogin); i {
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
		file_cmsg_client_msg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CReqSendChatMessage); i {
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
		file_cmsg_client_msg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SRespSendChatMessage); i {
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
		file_cmsg_client_msg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SNotifyUserEnter); i {
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
		file_cmsg_client_msg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SNotifyUserLeave); i {
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
		file_cmsg_client_msg_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SNotifyUserChatMessage); i {
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
		file_cmsg_client_msg_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SNotifyUserList); i {
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
			RawDescriptor: file_cmsg_client_msg_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmsg_client_msg_proto_goTypes,
		DependencyIndexes: file_cmsg_client_msg_proto_depIdxs,
		EnumInfos:         file_cmsg_client_msg_proto_enumTypes,
		MessageInfos:      file_cmsg_client_msg_proto_msgTypes,
	}.Build()
	File_cmsg_client_msg_proto = out.File
	file_cmsg_client_msg_proto_rawDesc = nil
	file_cmsg_client_msg_proto_goTypes = nil
	file_cmsg_client_msg_proto_depIdxs = nil
}
