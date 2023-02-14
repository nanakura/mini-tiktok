// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: idl/chat/chat.proto

package chatservice

import (
	context "context"
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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content    string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	CreateTime string `protobuf:"bytes,3,opt,name=createTime,proto3" json:"createTime,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_chat_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_idl_chat_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_idl_chat_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

type MessageActionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ToUserKey  string `protobuf:"bytes,2,opt,name=toUserKey,proto3" json:"toUserKey,omitempty"`
	ActionType string `protobuf:"bytes,3,opt,name=actionType,proto3" json:"actionType,omitempty"`
	Content    string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *MessageActionReq) Reset() {
	*x = MessageActionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_chat_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageActionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageActionReq) ProtoMessage() {}

func (x *MessageActionReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_chat_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageActionReq.ProtoReflect.Descriptor instead.
func (*MessageActionReq) Descriptor() ([]byte, []int) {
	return file_idl_chat_chat_proto_rawDescGZIP(), []int{1}
}

func (x *MessageActionReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *MessageActionReq) GetToUserKey() string {
	if x != nil {
		return x.ToUserKey
	}
	return ""
}

func (x *MessageActionReq) GetActionType() string {
	if x != nil {
		return x.ActionType
	}
	return ""
}

func (x *MessageActionReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type MessageActionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MessageActionResp) Reset() {
	*x = MessageActionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_chat_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageActionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageActionResp) ProtoMessage() {}

func (x *MessageActionResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_chat_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageActionResp.ProtoReflect.Descriptor instead.
func (*MessageActionResp) Descriptor() ([]byte, []int) {
	return file_idl_chat_chat_proto_rawDescGZIP(), []int{2}
}

type MessageChatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatKey    string `protobuf:"bytes,1,opt,name=ChatKey,proto3" json:"ChatKey,omitempty"`
	MsgContent string `protobuf:"bytes,2,opt,name=MsgContent,proto3" json:"MsgContent,omitempty"`
}

func (x *MessageChatReq) Reset() {
	*x = MessageChatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_chat_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageChatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageChatReq) ProtoMessage() {}

func (x *MessageChatReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_chat_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageChatReq.ProtoReflect.Descriptor instead.
func (*MessageChatReq) Descriptor() ([]byte, []int) {
	return file_idl_chat_chat_proto_rawDescGZIP(), []int{3}
}

func (x *MessageChatReq) GetChatKey() string {
	if x != nil {
		return x.ChatKey
	}
	return ""
}

func (x *MessageChatReq) GetMsgContent() string {
	if x != nil {
		return x.MsgContent
	}
	return ""
}

type MessageChatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageList []*Message `protobuf:"bytes,1,rep,name=messageList,proto3" json:"messageList,omitempty"`
}

func (x *MessageChatResp) Reset() {
	*x = MessageChatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_chat_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageChatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageChatResp) ProtoMessage() {}

func (x *MessageChatResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_chat_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageChatResp.ProtoReflect.Descriptor instead.
func (*MessageChatResp) Descriptor() ([]byte, []int) {
	return file_idl_chat_chat_proto_rawDescGZIP(), []int{4}
}

func (x *MessageChatResp) GetMessageList() []*Message {
	if x != nil {
		return x.MessageList
	}
	return nil
}

var File_idl_chat_chat_proto protoreflect.FileDescriptor

var file_idl_chat_chat_proto_rawDesc = []byte{
	0x0a, 0x13, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x53, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x80, 0x01, 0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f,
	0x55, 0x73, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x4a, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x68, 0x61,
	0x74, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x68, 0x61, 0x74,
	0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x73, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4d, 0x73, 0x67, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x51, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x68,
	0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3e, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xca, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x0d, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x1a, 0x26, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x58, 0x0a, 0x0b, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x42, 0x23, 0x5a, 0x21, 0x6d, 0x69, 0x6e, 0x69, 0x5f, 0x74, 0x69, 0x6b, 0x74,
	0x6f, 0x6b, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61,
	0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_chat_chat_proto_rawDescOnce sync.Once
	file_idl_chat_chat_proto_rawDescData = file_idl_chat_chat_proto_rawDesc
)

func file_idl_chat_chat_proto_rawDescGZIP() []byte {
	file_idl_chat_chat_proto_rawDescOnce.Do(func() {
		file_idl_chat_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_chat_chat_proto_rawDescData)
	})
	return file_idl_chat_chat_proto_rawDescData
}

var file_idl_chat_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_idl_chat_chat_proto_goTypes = []interface{}{
	(*Message)(nil),           // 0: service.ChatService.Message
	(*MessageActionReq)(nil),  // 1: service.ChatService.MessageActionReq
	(*MessageActionResp)(nil), // 2: service.ChatService.MessageActionResp
	(*MessageChatReq)(nil),    // 3: service.ChatService.MessageChatReq
	(*MessageChatResp)(nil),   // 4: service.ChatService.MessageChatResp
}
var file_idl_chat_chat_proto_depIdxs = []int32{
	0, // 0: service.ChatService.MessageChatResp.messageList:type_name -> service.ChatService.Message
	1, // 1: service.ChatService.MessageService.MessageAction:input_type -> service.ChatService.MessageActionReq
	3, // 2: service.ChatService.MessageService.MessageChat:input_type -> service.ChatService.MessageChatReq
	2, // 3: service.ChatService.MessageService.MessageAction:output_type -> service.ChatService.MessageActionResp
	4, // 4: service.ChatService.MessageService.MessageChat:output_type -> service.ChatService.MessageChatResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_idl_chat_chat_proto_init() }
func file_idl_chat_chat_proto_init() {
	if File_idl_chat_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_chat_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_idl_chat_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageActionReq); i {
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
		file_idl_chat_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageActionResp); i {
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
		file_idl_chat_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageChatReq); i {
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
		file_idl_chat_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageChatResp); i {
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
			RawDescriptor: file_idl_chat_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_chat_chat_proto_goTypes,
		DependencyIndexes: file_idl_chat_chat_proto_depIdxs,
		MessageInfos:      file_idl_chat_chat_proto_msgTypes,
	}.Build()
	File_idl_chat_chat_proto = out.File
	file_idl_chat_chat_proto_rawDesc = nil
	file_idl_chat_chat_proto_goTypes = nil
	file_idl_chat_chat_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.4.4. DO NOT EDIT.

type MessageService interface {
	MessageAction(ctx context.Context, req *MessageActionReq) (res *MessageActionResp, err error)
	MessageChat(ctx context.Context, req *MessageChatReq) (res *MessageChatResp, err error)
}
