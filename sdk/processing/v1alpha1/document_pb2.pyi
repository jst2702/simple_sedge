"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import collections.abc
import google.protobuf.descriptor
import google.protobuf.internal.containers
import google.protobuf.message
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class Document(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    GUID_FIELD_NUMBER: builtins.int
    URL_FIELD_NUMBER: builtins.int
    SITE_FIELD_NUMBER: builtins.int
    SITE_FULL_FIELD_NUMBER: builtins.int
    SITE_SECTION_FIELD_NUMBER: builtins.int
    HEADLINE_FIELD_NUMBER: builtins.int
    TITLE_FIELD_NUMBER: builtins.int
    BODY_FIELD_NUMBER: builtins.int
    TICKER_FIELD_NUMBER: builtins.int
    TICKERS_FIELD_NUMBER: builtins.int
    PUBLISHED_FIELD_NUMBER: builtins.int
    LANGUAGE_FIELD_NUMBER: builtins.int
    guid: builtins.str
    url: builtins.str
    site: builtins.str
    site_full: builtins.str
    site_section: builtins.str
    headline: builtins.str
    title: builtins.str
    body: builtins.str
    ticker: builtins.str
    @property
    def tickers(self) -> google.protobuf.internal.containers.RepeatedScalarFieldContainer[builtins.str]: ...
    published: builtins.str
    language: builtins.str
    def __init__(
        self,
        *,
        guid: builtins.str = ...,
        url: builtins.str = ...,
        site: builtins.str = ...,
        site_full: builtins.str = ...,
        site_section: builtins.str = ...,
        headline: builtins.str = ...,
        title: builtins.str = ...,
        body: builtins.str = ...,
        ticker: builtins.str = ...,
        tickers: collections.abc.Iterable[builtins.str] | None = ...,
        published: builtins.str = ...,
        language: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["body", b"body", "guid", b"guid", "headline", b"headline", "language", b"language", "published", b"published", "site", b"site", "site_full", b"site_full", "site_section", b"site_section", "ticker", b"ticker", "tickers", b"tickers", "title", b"title", "url", b"url"]) -> None: ...

global___Document = Document
