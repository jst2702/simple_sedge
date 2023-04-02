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
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n&processing/v1alpha1/rpc_document.proto\x12\x13processing.v1alpha1\x1a\"processing/v1alpha1/document.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x8d\x03\n\x15IngestDocumentRequest\x12\x12\n\x04guid\x18\x01 \x01(\tR\x04guid\x12\x10\n\x03url\x18\x02 \x01(\tR\x03url\x12\x12\n\x04site\x18\x03 \x01(\tR\x04site\x12\x1c\n\tsite_full\x18\x04 \x01(\tR\tsite_full\x12\"\n\x0csite_section\x18\x05 \x01(\tR\x0csite_section\x12\x1a\n\x08headline\x18\x06 \x01(\tR\x08headline\x12\x14\n\x05title\x18\x07 \x01(\tR\x05title\x12\x12\n\x04\x62ody\x18\x08 \x01(\tR\x04\x62ody\x12\x1b\n\x06ticker\x18\t \x01(\tH\x00R\x06ticker\x88\x01\x01\x12\x17\n\x07tickers\x18\n \x03(\tR\x06ticker\x12;\n\x0cpublished_at\x18\x0b \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tpublished\x12\x1a\n\x08language\x18\x0c \x01(\tR\x08language\x12\x18\n\x07\x61pi_key\x18\r \x01(\tR\x07\x61pi_keyB\t\n\x07_ticker\"S\n\x16IngestDocumentResponse\x12\x39\n\x08\x64ocument\x18\x01 \x01(\x0b\x32\x1d.processing.v1alpha1.DocumentR\x08\x64ocument\"J\n\x14ListDocumentsRequest\x12\x16\n\x06pageID\x18\x01 \x01(\x05R\x06pageID\x12\x1a\n\x08pageSize\x18\x02 \x01(\x05R\x08pageSize\"T\n\x15ListDocumentsResponse\x12;\n\tdocuments\x18\x01 \x03(\x0b\x32\x1d.processing.v1alpha1.DocumentR\tdocumentsB\x98\x01\n\x17\x63om.processing.v1alpha1B\x10RpcDocumentProtoP\x01\xa2\x02\x03PXX\xaa\x02\x13Processing.V1alpha1\xca\x02\x13Processing\\V1alpha1\xe2\x02\x1fProcessing\\V1alpha1\\GPBMetadata\xea\x02\x14Processing::V1alpha1b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'processing.v1alpha1.rpc_document_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\027com.processing.v1alpha1B\020RpcDocumentProtoP\001\242\002\003PXX\252\002\023Processing.V1alpha1\312\002\023Processing\\V1alpha1\342\002\037Processing\\V1alpha1\\GPBMetadata\352\002\024Processing::V1alpha1'
  _INGESTDOCUMENTREQUEST._serialized_start=133
  _INGESTDOCUMENTREQUEST._serialized_end=530
  _INGESTDOCUMENTRESPONSE._serialized_start=532
  _INGESTDOCUMENTRESPONSE._serialized_end=615
  _LISTDOCUMENTSREQUEST._serialized_start=617
  _LISTDOCUMENTSREQUEST._serialized_end=691
  _LISTDOCUMENTSRESPONSE._serialized_start=693
  _LISTDOCUMENTSRESPONSE._serialized_end=777
# @@protoc_insertion_point(module_scope)
