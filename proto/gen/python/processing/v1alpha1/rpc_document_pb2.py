# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: processing/v1alpha1/rpc_document.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from processing.v1alpha1 import document_pb2 as processing_dot_v1alpha1_dot_document__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n&processing/v1alpha1/rpc_document.proto\x12\x13processing.v1alpha1\x1a\"processing/v1alpha1/document.proto\"\xa8\x01\n\x15IngestDocumentRequest\x12\x39\n\x08platform\x18\x01 \x01(\x0e\x32\x1d.processing.v1alpha1.PlatformR\x08platform\x12\x16\n\x07payload\x18\x02 \x01(\x0cR\x05input\x12<\n\x08metadata\x18\x03 \x01(\x0b\x32 .processing.v1alpha1.DocMetadataR\x08metadata\"(\n\x16IngestDocumentResponse\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02idB\x98\x01\n\x17\x63om.processing.v1alpha1B\x10RpcDocumentProtoP\x01\xa2\x02\x03PXX\xaa\x02\x13Processing.V1alpha1\xca\x02\x13Processing\\V1alpha1\xe2\x02\x1fProcessing\\V1alpha1\\GPBMetadata\xea\x02\x14Processing::V1alpha1b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'processing.v1alpha1.rpc_document_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\027com.processing.v1alpha1B\020RpcDocumentProtoP\001\242\002\003PXX\252\002\023Processing.V1alpha1\312\002\023Processing\\V1alpha1\342\002\037Processing\\V1alpha1\\GPBMetadata\352\002\024Processing::V1alpha1'
  _INGESTDOCUMENTREQUEST._serialized_start=100
  _INGESTDOCUMENTREQUEST._serialized_end=268
  _INGESTDOCUMENTRESPONSE._serialized_start=270
  _INGESTDOCUMENTRESPONSE._serialized_end=310
# @@protoc_insertion_point(module_scope)
