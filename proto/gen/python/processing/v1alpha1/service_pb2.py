# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: processing/v1alpha1/service.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from processing.v1alpha1 import document_pb2 as processing_dot_v1alpha1_dot_document__pb2
from processing.v1alpha1 import rpc_document_pb2 as processing_dot_v1alpha1_dot_rpc__document__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n!processing/v1alpha1/service.proto\x12\x13processing.v1alpha1\x1a\"processing/v1alpha1/document.proto\x1a&processing/v1alpha1/rpc_document.proto\x1a\x1cgoogle/api/annotations.proto2\xb0\x01\n\x10IngestionService\x12\x9b\x01\n\x0eIngestDocument\x12*.processing.v1alpha1.IngestDocumentRequest\x1a+.processing.v1alpha1.IngestDocumentResponse\"0\x82\xd3\xe4\x93\x02*\"%/processing/v2alpha1/documents/ingest:\x01*B\x94\x01\n\x17\x63om.processing.v1alpha1B\x0cServiceProtoP\x01\xa2\x02\x03PXX\xaa\x02\x13Processing.V1alpha1\xca\x02\x13Processing\\V1alpha1\xe2\x02\x1fProcessing\\V1alpha1\\GPBMetadata\xea\x02\x14Processing::V1alpha1b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'processing.v1alpha1.service_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\027com.processing.v1alpha1B\014ServiceProtoP\001\242\002\003PXX\252\002\023Processing.V1alpha1\312\002\023Processing\\V1alpha1\342\002\037Processing\\V1alpha1\\GPBMetadata\352\002\024Processing::V1alpha1'
  _INGESTIONSERVICE.methods_by_name['IngestDocument']._options = None
  _INGESTIONSERVICE.methods_by_name['IngestDocument']._serialized_options = b'\202\323\344\223\002*\"%/processing/v2alpha1/documents/ingest:\001*'
  _INGESTIONSERVICE._serialized_start=165
  _INGESTIONSERVICE._serialized_end=341
# @@protoc_insertion_point(module_scope)
