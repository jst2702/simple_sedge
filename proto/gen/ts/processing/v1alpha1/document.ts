/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Timestamp } from "../../google/protobuf/timestamp";

export const protobufPackage = "processing.v1alpha1";

export enum Platform {
  PLATFORM_UNDEFINED = 0,
  PLATFORM_BUSINESS_NEWSWIRE = 1,
  PLATFORM_GLOBAL_NEWSWIRE = 2,
  PLATFORM_PR_NEWSWIRE = 3,
  PLATFORM_ACCESSWIRE = 4,
  PLATFORM_MOTLEY_FOOL = 5,
  UNRECOGNIZED = -1,
}

export function platformFromJSON(object: any): Platform {
  switch (object) {
    case 0:
    case "PLATFORM_UNDEFINED":
      return Platform.PLATFORM_UNDEFINED;
    case 1:
    case "PLATFORM_BUSINESS_NEWSWIRE":
      return Platform.PLATFORM_BUSINESS_NEWSWIRE;
    case 2:
    case "PLATFORM_GLOBAL_NEWSWIRE":
      return Platform.PLATFORM_GLOBAL_NEWSWIRE;
    case 3:
    case "PLATFORM_PR_NEWSWIRE":
      return Platform.PLATFORM_PR_NEWSWIRE;
    case 4:
    case "PLATFORM_ACCESSWIRE":
      return Platform.PLATFORM_ACCESSWIRE;
    case 5:
    case "PLATFORM_MOTLEY_FOOL":
      return Platform.PLATFORM_MOTLEY_FOOL;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Platform.UNRECOGNIZED;
  }
}

export function platformToJSON(object: Platform): string {
  switch (object) {
    case Platform.PLATFORM_UNDEFINED:
      return "PLATFORM_UNDEFINED";
    case Platform.PLATFORM_BUSINESS_NEWSWIRE:
      return "PLATFORM_BUSINESS_NEWSWIRE";
    case Platform.PLATFORM_GLOBAL_NEWSWIRE:
      return "PLATFORM_GLOBAL_NEWSWIRE";
    case Platform.PLATFORM_PR_NEWSWIRE:
      return "PLATFORM_PR_NEWSWIRE";
    case Platform.PLATFORM_ACCESSWIRE:
      return "PLATFORM_ACCESSWIRE";
    case Platform.PLATFORM_MOTLEY_FOOL:
      return "PLATFORM_MOTLEY_FOOL";
    case Platform.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum DocStatus {
  STATUS_UNDEFINED = 0,
  STATUS_SUBMITTED = 1,
  STATUS_FAULTED = 2,
  STATUS_PROCESSED = 3,
  STATUS_COMPLETED = 4,
  UNRECOGNIZED = -1,
}

export function docStatusFromJSON(object: any): DocStatus {
  switch (object) {
    case 0:
    case "STATUS_UNDEFINED":
      return DocStatus.STATUS_UNDEFINED;
    case 1:
    case "STATUS_SUBMITTED":
      return DocStatus.STATUS_SUBMITTED;
    case 2:
    case "STATUS_FAULTED":
      return DocStatus.STATUS_FAULTED;
    case 3:
    case "STATUS_PROCESSED":
      return DocStatus.STATUS_PROCESSED;
    case 4:
    case "STATUS_COMPLETED":
      return DocStatus.STATUS_COMPLETED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return DocStatus.UNRECOGNIZED;
  }
}

export function docStatusToJSON(object: DocStatus): string {
  switch (object) {
    case DocStatus.STATUS_UNDEFINED:
      return "STATUS_UNDEFINED";
    case DocStatus.STATUS_SUBMITTED:
      return "STATUS_SUBMITTED";
    case DocStatus.STATUS_FAULTED:
      return "STATUS_FAULTED";
    case DocStatus.STATUS_PROCESSED:
      return "STATUS_PROCESSED";
    case DocStatus.STATUS_COMPLETED:
      return "STATUS_COMPLETED";
    case DocStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface DocMetadata {
  uuid: string;
  url: string;
  site: string;
  siteFull: string;
  siteSection: string;
  headline: string;
  title: string;
  body: string;
  published: string;
  language: string;
}

export interface Document {
  id: string;
  platform: Platform;
  size: number;
  /** indicates if the document has been succesfully processed */
  status: DocStatus;
  text: string;
  createdAt: string | undefined;
  updatedAt: string | undefined;
  metadata: DocMetadata | undefined;
}

function createBaseDocMetadata(): DocMetadata {
  return {
    uuid: "",
    url: "",
    site: "",
    siteFull: "",
    siteSection: "",
    headline: "",
    title: "",
    body: "",
    published: "",
    language: "",
  };
}

export const DocMetadata = {
  encode(message: DocMetadata, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.uuid !== "") {
      writer.uint32(10).string(message.uuid);
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
    if (message.published !== "") {
      writer.uint32(74).string(message.published);
    }
    if (message.language !== "") {
      writer.uint32(82).string(message.language);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DocMetadata {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDocMetadata();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uuid = reader.string();
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
          message.published = reader.string();
          break;
        case 10:
          message.language = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DocMetadata {
    return {
      uuid: isSet(object.encounter_id) ? String(object.encounter_id) : "",
      url: isSet(object.url) ? String(object.url) : "",
      site: isSet(object.site) ? String(object.site) : "",
      siteFull: isSet(object.site_full) ? String(object.site_full) : "",
      siteSection: isSet(object.site_section) ? String(object.site_section) : "",
      headline: isSet(object.headline) ? String(object.headline) : "",
      title: isSet(object.title) ? String(object.title) : "",
      body: isSet(object.body) ? String(object.body) : "",
      published: isSet(object.published) ? String(object.published) : "",
      language: isSet(object.language) ? String(object.language) : "",
    };
  },

  toJSON(message: DocMetadata): unknown {
    const obj: any = {};
    message.uuid !== undefined && (obj.encounter_id = message.uuid);
    message.url !== undefined && (obj.url = message.url);
    message.site !== undefined && (obj.site = message.site);
    message.siteFull !== undefined && (obj.site_full = message.siteFull);
    message.siteSection !== undefined && (obj.site_section = message.siteSection);
    message.headline !== undefined && (obj.headline = message.headline);
    message.title !== undefined && (obj.title = message.title);
    message.body !== undefined && (obj.body = message.body);
    message.published !== undefined && (obj.published = message.published);
    message.language !== undefined && (obj.language = message.language);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DocMetadata>, I>>(object: I): DocMetadata {
    const message = createBaseDocMetadata();
    message.uuid = object.uuid ?? "";
    message.url = object.url ?? "";
    message.site = object.site ?? "";
    message.siteFull = object.siteFull ?? "";
    message.siteSection = object.siteSection ?? "";
    message.headline = object.headline ?? "";
    message.title = object.title ?? "";
    message.body = object.body ?? "";
    message.published = object.published ?? "";
    message.language = object.language ?? "";
    return message;
  },
};

function createBaseDocument(): Document {
  return {
    id: "",
    platform: 0,
    size: 0,
    status: 0,
    text: "",
    createdAt: undefined,
    updatedAt: undefined,
    metadata: undefined,
  };
}

export const Document = {
  encode(message: Document, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.platform !== 0) {
      writer.uint32(24).int32(message.platform);
    }
    if (message.size !== 0) {
      writer.uint32(32).int64(message.size);
    }
    if (message.status !== 0) {
      writer.uint32(40).int32(message.status);
    }
    if (message.text !== "") {
      writer.uint32(50).string(message.text);
    }
    if (message.createdAt !== undefined) {
      Timestamp.encode(toTimestamp(message.createdAt), writer.uint32(58).fork()).ldelim();
    }
    if (message.updatedAt !== undefined) {
      Timestamp.encode(toTimestamp(message.updatedAt), writer.uint32(66).fork()).ldelim();
    }
    if (message.metadata !== undefined) {
      DocMetadata.encode(message.metadata, writer.uint32(74).fork()).ldelim();
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
          message.id = reader.string();
          break;
        case 3:
          message.platform = reader.int32() as any;
          break;
        case 4:
          message.size = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.status = reader.int32() as any;
          break;
        case 6:
          message.text = reader.string();
          break;
        case 7:
          message.createdAt = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 8:
          message.updatedAt = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 9:
          message.metadata = DocMetadata.decode(reader, reader.uint32());
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
      id: isSet(object.id) ? String(object.id) : "",
      platform: isSet(object.platform) ? platformFromJSON(object.platform) : 0,
      size: isSet(object.size) ? Number(object.size) : 0,
      status: isSet(object.status) ? docStatusFromJSON(object.status) : 0,
      text: isSet(object.text) ? String(object.text) : "",
      createdAt: isSet(object.created_at) ? String(object.created_at) : undefined,
      updatedAt: isSet(object.updated_at) ? String(object.updated_at) : undefined,
      metadata: isSet(object.updated_at) ? DocMetadata.fromJSON(object.updated_at) : undefined,
    };
  },

  toJSON(message: Document): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.platform !== undefined && (obj.platform = platformToJSON(message.platform));
    message.size !== undefined && (obj.size = Math.round(message.size));
    message.status !== undefined && (obj.status = docStatusToJSON(message.status));
    message.text !== undefined && (obj.text = message.text);
    message.createdAt !== undefined && (obj.created_at = message.createdAt);
    message.updatedAt !== undefined && (obj.updated_at = message.updatedAt);
    message.metadata !== undefined &&
      (obj.updated_at = message.metadata ? DocMetadata.toJSON(message.metadata) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Document>, I>>(object: I): Document {
    const message = createBaseDocument();
    message.id = object.id ?? "";
    message.platform = object.platform ?? 0;
    message.size = object.size ?? 0;
    message.status = object.status ?? 0;
    message.text = object.text ?? "";
    message.createdAt = object.createdAt ?? undefined;
    message.updatedAt = object.updatedAt ?? undefined;
    message.metadata = (object.metadata !== undefined && object.metadata !== null)
      ? DocMetadata.fromPartial(object.metadata)
      : undefined;
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var tsProtoGlobalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new tsProtoGlobalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
