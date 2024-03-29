# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: apis/v1alpha1/service.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from apis.v1alpha1 import rpc_document_pb2 as apis_dot_v1alpha1_dot_rpc__document__pb2
from apis.v1alpha1 import rpc_user_pb2 as apis_dot_v1alpha1_dot_rpc__user__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1b\x61pis/v1alpha1/service.proto\x12\x04\x61pis\x1a apis/v1alpha1/rpc_document.proto\x1a\x1c\x61pis/v1alpha1/rpc_user.proto\x1a\x1cgoogle/api/annotations.proto2\x80\x05\n\x0bSimpleSedge\x12\x61\n\nCreateUser\x12\x17.apis.CreateUserRequest\x1a\x18.apis.CreateUserResponse\" \x82\xd3\xe4\x93\x02\x1a\"\x15/v1alpha1/create_user:\x01*\x12]\n\tLoginUser\x12\x16.apis.LoginUserRequest\x1a\x17.apis.LoginUserResponse\"\x1f\x82\xd3\xe4\x93\x02\x19\"\x14/v1alpha1/login_user:\x01*\x12\x61\n\nUpdateUser\x12\x17.apis.UpdateUserRequest\x1a\x18.apis.UpdateUserResponse\" \x82\xd3\xe4\x93\x02\x1a\"\x15/v1alpha1/update_user:\x01*\x12\x62\n\x0bVerifyEmail\x12\x18.apis.VerifyEmailRequest\x1a\x19.apis.VerifyEmailResponse\"\x1e\x82\xd3\xe4\x93\x02\x18\x12\x16/v1alpha1/verify_email\x12j\n\rListDocuments\x12\x1a.apis.ListDocumentsRequest\x1a\x1b.apis.ListDocumentsResponse\" \x82\xd3\xe4\x93\x02\x1a\x12\x18/v1alpha1/list_documents\x12|\n\x0eIngestDocument\x12\x1b.apis.IngestDocumentRequest\x1a\x1c.apis.IngestDocumentResponse\"/\x82\xd3\xe4\x93\x02)\"$/v1alpha1/processing/ingest_document:\x01*BH\n\x08\x63om.apisB\x0cServiceProtoP\x01\xa2\x02\x03\x41XX\xaa\x02\x04\x41pis\xca\x02\x04\x41pis\xe2\x02\x10\x41pis\\GPBMetadata\xea\x02\x04\x41pisb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'apis.v1alpha1.service_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\010com.apisB\014ServiceProtoP\001\242\002\003AXX\252\002\004Apis\312\002\004Apis\342\002\020Apis\\GPBMetadata\352\002\004Apis'
  _SIMPLESEDGE.methods_by_name['CreateUser']._options = None
  _SIMPLESEDGE.methods_by_name['CreateUser']._serialized_options = b'\202\323\344\223\002\032\"\025/v1alpha1/create_user:\001*'
  _SIMPLESEDGE.methods_by_name['LoginUser']._options = None
  _SIMPLESEDGE.methods_by_name['LoginUser']._serialized_options = b'\202\323\344\223\002\031\"\024/v1alpha1/login_user:\001*'
  _SIMPLESEDGE.methods_by_name['UpdateUser']._options = None
  _SIMPLESEDGE.methods_by_name['UpdateUser']._serialized_options = b'\202\323\344\223\002\032\"\025/v1alpha1/update_user:\001*'
  _SIMPLESEDGE.methods_by_name['VerifyEmail']._options = None
  _SIMPLESEDGE.methods_by_name['VerifyEmail']._serialized_options = b'\202\323\344\223\002\030\022\026/v1alpha1/verify_email'
  _SIMPLESEDGE.methods_by_name['ListDocuments']._options = None
  _SIMPLESEDGE.methods_by_name['ListDocuments']._serialized_options = b'\202\323\344\223\002\032\022\030/v1alpha1/list_documents'
  _SIMPLESEDGE.methods_by_name['IngestDocument']._options = None
  _SIMPLESEDGE.methods_by_name['IngestDocument']._serialized_options = b'\202\323\344\223\002)\"$/v1alpha1/processing/ingest_document:\001*'
  _SIMPLESEDGE._serialized_start=132
  _SIMPLESEDGE._serialized_end=772
# @@protoc_insertion_point(module_scope)
