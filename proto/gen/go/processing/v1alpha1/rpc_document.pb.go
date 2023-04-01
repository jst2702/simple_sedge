// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: processing/v1alpha1/rpc_document.proto

package processingv1alpha1

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

type IngestDocumentRequest struct {
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
	Tickers     []string `protobuf:"bytes,10,rep,name=tickers,json=ticker,proto3" json:"tickers,omitempty"`
	Published   string   `protobuf:"bytes,11,opt,name=published,proto3" json:"published,omitempty"`
	Language    string   `protobuf:"bytes,12,opt,name=language,proto3" json:"language,omitempty"`
	ApiKey      string   `protobuf:"bytes,13,opt,name=api_key,proto3" json:"api_key,omitempty"`
}

func (x *IngestDocumentRequest) Reset() {
	*x = IngestDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_processing_v1alpha1_rpc_document_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestDocumentRequest) ProtoMessage() {}

func (x *IngestDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_processing_v1alpha1_rpc_document_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngestDocumentRequest.ProtoReflect.Descriptor instead.
func (*IngestDocumentRequest) Descriptor() ([]byte, []int) {
	return file_processing_v1alpha1_rpc_document_proto_rawDescGZIP(), []int{0}
}

func (x *IngestDocumentRequest) GetGuid() string {
	if x != nil {
		return x.Guid
	}
	return ""
}

func (x *IngestDocumentRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *IngestDocumentRequest) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

func (x *IngestDocumentRequest) GetSiteFull() string {
	if x != nil {
		return x.SiteFull
	}
	return ""
}

func (x *IngestDocumentRequest) GetSiteSection() string {
	if x != nil {
		return x.SiteSection
	}
	return ""
}

func (x *IngestDocumentRequest) GetHeadline() string {
	if x != nil {
		return x.Headline
	}
	return ""
}

func (x *IngestDocumentRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *IngestDocumentRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *IngestDocumentRequest) GetTicker() string {
	if x != nil && x.Ticker != nil {
		return *x.Ticker
	}
	return ""
}

func (x *IngestDocumentRequest) GetTickers() []string {
	if x != nil {
		return x.Tickers
	}
	return nil
}

func (x *IngestDocumentRequest) GetPublished() string {
	if x != nil {
		return x.Published
	}
	return ""
}

func (x *IngestDocumentRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *IngestDocumentRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

type IngestDocumentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Document *Document `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
}

func (x *IngestDocumentResponse) Reset() {
	*x = IngestDocumentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_processing_v1alpha1_rpc_document_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestDocumentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestDocumentResponse) ProtoMessage() {}

func (x *IngestDocumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_processing_v1alpha1_rpc_document_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngestDocumentResponse.ProtoReflect.Descriptor instead.
func (*IngestDocumentResponse) Descriptor() ([]byte, []int) {
	return file_processing_v1alpha1_rpc_document_proto_rawDescGZIP(), []int{1}
}

func (x *IngestDocumentResponse) GetDocument() *Document {
	if x != nil {
		return x.Document
	}
	return nil
}

type ListDocumentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListDocumentsRequest) Reset() {
	*x = ListDocumentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_processing_v1alpha1_rpc_document_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDocumentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDocumentsRequest) ProtoMessage() {}

func (x *ListDocumentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_processing_v1alpha1_rpc_document_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDocumentsRequest.ProtoReflect.Descriptor instead.
func (*ListDocumentsRequest) Descriptor() ([]byte, []int) {
	return file_processing_v1alpha1_rpc_document_proto_rawDescGZIP(), []int{2}
}

func (x *ListDocumentsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListDocumentsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListDocumentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Documents []*Document `protobuf:"bytes,1,rep,name=documents,proto3" json:"documents,omitempty"`
}

func (x *ListDocumentsResponse) Reset() {
	*x = ListDocumentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_processing_v1alpha1_rpc_document_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDocumentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDocumentsResponse) ProtoMessage() {}

func (x *ListDocumentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_processing_v1alpha1_rpc_document_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDocumentsResponse.ProtoReflect.Descriptor instead.
func (*ListDocumentsResponse) Descriptor() ([]byte, []int) {
	return file_processing_v1alpha1_rpc_document_proto_rawDescGZIP(), []int{3}
}

func (x *ListDocumentsResponse) GetDocuments() []*Document {
	if x != nil {
		return x.Documents
	}
	return nil
}

var File_processing_v1alpha1_rpc_document_proto protoreflect.FileDescriptor

var file_processing_v1alpha1_rpc_document_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x22, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xee, 0x02, 0x0a, 0x15, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x67,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x66, 0x75,
	0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x66,
	0x75, 0x6c, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x69, 0x74, 0x65, 0x5f,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x65, 0x61, 0x64, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x65, 0x61, 0x64, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1b, 0x0a,
	0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x72, 0x22, 0x53, 0x0a, 0x16, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x44, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x54, 0x0a,
	0x15, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x42, 0xdd, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42,
	0x10, 0x52, 0x70, 0x63, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x43, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x65, 0x64, 0x67, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02,
	0x13, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x13, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x1f, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_processing_v1alpha1_rpc_document_proto_rawDescOnce sync.Once
	file_processing_v1alpha1_rpc_document_proto_rawDescData = file_processing_v1alpha1_rpc_document_proto_rawDesc
)

func file_processing_v1alpha1_rpc_document_proto_rawDescGZIP() []byte {
	file_processing_v1alpha1_rpc_document_proto_rawDescOnce.Do(func() {
		file_processing_v1alpha1_rpc_document_proto_rawDescData = protoimpl.X.CompressGZIP(file_processing_v1alpha1_rpc_document_proto_rawDescData)
	})
	return file_processing_v1alpha1_rpc_document_proto_rawDescData
}

var file_processing_v1alpha1_rpc_document_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_processing_v1alpha1_rpc_document_proto_goTypes = []interface{}{
	(*IngestDocumentRequest)(nil),  // 0: processing.v1alpha1.IngestDocumentRequest
	(*IngestDocumentResponse)(nil), // 1: processing.v1alpha1.IngestDocumentResponse
	(*ListDocumentsRequest)(nil),   // 2: processing.v1alpha1.ListDocumentsRequest
	(*ListDocumentsResponse)(nil),  // 3: processing.v1alpha1.ListDocumentsResponse
	(*Document)(nil),               // 4: processing.v1alpha1.Document
}
var file_processing_v1alpha1_rpc_document_proto_depIdxs = []int32{
	4, // 0: processing.v1alpha1.IngestDocumentResponse.document:type_name -> processing.v1alpha1.Document
	4, // 1: processing.v1alpha1.ListDocumentsResponse.documents:type_name -> processing.v1alpha1.Document
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_processing_v1alpha1_rpc_document_proto_init() }
func file_processing_v1alpha1_rpc_document_proto_init() {
	if File_processing_v1alpha1_rpc_document_proto != nil {
		return
	}
	file_processing_v1alpha1_document_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_processing_v1alpha1_rpc_document_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngestDocumentRequest); i {
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
		file_processing_v1alpha1_rpc_document_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngestDocumentResponse); i {
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
		file_processing_v1alpha1_rpc_document_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDocumentsRequest); i {
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
		file_processing_v1alpha1_rpc_document_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDocumentsResponse); i {
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
	file_processing_v1alpha1_rpc_document_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_processing_v1alpha1_rpc_document_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_processing_v1alpha1_rpc_document_proto_goTypes,
		DependencyIndexes: file_processing_v1alpha1_rpc_document_proto_depIdxs,
		MessageInfos:      file_processing_v1alpha1_rpc_document_proto_msgTypes,
	}.Build()
	File_processing_v1alpha1_rpc_document_proto = out.File
	file_processing_v1alpha1_rpc_document_proto_rawDesc = nil
	file_processing_v1alpha1_rpc_document_proto_goTypes = nil
	file_processing_v1alpha1_rpc_document_proto_depIdxs = nil
}
