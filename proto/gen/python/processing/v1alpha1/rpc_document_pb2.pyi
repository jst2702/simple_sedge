"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import google.protobuf.descriptor
import google.protobuf.message
import processing.v1alpha1.document_pb2
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class IngestDocumentRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    PLATFORM_FIELD_NUMBER: builtins.int
    PAYLOAD_FIELD_NUMBER: builtins.int
    METADATA_FIELD_NUMBER: builtins.int
    platform: processing.v1alpha1.document_pb2.Platform.ValueType
    """which news platform"""
    payload: builtins.bytes
    """raw bytes containing the payload"""
    @property
    def metadata(self) -> processing.v1alpha1.document_pb2.DocMetadata:
        """specified metadata"""
    def __init__(
        self,
        *,
        platform: processing.v1alpha1.document_pb2.Platform.ValueType = ...,
        payload: builtins.bytes = ...,
        metadata: processing.v1alpha1.document_pb2.DocMetadata | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["metadata", b"metadata"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["metadata", b"metadata", "payload", b"payload", "platform", b"platform"]) -> None: ...

global___IngestDocumentRequest = IngestDocumentRequest

class IngestDocumentResponse(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ID_FIELD_NUMBER: builtins.int
    id: builtins.str
    """unique identifier for external reference"""
    def __init__(
        self,
        *,
        id: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["id", b"id"]) -> None: ...

global___IngestDocumentResponse = IngestDocumentResponse
