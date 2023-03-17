/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { DocMetadata, Platform, platformFromJSON, platformToJSON } from "./document";

export const protobufPackage = "processing.v1alpha1";

export interface IngestDocumentRequest {
  /** which news platform */
  platform: Platform;
  /** raw bytes containing the payload */
  payload: Uint8Array;
  /** specified metadata */
  metadata: DocMetadata | undefined;
}

export interface IngestDocumentResponse {
  /** unique identifier for external reference */
  id: string;
}

function createBaseIngestDocumentRequest(): IngestDocumentRequest {
  return { platform: 0, payload: new Uint8Array(), metadata: undefined };
}

export const IngestDocumentRequest = {
  encode(message: IngestDocumentRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.platform !== 0) {
      writer.uint32(8).int32(message.platform);
    }
    if (message.payload.length !== 0) {
      writer.uint32(18).bytes(message.payload);
    }
    if (message.metadata !== undefined) {
      DocMetadata.encode(message.metadata, writer.uint32(26).fork()).ldelim();
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
          message.platform = reader.int32() as any;
          break;
        case 2:
          message.payload = reader.bytes();
          break;
        case 3:
          message.metadata = DocMetadata.decode(reader, reader.uint32());
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
      platform: isSet(object.platform) ? platformFromJSON(object.platform) : 0,
      payload: isSet(object.input) ? bytesFromBase64(object.input) : new Uint8Array(),
      metadata: isSet(object.metadata) ? DocMetadata.fromJSON(object.metadata) : undefined,
    };
  },

  toJSON(message: IngestDocumentRequest): unknown {
    const obj: any = {};
    message.platform !== undefined && (obj.platform = platformToJSON(message.platform));
    message.payload !== undefined &&
      (obj.input = base64FromBytes(message.payload !== undefined ? message.payload : new Uint8Array()));
    message.metadata !== undefined &&
      (obj.metadata = message.metadata ? DocMetadata.toJSON(message.metadata) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<IngestDocumentRequest>, I>>(object: I): IngestDocumentRequest {
    const message = createBaseIngestDocumentRequest();
    message.platform = object.platform ?? 0;
    message.payload = object.payload ?? new Uint8Array();
    message.metadata = (object.metadata !== undefined && object.metadata !== null)
      ? DocMetadata.fromPartial(object.metadata)
      : undefined;
    return message;
  },
};

function createBaseIngestDocumentResponse(): IngestDocumentResponse {
  return { id: "" };
}

export const IngestDocumentResponse = {
  encode(message: IngestDocumentResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
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
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): IngestDocumentResponse {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: IngestDocumentResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<IngestDocumentResponse>, I>>(object: I): IngestDocumentResponse {
    const message = createBaseIngestDocumentResponse();
    message.id = object.id ?? "";
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

function bytesFromBase64(b64: string): Uint8Array {
  if (tsProtoGlobalThis.Buffer) {
    return Uint8Array.from(tsProtoGlobalThis.Buffer.from(b64, "base64"));
  } else {
    const bin = tsProtoGlobalThis.atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
      arr[i] = bin.charCodeAt(i);
    }
    return arr;
  }
}

function base64FromBytes(arr: Uint8Array): string {
  if (tsProtoGlobalThis.Buffer) {
    return tsProtoGlobalThis.Buffer.from(arr).toString("base64");
  } else {
    const bin: string[] = [];
    arr.forEach((byte) => {
      bin.push(String.fromCharCode(byte));
    });
    return tsProtoGlobalThis.btoa(bin.join(""));
  }
}

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
