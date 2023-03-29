/* eslint-disable */
import _m0 from "protobufjs/minimal";
import {
  CreateUserRequest,
  CreateUserResponse,
  LoginUserRequest,
  LoginUserResponse,
  UpdateUserRequest,
  UpdateUserResponse,
  VerifyEmailRequest,
  VerifyEmailResponse,
} from "./user_rpc";

export const protobufPackage = "processing.v1alpha1";

export interface SimpleSedge {
  CreateUser(request: CreateUserRequest): Promise<CreateUserResponse>;
  LoginUser(request: LoginUserRequest): Promise<LoginUserResponse>;
  UpdateUser(request: UpdateUserRequest): Promise<UpdateUserResponse>;
  VerifyEmail(request: VerifyEmailRequest): Promise<VerifyEmailResponse>;
}

export class SimpleSedgeClientImpl implements SimpleSedge {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "processing.v1alpha1.SimpleSedge";
    this.rpc = rpc;
    this.CreateUser = this.CreateUser.bind(this);
    this.LoginUser = this.LoginUser.bind(this);
    this.UpdateUser = this.UpdateUser.bind(this);
    this.VerifyEmail = this.VerifyEmail.bind(this);
  }
  CreateUser(request: CreateUserRequest): Promise<CreateUserResponse> {
    const data = CreateUserRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateUser", data);
    return promise.then((data) => CreateUserResponse.decode(new _m0.Reader(data)));
  }

  LoginUser(request: LoginUserRequest): Promise<LoginUserResponse> {
    const data = LoginUserRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "LoginUser", data);
    return promise.then((data) => LoginUserResponse.decode(new _m0.Reader(data)));
  }

  UpdateUser(request: UpdateUserRequest): Promise<UpdateUserResponse> {
    const data = UpdateUserRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateUser", data);
    return promise.then((data) => UpdateUserResponse.decode(new _m0.Reader(data)));
  }

  VerifyEmail(request: VerifyEmailRequest): Promise<VerifyEmailResponse> {
    const data = VerifyEmailRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "VerifyEmail", data);
    return promise.then((data) => VerifyEmailResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
