// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.0--rc3
// source: article/v1/article.proto

package v1

import (
	pbentity "grpc/app/article/api/pbentity"
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page uint32 `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	Size uint32 `protobuf:"varint,2,opt,name=Size,proto3" json:"Size,omitempty"`
}

func (x *ListReq) Reset() {
	*x = ListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_v1_article_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReq) ProtoMessage() {}

func (x *ListReq) ProtoReflect() protoreflect.Message {
	mi := &file_article_v1_article_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReq.ProtoReflect.Descriptor instead.
func (*ListReq) Descriptor() ([]byte, []int) {
	return file_article_v1_article_proto_rawDescGZIP(), []int{0}
}

func (x *ListReq) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListReq) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article []*pbentity.Article `protobuf:"bytes,1,rep,name=Article,proto3" json:"Article,omitempty"`
}

func (x *ListRes) Reset() {
	*x = ListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_v1_article_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRes) ProtoMessage() {}

func (x *ListRes) ProtoReflect() protoreflect.Message {
	mi := &file_article_v1_article_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRes.ProtoReflect.Descriptor instead.
func (*ListRes) Descriptor() ([]byte, []int) {
	return file_article_v1_article_proto_rawDescGZIP(), []int{1}
}

func (x *ListRes) GetArticle() []*pbentity.Article {
	if x != nil {
		return x.Article
	}
	return nil
}

var File_article_v1_article_proto protoreflect.FileDescriptor

var file_article_v1_article_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x1a, 0x16, 0x70, 0x62, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x07, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x36,
	0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x07, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x07, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x32, 0x35, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x12, 0x2a, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x42, 0x1e, 0x5a,
	0x1c, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_article_v1_article_proto_rawDescOnce sync.Once
	file_article_v1_article_proto_rawDescData = file_article_v1_article_proto_rawDesc
)

func file_article_v1_article_proto_rawDescGZIP() []byte {
	file_article_v1_article_proto_rawDescOnce.Do(func() {
		file_article_v1_article_proto_rawDescData = protoimpl.X.CompressGZIP(file_article_v1_article_proto_rawDescData)
	})
	return file_article_v1_article_proto_rawDescData
}

var file_article_v1_article_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_article_v1_article_proto_goTypes = []interface{}{
	(*ListReq)(nil),          // 0: article.ListReq
	(*ListRes)(nil),          // 1: article.ListRes
	(*pbentity.Article)(nil), // 2: pbentity.Article
}
var file_article_v1_article_proto_depIdxs = []int32{
	2, // 0: article.ListRes.Article:type_name -> pbentity.Article
	0, // 1: article.Article.List:input_type -> article.ListReq
	1, // 2: article.Article.List:output_type -> article.ListRes
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_article_v1_article_proto_init() }
func file_article_v1_article_proto_init() {
	if File_article_v1_article_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_article_v1_article_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReq); i {
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
		file_article_v1_article_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRes); i {
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
			RawDescriptor: file_article_v1_article_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_article_v1_article_proto_goTypes,
		DependencyIndexes: file_article_v1_article_proto_depIdxs,
		MessageInfos:      file_article_v1_article_proto_msgTypes,
	}.Build()
	File_article_v1_article_proto = out.File
	file_article_v1_article_proto_rawDesc = nil
	file_article_v1_article_proto_goTypes = nil
	file_article_v1_article_proto_depIdxs = nil
}