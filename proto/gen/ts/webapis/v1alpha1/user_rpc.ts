/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Timestamp } from "../../google/protobuf/timestamp";
import { User } from "./user";

export const protobufPackage = "processing.v1alpha1";

export interface CreateUserRequest {
  email: string;
  username: string;
  password: string;
}

export interface CreateUserResponse {
  user: User | undefined;
}

export interface LoginUserRequest {
  username: string;
  password: string;
}

export interface LoginUserResponse {
  user: User | undefined;
  sessionId: string;
  accessToken: string;
  refreshToken: string;
  accessTokenExpiresAt: string | undefined;
  refreshTokenExpiresAt: string | undefined;
}

export interface UpdateUserRequest {
  username: string;
  email?: string | undefined;
  password?: string | undefined;
}

export interface UpdateUserResponse {
  user: User | undefined;
}

function createBaseCreateUserRequest(): CreateUserRequest {
  return { email: "", username: "", password: "" };
}

export const CreateUserRequest = {
  encode(message: CreateUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.email !== "") {
      writer.uint32(10).string(message.email);
    }
    if (message.username !== "") {
      writer.uint32(18).string(message.username);
    }
    if (message.password !== "") {
      writer.uint32(26).string(message.password);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateUserRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.email = reader.string();
          break;
        case 2:
          message.username = reader.string();
          break;
        case 3:
          message.password = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateUserRequest {
    return {
      email: isSet(object.email) ? String(object.email) : "",
      username: isSet(object.username) ? String(object.username) : "",
      password: isSet(object.password) ? String(object.password) : "",
    };
  },

  toJSON(message: CreateUserRequest): unknown {
    const obj: any = {};
    message.email !== undefined && (obj.email = message.email);
    message.username !== undefined && (obj.username = message.username);
    message.password !== undefined && (obj.password = message.password);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateUserRequest>, I>>(object: I): CreateUserRequest {
    const message = createBaseCreateUserRequest();
    message.email = object.email ?? "";
    message.username = object.username ?? "";
    message.password = object.password ?? "";
    return message;
  },
};

function createBaseCreateUserResponse(): CreateUserResponse {
  return { user: undefined };
}

export const CreateUserResponse = {
  encode(message: CreateUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateUserResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user = User.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateUserResponse {
    return { user: isSet(object.user) ? User.fromJSON(object.user) : undefined };
  },

  toJSON(message: CreateUserResponse): unknown {
    const obj: any = {};
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateUserResponse>, I>>(object: I): CreateUserResponse {
    const message = createBaseCreateUserResponse();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    return message;
  },
};

function createBaseLoginUserRequest(): LoginUserRequest {
  return { username: "", password: "" };
}

export const LoginUserRequest = {
  encode(message: LoginUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.username !== "") {
      writer.uint32(10).string(message.username);
    }
    if (message.password !== "") {
      writer.uint32(18).string(message.password);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LoginUserRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLoginUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.username = reader.string();
          break;
        case 2:
          message.password = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LoginUserRequest {
    return {
      username: isSet(object.username) ? String(object.username) : "",
      password: isSet(object.password) ? String(object.password) : "",
    };
  },

  toJSON(message: LoginUserRequest): unknown {
    const obj: any = {};
    message.username !== undefined && (obj.username = message.username);
    message.password !== undefined && (obj.password = message.password);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LoginUserRequest>, I>>(object: I): LoginUserRequest {
    const message = createBaseLoginUserRequest();
    message.username = object.username ?? "";
    message.password = object.password ?? "";
    return message;
  },
};

function createBaseLoginUserResponse(): LoginUserResponse {
  return {
    user: undefined,
    sessionId: "",
    accessToken: "",
    refreshToken: "",
    accessTokenExpiresAt: undefined,
    refreshTokenExpiresAt: undefined,
  };
}

export const LoginUserResponse = {
  encode(message: LoginUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    if (message.sessionId !== "") {
      writer.uint32(18).string(message.sessionId);
    }
    if (message.accessToken !== "") {
      writer.uint32(26).string(message.accessToken);
    }
    if (message.refreshToken !== "") {
      writer.uint32(34).string(message.refreshToken);
    }
    if (message.accessTokenExpiresAt !== undefined) {
      Timestamp.encode(toTimestamp(message.accessTokenExpiresAt), writer.uint32(42).fork()).ldelim();
    }
    if (message.refreshTokenExpiresAt !== undefined) {
      Timestamp.encode(toTimestamp(message.refreshTokenExpiresAt), writer.uint32(50).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LoginUserResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLoginUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user = User.decode(reader, reader.uint32());
          break;
        case 2:
          message.sessionId = reader.string();
          break;
        case 3:
          message.accessToken = reader.string();
          break;
        case 4:
          message.refreshToken = reader.string();
          break;
        case 5:
          message.accessTokenExpiresAt = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 6:
          message.refreshTokenExpiresAt = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LoginUserResponse {
    return {
      user: isSet(object.user) ? User.fromJSON(object.user) : undefined,
      sessionId: isSet(object.sessionId) ? String(object.sessionId) : "",
      accessToken: isSet(object.accessToken) ? String(object.accessToken) : "",
      refreshToken: isSet(object.refreshToken) ? String(object.refreshToken) : "",
      accessTokenExpiresAt: isSet(object.accessTokenExpiresAt) ? String(object.accessTokenExpiresAt) : undefined,
      refreshTokenExpiresAt: isSet(object.refreshTokenExpiresAt) ? String(object.refreshTokenExpiresAt) : undefined,
    };
  },

  toJSON(message: LoginUserResponse): unknown {
    const obj: any = {};
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    message.sessionId !== undefined && (obj.sessionId = message.sessionId);
    message.accessToken !== undefined && (obj.accessToken = message.accessToken);
    message.refreshToken !== undefined && (obj.refreshToken = message.refreshToken);
    message.accessTokenExpiresAt !== undefined && (obj.accessTokenExpiresAt = message.accessTokenExpiresAt);
    message.refreshTokenExpiresAt !== undefined && (obj.refreshTokenExpiresAt = message.refreshTokenExpiresAt);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LoginUserResponse>, I>>(object: I): LoginUserResponse {
    const message = createBaseLoginUserResponse();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    message.sessionId = object.sessionId ?? "";
    message.accessToken = object.accessToken ?? "";
    message.refreshToken = object.refreshToken ?? "";
    message.accessTokenExpiresAt = object.accessTokenExpiresAt ?? undefined;
    message.refreshTokenExpiresAt = object.refreshTokenExpiresAt ?? undefined;
    return message;
  },
};

function createBaseUpdateUserRequest(): UpdateUserRequest {
  return { username: "", email: undefined, password: undefined };
}

export const UpdateUserRequest = {
  encode(message: UpdateUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.username !== "") {
      writer.uint32(10).string(message.username);
    }
    if (message.email !== undefined) {
      writer.uint32(18).string(message.email);
    }
    if (message.password !== undefined) {
      writer.uint32(26).string(message.password);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateUserRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.username = reader.string();
          break;
        case 2:
          message.email = reader.string();
          break;
        case 3:
          message.password = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateUserRequest {
    return {
      username: isSet(object.username) ? String(object.username) : "",
      email: isSet(object.email) ? String(object.email) : undefined,
      password: isSet(object.password) ? String(object.password) : undefined,
    };
  },

  toJSON(message: UpdateUserRequest): unknown {
    const obj: any = {};
    message.username !== undefined && (obj.username = message.username);
    message.email !== undefined && (obj.email = message.email);
    message.password !== undefined && (obj.password = message.password);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateUserRequest>, I>>(object: I): UpdateUserRequest {
    const message = createBaseUpdateUserRequest();
    message.username = object.username ?? "";
    message.email = object.email ?? undefined;
    message.password = object.password ?? undefined;
    return message;
  },
};

function createBaseUpdateUserResponse(): UpdateUserResponse {
  return { user: undefined };
}

export const UpdateUserResponse = {
  encode(message: UpdateUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateUserResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user = User.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateUserResponse {
    return { user: isSet(object.user) ? User.fromJSON(object.user) : undefined };
  },

  toJSON(message: UpdateUserResponse): unknown {
    const obj: any = {};
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateUserResponse>, I>>(object: I): UpdateUserResponse {
    const message = createBaseUpdateUserResponse();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
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
