from __future__ import annotations
from processing.v1alpha1.rpc_document_pb2 import (
    IngestDocumentRequest, IngestDocumentResponse)
from google.protobuf.timestamp_pb2 import Timestamp
from google.protobuf.json_format import MessageToDict, ParseDict
from news_scraping.common import NewsDoc
from dotenv import load_dotenv
import os
import requests

class IngestClient:
    def __init__(self, server_addr: str, api_key: str) -> None:
        self.server_addr = server_addr
        self.api_key = api_key
        self._timestamp = Timestamp()
        self._endpoint = f"{self.server_addr}/processing/v1alpha1/ingest_document"

    def ingest_doc(self, doc: NewsDoc) -> IngestDocumentResponse:
        req = IngestDocumentRequest(
            guid=doc.guid,
            url=doc.url,
            site=doc.site,
            site_full=doc.site_full,
            site_section=doc.site_section,
            headline=doc.headline,
            title=doc.title,
            body=doc.body,
            ticker=doc.ticker,
            tickers=doc.tickers,
            published_at=self._timestamp.FromDatetime(doc.published_date),
            language=doc.language,
            api_key=self.api_key
        )
        dict_payload = MessageToDict(req)
        resp = requests.post(self._endpoint, json=dict_payload)
        # Add error handling
        return ParseDict(
            js_dict=resp.json(), 
            message=IngestDocumentResponse(), 
            ignore_unknown_fields=True
        )

    @classmethod
    def from_default(cls, path = '.env') -> IngestClient:
        load_dotenv(path)

        server_addir = os.getenv('FEED_HTTP_SERVER_ADDRESS')
        if not server_addir:
            raise ValueError("FEED_HTTP_SERVER_ADDRESS not set")
        
        api_key = os.getenv('API_KEY')
        if not api_key:
            raise ValueError("API_KEY not set")

        return cls(
            server_addr=server_addir, 
            api_key=api_key
        )