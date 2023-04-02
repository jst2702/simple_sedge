/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Timestamp } from "../../google/protobuf/timestamp";
import { Document } from "./document";

export const protobufPackage = "processing.v1alpha1";

export interface IngestDocumentRequest {
  guid: string;
  url: string;
  site: string;
  siteFull: string;
  siteSection: string;
  headline: string;
  title: string;
  body: string;
  ticker?: string | undefined;
  tickers: string[];
  publishedAt: string | undefined;
  language: string;
  apiKey: string;
}

export interface IngestDocumentResponse {
  document: Document | undefined;
}

export interface ListDocumentsRequest {
  pageID: number;
  pageSize: number;
}

export interface ListDocumentsResponse {
  documents: Document[];
}

function createBaseIngestDocumentRequest(): IngestDocumentRequest {
  return {
    guid: "",
    url: "",
    site: "",
    siteFull: "",
    siteSection: "",
    headline: "",
    title: "",
    body: "",
    ticker: undefined,
    tickers: [],
    publishedAt: undefined,
    language: "",
    apiKey: "",
  };
}

export const IngestDocumentRequest = {
  encode(message: IngestDocumentRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.guid !== "") {
      writer.uint32(10).string(message.guid);
    }
    if (message.url !== "") {
      writer.uint32(18).string(message.url);
    }
    if (message.site !== "") {
      writer.uint32(26).string(message.site);
    }
    if (message.siteFull !== "") {
      writer.uint32(34).string(message.siteFull);
    }
    if (message.siteSection !== "") {
      writer.uint32(42).string(message.siteSection);
    }
    if (message.headline !== "") {
      writer.uint32(50).string(message.headline);
    }
    if (message.title !== "") {
      writer.uint32(58).string(message.title);
    }
    if (message.body !== "") {
      writer.uint32(66).string(message.body);
    }
    if (message.ticker !== undefined) {
      writer.uint32(74).string(message.ticker);
    }
    for (const v of message.tickers) {
      writer.uint32(82).string(v!);
    }
    if (message.publishedAt !== undefined) {
      Timestamp.encode(toTimestamp(message.publishedAt), writer.uint32(90).fork()).ldelim();
    }
    if (message.language !== "") {
      writer.uint32(98).string(message.language);
    }
    if (message.apiKey !== "") {
      writer.uint32(106).string(message.apiKey);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IngestDocumentRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIngestDocumentRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.guid = reader.string();
          break;
        case 2:
          message.url = reader.string();
          break;
        case 3:
          message.site = reader.string();
          break;
        case 4:
          message.siteFull = reader.string();
          break;
        case 5:
          message.siteSection = reader.string();
          break;
        case 6:
          message.headline = reader.string();
          break;
        case 7:
          message.title = reader.string();
          break;
        case 8:
          message.body = reader.string();
          break;
        case 9:
          message.ticker = reader.string();
          break;
        case 10:
          message.tickers.push(reader.string());
          break;
        case 11:
          message.publishedAt = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 12:
          message.language = reader.string();
          break;
        case 13:
          message.apiKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): IngestDocumentRequest {
    return {
      guid: isSet(object.guid) ? String(object.guid) : "",
      url: isSet(object.url) ? String(object.url) : "",
      site: isSet(object.site) ? String(object.site) : "",
      siteFull: isSet(object.site_full) ? String(object.site_full) : "",
      siteSection: isSet(object.site_section) ? String(object.site_section) : "",
      headline: isSet(object.headline) ? String(object.headline) : "",
      title: isSet(object.title) ? String(object.title) : "",
      body: isSet(object.body) ? String(object.body) : "",
      ticker: isSet(object.ticker) ? String(object.ticker) : undefined,
      tickers: Array.isArray(object?.ticker) ? object.ticker.map((e: any) => String(e)) : [],
      publishedAt: isSet(object.published) ? String(object.published) : undefined,
      language: isSet(object.language) ? String(object.language) : "",
      apiKey: isSet(object.api_key) ? String(object.api_key) : "",
    };
  },

  toJSON(message: IngestDocumentRequest): unknown {
    const obj: any = {};
    message.guid !== undefined && (obj.guid = message.guid);
    message.url !== undefined && (obj.url = message.url);
    message.site !== undefined && (obj.site = message.site);
    message.siteFull !== undefined && (obj.site_full = message.siteFull);
    message.siteSection !== undefined && (obj.site_section = message.siteSection);
    message.headline !== undefined && (obj.headline = message.headline);
    message.title !== undefined && (obj.title = message.title);
    message.body !== undefined && (obj.body = message.body);
    message.ticker !== undefined && (obj.ticker = message.ticker);
    if (message.tickers) {
      obj.ticker = message.tickers.map((e) => e);
    } else {
      obj.ticker = [];
    }
    message.publishedAt !== undefined && (obj.published = message.publishedAt);
    message.language !== undefined && (obj.language = message.language);
    message.apiKey !== undefined && (obj.api_key = message.apiKey);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<IngestDocumentRequest>, I>>(object: I): IngestDocumentRequest {
    const message = createBaseIngestDocumentRequest();
    message.guid = object.guid ?? "";
    message.url = object.url ?? "";
    message.site = object.site ?? "";
    message.siteFull = object.siteFull ?? "";
    message.siteSection = object.siteSection ?? "";
    message.headline = object.headline ?? "";
    message.title = object.title ?? "";
    message.body = object.body ?? "";
    message.ticker = object.ticker ?? undefined;
    message.tickers = object.tickers?.map((e) => e) || [];
    message.publishedAt = object.publishedAt ?? undefined;
    message.language = object.language ?? "";
    message.apiKey = object.apiKey ?? "";
    return message;
  },
};

function createBaseIngestDocumentResponse(): IngestDocumentResponse {
  return { document: undefined };
}

export const IngestDocumentResponse = {
  encode(message: IngestDocumentResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.document !== undefined) {
      Document.encode(message.document, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IngestDocumentResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIngestDocumentResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.document = Document.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): IngestDocumentResponse {
    return { document: isSet(object.document) ? Document.fromJSON(object.document) : undefined };
  },

  toJSON(message: IngestDocumentResponse): unknown {
    const obj: any = {};
    message.document !== undefined && (obj.document = message.document ? Document.toJSON(message.document) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<IngestDocumentResponse>, I>>(object: I): IngestDocumentResponse {
    const message = createBaseIngestDocumentResponse();
    message.document = (object.document !== undefined && object.document !== null)
      ? Document.fromPartial(object.document)
      : undefined;
    return message;
  },
};

function createBaseListDocumentsRequest(): ListDocumentsRequest {
  return { pageID: 0, pageSize: 0 };
}

export const ListDocumentsRequest = {
  encode(message: ListDocumentsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pageID !== 0) {
      writer.uint32(8).int32(message.pageID);
    }
    if (message.pageSize !== 0) {
      writer.uint32(16).int32(message.pageSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListDocumentsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListDocumentsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pageID = reader.int32();
          break;
        case 2:
          message.pageSize = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListDocumentsRequest {
    return {
      pageID: isSet(object.pageID) ? Number(object.pageID) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
    };
  },

  toJSON(message: ListDocumentsRequest): unknown {
    const obj: any = {};
    message.pageID !== undefined && (obj.pageID = Math.round(message.pageID));
    message.pageSize !== undefined && (obj.pageSize = Math.round(message.pageSize));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListDocumentsRequest>, I>>(object: I): ListDocumentsRequest {
    const message = createBaseListDocumentsRequest();
    message.pageID = object.pageID ?? 0;
    message.pageSize = object.pageSize ?? 0;
    return message;
  },
};

function createBaseListDocumentsResponse(): ListDocumentsResponse {
  return { documents: [] };
}

export const ListDocumentsResponse = {
  encode(message: ListDocumentsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.documents) {
      Document.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListDocumentsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListDocumentsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.documents.push(Document.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListDocumentsResponse {
    return {
      documents: Array.isArray(object?.documents) ? object.documents.map((e: any) => Document.fromJSON(e)) : [],
    };
  },

  toJSON(message: ListDocumentsResponse): unknown {
    const obj: any = {};
    if (message.documents) {
      obj.documents = message.documents.map((e) => e ? Document.toJSON(e) : undefined);
    } else {
      obj.documents = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListDocumentsResponse>, I>>(object: I): ListDocumentsResponse {
    const message = createBaseListDocumentsResponse();
    message.documents = object.documents?.map((e) => Document.fromPartial(e)) || [];
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function toTimestamp(dateStr: string): Timestamp {
  const date = new Date(dateStr);
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): string {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis).toISOString();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
