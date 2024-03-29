// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: apis/v1alpha1/document.proto

package v1alpha1

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

type Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guid        string   `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Url         string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Site        string   `protobuf:"bytes,3,opt,name=site,proto3" json:"site,omitempty"`
	SiteFull    string   `protobuf:"bytes,4,opt,name=site_full,proto3" json:"site_full,omitempty"`
	SiteSection string   `protobuf:"bytes,5,opt,name=site_section,proto3" json:"site_section,omitempty"`
	Headline    string   `protobuf:"bytes,6,opt,name=headline,proto3" json:"headline,omitempty"`
	Title       string   `protobuf:"bytes,7,opt,name=title,proto3" json:"title,omitempty"`
	Body        string   `protobuf:"bytes,8,opt,name=body,proto3" json:"body,omitempty"`
	Ticker      *string  `protobuf:"bytes,9,opt,name=ticker,proto3,oneof" json:"ticker,omitempty"`
	Tickers     []string `protobuf:"bytes,10,rep,name=tickers,proto3" json:"tickers,omitempty"`
	Published   string   `protobuf:"bytes,11,opt,name=published,proto3" json:"published,omitempty"`
	Language    string   `protobuf:"bytes,12,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *Document) Reset() {
	*x = Document{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1alpha1_document_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document) ProtoMessage() {}

func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1alpha1_document_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document.ProtoReflect.Descriptor instead.
func (*Document) Descriptor() ([]byte, []int) {
	return file_apis_v1alpha1_document_proto_rawDescGZIP(), []int{0}
}

func (x *Document) GetGuid() string {
	if x != nil {
		return x.Guid
	}
	return ""
}

func (x *Document) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Document) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

func (x *Document) GetSiteFull() string {
	if x != nil {
		return x.SiteFull
	}
	return ""
}

func (x *Document) GetSiteSection() string {
	if x != nil {
		return x.SiteSection
	}
	return ""
}

func (x *Document) GetHeadline() string {
	if x != nil {
		return x.Headline
	}
	return ""
}

func (x *Document) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Document) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Document) GetTicker() string {
	if x != nil && x.Ticker != nil {
		return *x.Ticker
	}
	return ""
}

func (x *Document) GetTickers() []string {
	if x != nil {
		return x.Tickers
	}
	return nil
}

func (x *Document) GetPublished() string {
	if x != nil {
		return x.Published
	}
	return ""
}

func (x *Document) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

var File_apis_v1alpha1_document_proto protoreflect.FileDescriptor

var file_apis_v1alpha1_document_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x61, 0x70, 0x69, 0x73, 0x22, 0xc8, 0x02, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x67, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x69, 0x74, 0x65, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x69, 0x74, 0x65, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x69, 0x74,
	0x65, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x12, 0x1b, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x88, 0x01, 0x01,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x42,
	0x75, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x42, 0x0d, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x73, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02,
	0x04, 0x41, 0x70, 0x69, 0x73, 0xca, 0x02, 0x04, 0x41, 0x70, 0x69, 0x73, 0xe2, 0x02, 0x10, 0x41,
	0x70, 0x69, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x04, 0x41, 0x70, 0x69, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apis_v1alpha1_document_proto_rawDescOnce sync.Once
	file_apis_v1alpha1_document_proto_rawDescData = file_apis_v1alpha1_document_proto_rawDesc
)

func file_apis_v1alpha1_document_proto_rawDescGZIP() []byte {
	file_apis_v1alpha1_document_proto_rawDescOnce.Do(func() {
		file_apis_v1alpha1_document_proto_rawDescData = protoimpl.X.CompressGZIP(file_apis_v1alpha1_document_proto_rawDescData)
	})
	return file_apis_v1alpha1_document_proto_rawDescData
}

var file_apis_v1alpha1_document_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_apis_v1alpha1_document_proto_goTypes = []interface{}{
	(*Document)(nil), // 0: apis.Document
}
var file_apis_v1alpha1_document_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apis_v1alpha1_document_proto_init() }
func file_apis_v1alpha1_document_proto_init() {
	if File_apis_v1alpha1_document_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apis_v1alpha1_document_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Document); i {
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
	file_apis_v1alpha1_document_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apis_v1alpha1_document_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apis_v1alpha1_document_proto_goTypes,
		DependencyIndexes: file_apis_v1alpha1_document_proto_depIdxs,
		MessageInfos:      file_apis_v1alpha1_document_proto_msgTypes,
	}.Build()
	File_apis_v1alpha1_document_proto = out.File
	file_apis_v1alpha1_document_proto_rawDesc = nil
	file_apis_v1alpha1_document_proto_goTypes = nil
	file_apis_v1alpha1_document_proto_depIdxs = nil
}
