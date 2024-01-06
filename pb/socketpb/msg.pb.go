// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: socket/msg.proto

package socketpb

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

type MsgType int32

const (
	MsgType_HEART      MsgType = 0 //心跳
	MsgType_LOGIN      MsgType = 1 //登录
	MsgType_RESP       MsgType = 2 //统一消息回执
	MsgType_NEW_NOTICE MsgType = 3 //新消息通知
	MsgType_RPA        MsgType = 4 //rpa
	MsgType_RPA_PLAN   MsgType = 5 //rpa计划
)

// Enum value maps for MsgType.
var (
	MsgType_name = map[int32]string{
		0: "HEART",
		1: "LOGIN",
		2: "RESP",
		3: "NEW_NOTICE",
		4: "RPA",
		5: "RPA_PLAN",
	}
	MsgType_value = map[string]int32{
		"HEART":      0,
		"LOGIN":      1,
		"RESP":       2,
		"NEW_NOTICE": 3,
		"RPA":        4,
		"RPA_PLAN":   5,
	}
)

func (x MsgType) Enum() *MsgType {
	p := new(MsgType)
	*p = x
	return p
}

func (x MsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_socket_msg_proto_enumTypes[0].Descriptor()
}

func (MsgType) Type() protoreflect.EnumType {
	return &file_socket_msg_proto_enumTypes[0]
}

func (x MsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgType.Descriptor instead.
func (MsgType) EnumDescriptor() ([]byte, []int) {
	return file_socket_msg_proto_rawDescGZIP(), []int{0}
}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgType int32  `protobuf:"varint,1,opt,name=msgType,proto3" json:"msgType,omitempty"` //消息类型
	ReqId   string `protobuf:"bytes,2,opt,name=reqId,proto3" json:"reqId,omitempty"`      //请求id
	Version int32  `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"` //版本号（消息版本号） 1开始
	Data    []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`        //消息体
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_socket_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_socket_msg_proto_rawDescGZIP(), []int{0}
}

func (x *Msg) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *Msg) GetReqId() string {
	if x != nil {
		return x.ReqId
	}
	return ""
}

func (x *Msg) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Msg) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_socket_msg_proto protoreflect.FileDescriptor

var file_socket_msg_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x63, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x50, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x45, 0x41, 0x52, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x53, 0x50,
	0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x45, 0x57, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x43, 0x45,
	0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x50, 0x41, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x52,
	0x50, 0x41, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x10, 0x05, 0x42, 0x1b, 0x5a, 0x19, 0x64, 0x69, 0x6c,
	0x75, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_socket_msg_proto_rawDescOnce sync.Once
	file_socket_msg_proto_rawDescData = file_socket_msg_proto_rawDesc
)

func file_socket_msg_proto_rawDescGZIP() []byte {
	file_socket_msg_proto_rawDescOnce.Do(func() {
		file_socket_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_socket_msg_proto_rawDescData)
	})
	return file_socket_msg_proto_rawDescData
}

var file_socket_msg_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_socket_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_socket_msg_proto_goTypes = []interface{}{
	(MsgType)(0), // 0: MsgType
	(*Msg)(nil),  // 1: Msg
}
var file_socket_msg_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_socket_msg_proto_init() }
func file_socket_msg_proto_init() {
	if File_socket_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_socket_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
			RawDescriptor: file_socket_msg_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_socket_msg_proto_goTypes,
		DependencyIndexes: file_socket_msg_proto_depIdxs,
		EnumInfos:         file_socket_msg_proto_enumTypes,
		MessageInfos:      file_socket_msg_proto_msgTypes,
	}.Build()
	File_socket_msg_proto = out.File
	file_socket_msg_proto_rawDesc = nil
	file_socket_msg_proto_goTypes = nil
	file_socket_msg_proto_depIdxs = nil
}
