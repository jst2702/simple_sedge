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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n&processing/v1alpha1/rpc_document.proto\x12\nprocessing\x1a\"processing/v1alpha1/document.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x8e\x03\n\x15IngestDocumentRequest\x12\x12\n\x04guid\x18\x01 \x01(\tR\x04guid\x12\x10\n\x03url\x18\x02 \x01(\tR\x03url\x12\x12\n\x04site\x18\x03 \x01(\tR\x04site\x12\x1c\n\tsite_full\x18\x04 \x01(\tR\tsite_full\x12\"\n\x0csite_section\x18\x05 \x01(\tR\x0csite_section\x12\x1a\n\x08headline\x18\x06 \x01(\tR\x08headline\x12\x14\n\x05title\x18\x07 \x01(\tR\x05title\x12\x12\n\x04\x62ody\x18\x08 \x01(\tR\x04\x62ody\x12\x1b\n\x06ticker\x18\t \x01(\tH\x00R\x06ticker\x88\x01\x01\x12\x18\n\x07tickers\x18\n \x03(\tR\x07tickers\x12;\n\x0cpublished_at\x18\x0b \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tpublished\x12\x1a\n\x08language\x18\x0c \x01(\tR\x08language\x12\x18\n\x07\x61pi_key\x18\r \x01(\tR\x07\x61pi_keyB\t\n\x07_ticker\"J\n\x16IngestDocumentResponse\x12\x30\n\x08\x64ocument\x18\x01 \x01(\x0b\x32\x14.processing.DocumentR\x08\x64ocument\"J\n\x14ListDocumentsRequest\x12\x16\n\x06pageID\x18\x01 \x01(\x05R\x06pageID\x12\x1a\n\x08pageSize\x18\x02 \x01(\x05R\x08pageSize\"K\n\x15ListDocumentsResponse\x12\x32\n\tdocuments\x18\x01 \x03(\x0b\x32\x14.processing.DocumentR\tdocumentsBj\n\x0e\x63om.processingB\x10RpcDocumentProtoP\x01\xa2\x02\x03PXX\xaa\x02\nProcessing\xca\x02\nProcessing\xe2\x02\x16Processing\\GPBMetadata\xea\x02\nProcessingb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'processing.v1alpha1.rpc_document_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\016com.processingB\020RpcDocumentProtoP\001\242\002\003PXX\252\002\nProcessing\312\002\nProcessing\342\002\026Processing\\GPBMetadata\352\002\nProcessing'
  _INGESTDOCUMENTREQUEST._serialized_start=124
  _INGESTDOCUMENTREQUEST._serialized_end=522
  _INGESTDOCUMENTRESPONSE._serialized_start=524
  _INGESTDOCUMENTRESPONSE._serialized_end=598
  _LISTDOCUMENTSREQUEST._serialized_start=600
  _LISTDOCUMENTSREQUEST._serialized_end=674
  _LISTDOCUMENTSRESPONSE._serialized_start=676
  _LISTDOCUMENTSRESPONSE._serialized_end=751
# @@protoc_insertion_point(module_scope)
