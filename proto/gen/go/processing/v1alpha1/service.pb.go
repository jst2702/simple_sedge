// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: processing/v1alpha1/service.proto

package v1alpha1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_processing_v1alpha1_service_proto protoreflect.FileDescriptor

var file_processing_v1alpha1_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x1a,
	0x26, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x95, 0x02, 0x0a, 0x10, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x0e, 0x49,
	0x6e, 0x67, 0x65, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x67, 0x65, 0x73,
	0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e,
	0x67, 0x65, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x22, 0x24, 0x2f, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x76, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x98, 0x01,
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x30, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0xca, 0x02, 0x0a, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0xe2, 0x02, 0x16, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_processing_v1alpha1_service_proto_goTypes = []interface{}{
	(*IngestDocumentRequest)(nil),  // 0: processing.IngestDocumentRequest
	(*ListDocumentsRequest)(nil),   // 1: processing.ListDocumentsRequest
	(*IngestDocumentResponse)(nil), // 2: processing.IngestDocumentResponse
	(*ListDocumentsResponse)(nil),  // 3: processing.ListDocumentsResponse
}
var file_processing_v1alpha1_service_proto_depIdxs = []int32{
	0, // 0: processing.IngestionService.IngestDocument:input_type -> processing.IngestDocumentRequest
	1, // 1: processing.IngestionService.ListDocuments:input_type -> processing.ListDocumentsRequest
	2, // 2: processing.IngestionService.IngestDocument:output_type -> processing.IngestDocumentResponse
	3, // 3: processing.IngestionService.ListDocuments:output_type -> processing.ListDocumentsResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_processing_v1alpha1_service_proto_init() }
func file_processing_v1alpha1_service_proto_init() {
	if File_processing_v1alpha1_service_proto != nil {
		return
	}
	file_processing_v1alpha1_rpc_document_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_processing_v1alpha1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_processing_v1alpha1_service_proto_goTypes,
		DependencyIndexes: file_processing_v1alpha1_service_proto_depIdxs,
	}.Build()
	File_processing_v1alpha1_service_proto = out.File
	file_processing_v1alpha1_service_proto_rawDesc = nil
	file_processing_v1alpha1_service_proto_goTypes = nil
	file_processing_v1alpha1_service_proto_depIdxs = nil
}
