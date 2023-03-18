// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: processing/v1alpha1/document.proto

package processingv1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Platform int32

const (
	Platform_PLATFORM_UNDEFINED         Platform = 0
	Platform_PLATFORM_BUSINESS_NEWSWIRE Platform = 1
	Platform_PLATFORM_GLOBAL_NEWSWIRE   Platform = 2
	Platform_PLATFORM_PR_NEWSWIRE       Platform = 3
	Platform_PLATFORM_ACCESSWIRE        Platform = 4
	Platform_PLATFORM_MOTLEY_FOOL       Platform = 5
)

// Enum value maps for Platform.
var (
	Platform_name = map[int32]string{
		0: "PLATFORM_UNDEFINED",
		1: "PLATFORM_BUSINESS_NEWSWIRE",
		2: "PLATFORM_GLOBAL_NEWSWIRE",
		3: "PLATFORM_PR_NEWSWIRE",
		4: "PLATFORM_ACCESSWIRE",
		5: "PLATFORM_MOTLEY_FOOL",
	}
	Platform_value = map[string]int32{
		"PLATFORM_UNDEFINED":         0,
		"PLATFORM_BUSINESS_NEWSWIRE": 1,
		"PLATFORM_GLOBAL_NEWSWIRE":   2,
		"PLATFORM_PR_NEWSWIRE":       3,
		"PLATFORM_ACCESSWIRE":        4,
		"PLATFORM_MOTLEY_FOOL":       5,
	}
)

func (x Platform) Enum() *Platform {
	p := new(Platform)
	*p = x
	return p
}

func (x Platform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Platform) Descriptor() protoreflect.EnumDescriptor {
	return file_processing_v1alpha1_document_proto_enumTypes[0].Descriptor()
}

func (Platform) Type() protoreflect.EnumType {
	return &file_processing_v1alpha1_document_proto_enumTypes[0]
}

func (x Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Platform.Descriptor instead.
func (Platform) EnumDescriptor() ([]byte, []int) {
	return file_processing_v1alpha1_document_proto_rawDescGZIP(), []int{0}
}

type DocStatus int32

const (
	DocStatus_STATUS_UNDEFINED DocStatus = 0
	DocStatus_STATUS_SUBMITTED DocStatus = 1
	DocStatus_STATUS_FAULTED   DocStatus = 2
	DocStatus_STATUS_PROCESSED DocStatus = 3
	DocStatus_STATUS_COMPLETED DocStatus = 4
)

// Enum value maps for DocStatus.
var (
	DocStatus_name = map[int32]string{
		0: "STATUS_UNDEFINED",
		1: "STATUS_SUBMITTED",
		2: "STATUS_FAULTED",
		3: "STATUS_PROCESSED",
		4: "STATUS_COMPLETED",
	}
	DocStatus_value = map[string]int32{
		"STATUS_UNDEFINED": 0,
		"STATUS_SUBMITTED": 1,
		"STATUS_FAULTED":   2,
		"STATUS_PROCESSED": 3,
		"STATUS_COMPLETED": 4,
	}
)

func (x DocStatus) Enum() *DocStatus {
	p := new(DocStatus)
	*p = x
	return p
}

func (x DocStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DocStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_processing_v1alpha1_document_proto_enumTypes[1].Descriptor()
}

func (DocStatus) Type() protoreflect.EnumType {
	return &file_processing_v1alpha1_document_proto_enumTypes[1]
}

func (x DocStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DocStatus.Descriptor instead.
func (DocStatus) EnumDescriptor() ([]byte, []int) {
	return file_processing_v1alpha1_document_proto_rawDescGZIP(), []int{1}
}

type DocMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string `protobuf:"bytes,1,opt,name=uuid,json=encounter_id,proto3" json:"uuid,omitempty"`
	Url         string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Site        string `protobuf:"bytes,3,opt,name=site,proto3" json:"site,omitempty"`
	SiteFull    string `protobuf:"bytes,4,opt,name=site_full,proto3" json:"site_full,omitempty"`
	SiteSection string `protobuf:"bytes,5,opt,name=site_section,proto3" json:"site_section,omitempty"`
	Headline    string `protobuf:"bytes,6,opt,name=headline,proto3" json:"headline,omitempty"`
	Title       string `protobuf:"bytes,7,opt,name=title,proto3" json:"title,omitempty"`
	Body        string `protobuf:"bytes,8,opt,name=body,proto3" json:"body,omitempty"`
	Published   string `protobuf:"bytes,9,opt,name=published,proto3" json:"published,omitempty"`
	Language    string `protobuf:"bytes,10,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *DocMetadata) Reset() {
	*x = DocMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_processing_v1alpha1_document_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocMetadata) ProtoMessage() {}

func (x *DocMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_processing_v1alpha1_document_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocMetadata.ProtoReflect.Descriptor instead.
func (*DocMetadata) Descriptor() ([]byte, []int) {
	return file_processing_v1alpha1_document_proto_rawDescGZIP(), []int{0}
}

func (x *DocMetadata) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *DocMetadata) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *DocMetadata) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

func (x *DocMetadata) GetSiteFull() string {
	if x != nil {
		return x.SiteFull
	}
	return ""
}

func (x *DocMetadata) GetSiteSection() string {
	if x != nil {
		return x.SiteSection
	}
	return ""
}

func (x *DocMetadata) GetHeadline() string {
	if x != nil {
		return x.Headline
	}
	return ""
}

func (x *DocMetadata) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DocMetadata) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *DocMetadata) GetPublished() string {
	if x != nil {
		return x.Published
	}
	return ""
}

func (x *DocMetadata) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Platform Platform `protobuf:"varint,3,opt,name=platform,proto3,enum=processing.v1alpha1.Platform" json:"platform,omitempty"`
	Size     int64    `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	// indicates if the document has been succesfully processed
	Status    DocStatus              `protobuf:"varint,5,opt,name=status,proto3,enum=processing.v1alpha1.DocStatus" json:"status,omitempty"`
	Text      string                 `protobuf:"bytes,6,opt,name=text,proto3" json:"text,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,proto3" json:"updated_at,omitempty"`
	Metadata  *DocMetadata           `protobuf:"bytes,9,opt,name=metadata,json=updated_at,proto3" json:"metadata,omitempty"`
}

func (x *Document) Reset() {
	*x = Document{}
	if protoimpl.UnsafeEnabled {
		mi := &file_processing_v1alpha1_document_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document) ProtoMessage() {}

func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_processing_v1alpha1_document_proto_msgTypes[1]
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
	return file_processing_v1alpha1_document_proto_rawDescGZIP(), []int{1}
}

func (x *Document) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Document) GetPlatform() Platform {
	if x != nil {
		return x.Platform
	}
	return Platform_PLATFORM_UNDEFINED
}

func (x *Document) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Document) GetStatus() DocStatus {
	if x != nil {
		return x.Status
	}
	return DocStatus_STATUS_UNDEFINED
}

func (x *Document) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Document) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Document) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Document) GetMetadata() *DocMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_processing_v1alpha1_document_proto protoreflect.FileDescriptor

var file_processing_v1alpha1_document_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x02, 0x0a, 0x0b, 0x44,
	0x6f, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x69, 0x74, 0x65, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x69,
	0x74, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0xed,
	0x02, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x08, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x08, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x44, 0x6f, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x12, 0x3e,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x2a, 0xad,
	0x01, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x16, 0x0a, 0x12, 0x50,
	0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f,
	0x42, 0x55, 0x53, 0x49, 0x4e, 0x45, 0x53, 0x53, 0x5f, 0x4e, 0x45, 0x57, 0x53, 0x57, 0x49, 0x52,
	0x45, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f,
	0x47, 0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x5f, 0x4e, 0x45, 0x57, 0x53, 0x57, 0x49, 0x52, 0x45, 0x10,
	0x02, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x50, 0x52,
	0x5f, 0x4e, 0x45, 0x57, 0x53, 0x57, 0x49, 0x52, 0x45, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x50,
	0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x57, 0x49,
	0x52, 0x45, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d,
	0x5f, 0x4d, 0x4f, 0x54, 0x4c, 0x45, 0x59, 0x5f, 0x46, 0x4f, 0x4f, 0x4c, 0x10, 0x05, 0x2a, 0x77,
	0x0a, 0x09, 0x44, 0x6f, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x10, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x42, 0x4d,
	0x49, 0x54, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x45, 0x44, 0x10,
	0x03, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x50,
	0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x04, 0x42, 0xda, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x42, 0x0d, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x65, 0x64, 0x67,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa,
	0x02, 0x13, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x13, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x1f, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_processing_v1alpha1_document_proto_rawDescOnce sync.Once
	file_processing_v1alpha1_document_proto_rawDescData = file_processing_v1alpha1_document_proto_rawDesc
)

func file_processing_v1alpha1_document_proto_rawDescGZIP() []byte {
	file_processing_v1alpha1_document_proto_rawDescOnce.Do(func() {
		file_processing_v1alpha1_document_proto_rawDescData = protoimpl.X.CompressGZIP(file_processing_v1alpha1_document_proto_rawDescData)
	})
	return file_processing_v1alpha1_document_proto_rawDescData
}

var file_processing_v1alpha1_document_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_processing_v1alpha1_document_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_processing_v1alpha1_document_proto_goTypes = []interface{}{
	(Platform)(0),                 // 0: processing.v1alpha1.Platform
	(DocStatus)(0),                // 1: processing.v1alpha1.DocStatus
	(*DocMetadata)(nil),           // 2: processing.v1alpha1.DocMetadata
	(*Document)(nil),              // 3: processing.v1alpha1.Document
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_processing_v1alpha1_document_proto_depIdxs = []int32{
	0, // 0: processing.v1alpha1.Document.platform:type_name -> processing.v1alpha1.Platform
	1, // 1: processing.v1alpha1.Document.status:type_name -> processing.v1alpha1.DocStatus
	4, // 2: processing.v1alpha1.Document.created_at:type_name -> google.protobuf.Timestamp
	4, // 3: processing.v1alpha1.Document.updated_at:type_name -> google.protobuf.Timestamp
	2, // 4: processing.v1alpha1.Document.metadata:type_name -> processing.v1alpha1.DocMetadata
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_processing_v1alpha1_document_proto_init() }
func file_processing_v1alpha1_document_proto_init() {
	if File_processing_v1alpha1_document_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_processing_v1alpha1_document_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocMetadata); i {
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
		file_processing_v1alpha1_document_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_processing_v1alpha1_document_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_processing_v1alpha1_document_proto_goTypes,
		DependencyIndexes: file_processing_v1alpha1_document_proto_depIdxs,
		EnumInfos:         file_processing_v1alpha1_document_proto_enumTypes,
		MessageInfos:      file_processing_v1alpha1_document_proto_msgTypes,
	}.Build()
	File_processing_v1alpha1_document_proto = out.File
	file_processing_v1alpha1_document_proto_rawDesc = nil
	file_processing_v1alpha1_document_proto_goTypes = nil
	file_processing_v1alpha1_document_proto_depIdxs = nil
}