// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.0
// source: shorturl.proto

package rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 定义请求消息，对应 types.ConvertRequest
type ConvertRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LongUrl       string                 `protobuf:"bytes,1,opt,name=long_url,json=longUrl,proto3" json:"long_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConvertRequest) Reset() {
	*x = ConvertRequest{}
	mi := &file_shorturl_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConvertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertRequest) ProtoMessage() {}

func (x *ConvertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shorturl_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertRequest.ProtoReflect.Descriptor instead.
func (*ConvertRequest) Descriptor() ([]byte, []int) {
	return file_shorturl_proto_rawDescGZIP(), []int{0}
}

func (x *ConvertRequest) GetLongUrl() string {
	if x != nil {
		return x.LongUrl
	}
	return ""
}

// 定义响应消息，对应 types.ConvertResponse
type ConvertResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShortUrl      string                 `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConvertResponse) Reset() {
	*x = ConvertResponse{}
	mi := &file_shorturl_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConvertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertResponse) ProtoMessage() {}

func (x *ConvertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shorturl_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertResponse.ProtoReflect.Descriptor instead.
func (*ConvertResponse) Descriptor() ([]byte, []int) {
	return file_shorturl_proto_rawDescGZIP(), []int{1}
}

func (x *ConvertResponse) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

var File_shorturl_proto protoreflect.FileDescriptor

const file_shorturl_proto_rawDesc = "" +
	"\n" +
	"\x0eshorturl.proto\x12\bshortURL\"+\n" +
	"\x0eConvertRequest\x12\x19\n" +
	"\blong_url\x18\x01 \x01(\tR\alongUrl\".\n" +
	"\x0fConvertResponse\x12\x1b\n" +
	"\tshort_url\x18\x01 \x01(\tR\bshortUrl2Q\n" +
	"\x0fShortUrlService\x12>\n" +
	"\aConvert\x12\x18.shortURL.ConvertRequest\x1a\x19.shortURL.ConvertResponseB\aZ\x05./rpcb\x06proto3"

var (
	file_shorturl_proto_rawDescOnce sync.Once
	file_shorturl_proto_rawDescData []byte
)

func file_shorturl_proto_rawDescGZIP() []byte {
	file_shorturl_proto_rawDescOnce.Do(func() {
		file_shorturl_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_shorturl_proto_rawDesc), len(file_shorturl_proto_rawDesc)))
	})
	return file_shorturl_proto_rawDescData
}

var file_shorturl_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shorturl_proto_goTypes = []any{
	(*ConvertRequest)(nil),  // 0: shortURL.ConvertRequest
	(*ConvertResponse)(nil), // 1: shortURL.ConvertResponse
}
var file_shorturl_proto_depIdxs = []int32{
	0, // 0: shortURL.ShortUrlService.Convert:input_type -> shortURL.ConvertRequest
	1, // 1: shortURL.ShortUrlService.Convert:output_type -> shortURL.ConvertResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shorturl_proto_init() }
func file_shorturl_proto_init() {
	if File_shorturl_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_shorturl_proto_rawDesc), len(file_shorturl_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shorturl_proto_goTypes,
		DependencyIndexes: file_shorturl_proto_depIdxs,
		MessageInfos:      file_shorturl_proto_msgTypes,
	}.Build()
	File_shorturl_proto = out.File
	file_shorturl_proto_goTypes = nil
	file_shorturl_proto_depIdxs = nil
}
