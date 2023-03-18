"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import google.protobuf.descriptor
import google.protobuf.internal.enum_type_wrapper
import google.protobuf.message
import google.protobuf.timestamp_pb2
import sys
import typing

if sys.version_info >= (3, 10):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class _Platform:
    ValueType = typing.NewType("ValueType", builtins.int)
    V: typing_extensions.TypeAlias = ValueType

class _PlatformEnumTypeWrapper(google.protobuf.internal.enum_type_wrapper._EnumTypeWrapper[_Platform.ValueType], builtins.type):  # noqa: F821
    DESCRIPTOR: google.protobuf.descriptor.EnumDescriptor
    PLATFORM_UNDEFINED: _Platform.ValueType  # 0
    PLATFORM_BUSINESS_NEWSWIRE: _Platform.ValueType  # 1
    PLATFORM_GLOBAL_NEWSWIRE: _Platform.ValueType  # 2
    PLATFORM_PR_NEWSWIRE: _Platform.ValueType  # 3
    PLATFORM_ACCESSWIRE: _Platform.ValueType  # 4
    PLATFORM_MOTLEY_FOOL: _Platform.ValueType  # 5

class Platform(_Platform, metaclass=_PlatformEnumTypeWrapper): ...

PLATFORM_UNDEFINED: Platform.ValueType  # 0
PLATFORM_BUSINESS_NEWSWIRE: Platform.ValueType  # 1
PLATFORM_GLOBAL_NEWSWIRE: Platform.ValueType  # 2
PLATFORM_PR_NEWSWIRE: Platform.ValueType  # 3
PLATFORM_ACCESSWIRE: Platform.ValueType  # 4
PLATFORM_MOTLEY_FOOL: Platform.ValueType  # 5
global___Platform = Platform

class _DocStatus:
    ValueType = typing.NewType("ValueType", builtins.int)
    V: typing_extensions.TypeAlias = ValueType

class _DocStatusEnumTypeWrapper(google.protobuf.internal.enum_type_wrapper._EnumTypeWrapper[_DocStatus.ValueType], builtins.type):  # noqa: F821
    DESCRIPTOR: google.protobuf.descriptor.EnumDescriptor
    STATUS_UNDEFINED: _DocStatus.ValueType  # 0
    STATUS_SUBMITTED: _DocStatus.ValueType  # 1
    STATUS_FAULTED: _DocStatus.ValueType  # 2
    STATUS_PROCESSED: _DocStatus.ValueType  # 3
    STATUS_COMPLETED: _DocStatus.ValueType  # 4

class DocStatus(_DocStatus, metaclass=_DocStatusEnumTypeWrapper): ...

STATUS_UNDEFINED: DocStatus.ValueType  # 0
STATUS_SUBMITTED: DocStatus.ValueType  # 1
STATUS_FAULTED: DocStatus.ValueType  # 2
STATUS_PROCESSED: DocStatus.ValueType  # 3
STATUS_COMPLETED: DocStatus.ValueType  # 4
global___DocStatus = DocStatus

class DocMetadata(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    UUID_FIELD_NUMBER: builtins.int
    URL_FIELD_NUMBER: builtins.int
    SITE_FIELD_NUMBER: builtins.int
    SITE_FULL_FIELD_NUMBER: builtins.int
    SITE_SECTION_FIELD_NUMBER: builtins.int
    HEADLINE_FIELD_NUMBER: builtins.int
    TITLE_FIELD_NUMBER: builtins.int
    BODY_FIELD_NUMBER: builtins.int
    PUBLISHED_FIELD_NUMBER: builtins.int
    LANGUAGE_FIELD_NUMBER: builtins.int
    uuid: builtins.str
    url: builtins.str
    site: builtins.str
    site_full: builtins.str
    site_section: builtins.str
    headline: builtins.str
    title: builtins.str
    body: builtins.str
    published: builtins.str
    language: builtins.str
    def __init__(
        self,
        *,
        uuid: builtins.str = ...,
        url: builtins.str = ...,
        site: builtins.str = ...,
        site_full: builtins.str = ...,
        site_section: builtins.str = ...,
        headline: builtins.str = ...,
        title: builtins.str = ...,
        body: builtins.str = ...,
        published: builtins.str = ...,
        language: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["body", b"body", "headline", b"headline", "language", b"language", "published", b"published", "site", b"site", "site_full", b"site_full", "site_section", b"site_section", "title", b"title", "url", b"url", "uuid", b"uuid"]) -> None: ...

global___DocMetadata = DocMetadata

class Document(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ID_FIELD_NUMBER: builtins.int
    PLATFORM_FIELD_NUMBER: builtins.int
    SIZE_FIELD_NUMBER: builtins.int
    STATUS_FIELD_NUMBER: builtins.int
    TEXT_FIELD_NUMBER: builtins.int
    CREATED_AT_FIELD_NUMBER: builtins.int
    UPDATED_AT_FIELD_NUMBER: builtins.int
    METADATA_FIELD_NUMBER: builtins.int
    id: builtins.str
    platform: global___Platform.ValueType
    size: builtins.int
    status: global___DocStatus.ValueType
    """indicates if the document has been succesfully processed"""
    text: builtins.str
    @property
    def created_at(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
    @property
    def updated_at(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
    @property
    def metadata(self) -> global___DocMetadata: ...
    def __init__(
        self,
        *,
        id: builtins.str = ...,
        platform: global___Platform.ValueType = ...,
        size: builtins.int = ...,
        status: global___DocStatus.ValueType = ...,
        text: builtins.str = ...,
        created_at: google.protobuf.timestamp_pb2.Timestamp | None = ...,
        updated_at: google.protobuf.timestamp_pb2.Timestamp | None = ...,
        metadata: global___DocMetadata | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["created_at", b"created_at", "metadata", b"metadata", "updated_at", b"updated_at"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["created_at", b"created_at", "id", b"id", "metadata", b"metadata", "platform", b"platform", "size", b"size", "status", b"status", "text", b"text", "updated_at", b"updated_at"]) -> None: ...

global___Document = Document