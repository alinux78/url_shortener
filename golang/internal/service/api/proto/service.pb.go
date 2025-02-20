// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.6.1
// source: service.proto

package proto

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

type UrlShortenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UrlShortenRequest) Reset() {
	*x = UrlShortenRequest{}
	mi := &file_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UrlShortenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlShortenRequest) ProtoMessage() {}

func (x *UrlShortenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlShortenRequest.ProtoReflect.Descriptor instead.
func (*UrlShortenRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *UrlShortenRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type UrlShortenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShortUrl      string                 `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UrlShortenResponse) Reset() {
	*x = UrlShortenResponse{}
	mi := &file_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UrlShortenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlShortenResponse) ProtoMessage() {}

func (x *UrlShortenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlShortenResponse.ProtoReflect.Descriptor instead.
func (*UrlShortenResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *UrlShortenResponse) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type UrlResolveRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShortUrl      string                 `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UrlResolveRequest) Reset() {
	*x = UrlResolveRequest{}
	mi := &file_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UrlResolveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlResolveRequest) ProtoMessage() {}

func (x *UrlResolveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlResolveRequest.ProtoReflect.Descriptor instead.
func (*UrlResolveRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *UrlResolveRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type UrlResolveResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UrlResolveResponse) Reset() {
	*x = UrlResolveResponse{}
	mi := &file_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UrlResolveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlResolveResponse) ProtoMessage() {}

func (x *UrlResolveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlResolveResponse.ProtoReflect.Descriptor instead.
func (*UrlResolveResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *UrlResolveResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x25, 0x0a, 0x11, 0x55, 0x72, 0x6c, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22,
	0x31, 0x0a, 0x12, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55,
	0x72, 0x6c, 0x22, 0x30, 0x0a, 0x11, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x72, 0x6c, 0x22, 0x26, 0x0a, 0x12, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32, 0x9d, 0x01, 0x0a,
	0x13, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x12,
	0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x72,
	0x6c, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData []byte
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_service_proto_rawDesc), len(file_service_proto_rawDesc)))
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_service_proto_goTypes = []any{
	(*UrlShortenRequest)(nil),  // 0: service.UrlShortenRequest
	(*UrlShortenResponse)(nil), // 1: service.UrlShortenResponse
	(*UrlResolveRequest)(nil),  // 2: service.UrlResolveRequest
	(*UrlResolveResponse)(nil), // 3: service.UrlResolveResponse
}
var file_service_proto_depIdxs = []int32{
	0, // 0: service.UrlShortenerService.Shorten:input_type -> service.UrlShortenRequest
	2, // 1: service.UrlShortenerService.Resolve:input_type -> service.UrlResolveRequest
	1, // 2: service.UrlShortenerService.Shorten:output_type -> service.UrlShortenResponse
	3, // 3: service.UrlShortenerService.Resolve:output_type -> service.UrlResolveResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_service_proto_rawDesc), len(file_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
