/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "processing.v1alpha1";

export interface Document {
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
  published: string;
  language: string;
}

function createBaseDocument(): Document {
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
    published: "",
    language: "",
  };
}

export const Document = {
  encode(message: Document, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
    if (message.published !== "") {
      writer.uint32(90).string(message.published);
    }
    if (message.language !== "") {
      writer.uint32(98).string(message.language);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Document {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDocument();
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
          message.published = reader.string();
          break;
        case 12:
          message.language = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Document {
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
      published: isSet(object.published) ? String(object.published) : "",
      language: isSet(object.language) ? String(object.language) : "",
    };
  },

  toJSON(message: Document): unknown {
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
    message.published !== undefined && (obj.published = message.published);
    message.language !== undefined && (obj.language = message.language);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Document>, I>>(object: I): Document {
    const message = createBaseDocument();
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
    message.published = object.published ?? "";
    message.language = object.language ?? "";
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
