// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.0--rc3
// source: protobuf/user/user.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Passport string `protobuf:"bytes,1,opt,name=Passport,proto3" json:"Passport,omitempty"` // v: required
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"` // v: required
	Nickname string `protobuf:"bytes,3,opt,name=Nickname,proto3" json:"Nickname,omitempty"` // v: required
}

func (x *CreateReq) Reset() {
	*x = CreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReq) ProtoMessage() {}

func (x *CreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReq.ProtoReflect.Descriptor instead.
func (*CreateReq) Descriptor() ([]byte, []int) {
	return file_protobuf_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *CreateReq) GetPassport() string {
	if x != nil {
		return x.Passport
	}
	return ""
}

func (x *CreateReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

type CreateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateRes) Reset() {
	*x = CreateRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRes) ProtoMessage() {}

func (x *CreateRes) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRes.ProtoReflect.Descriptor instead.
func (*CreateRes) Descriptor() ([]byte, []int) {
	return file_protobuf_user_user_proto_rawDescGZIP(), []int{1}
}

var File_protobuf_user_user_proto protoreflect.FileDescriptor

var file_protobuf_user_user_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x62, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0b, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x32, 0x3c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protobuf_user_user_proto_rawDescOnce sync.Once
	file_protobuf_user_user_proto_rawDescData = file_protobuf_user_user_proto_rawDesc
)

func file_protobuf_user_user_proto_rawDescGZIP() []byte {
	file_protobuf_user_user_proto_rawDescOnce.Do(func() {
		file_protobuf_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_user_user_proto_rawDescData)
	})
	return file_protobuf_user_user_proto_rawDescData
}

var file_protobuf_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protobuf_user_user_proto_goTypes = []interface{}{
	(*CreateReq)(nil), // 0: pbentity.CreateReq
	(*CreateRes)(nil), // 1: pbentity.CreateRes
}
var file_protobuf_user_user_proto_depIdxs = []int32{
	0, // 0: pbentity.User.Create:input_type -> pbentity.CreateReq
	1, // 1: pbentity.User.Create:output_type -> pbentity.CreateRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_user_user_proto_init() }
func file_protobuf_user_user_proto_init() {
	if File_protobuf_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReq); i {
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
		file_protobuf_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRes); i {
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
			RawDescriptor: file_protobuf_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_user_user_proto_goTypes,
		DependencyIndexes: file_protobuf_user_user_proto_depIdxs,
		MessageInfos:      file_protobuf_user_user_proto_msgTypes,
	}.Build()
	File_protobuf_user_user_proto = out.File
	file_protobuf_user_user_proto_rawDesc = nil
	file_protobuf_user_user_proto_goTypes = nil
	file_protobuf_user_user_proto_depIdxs = nil
}
