/* eslint-disable */
import _m0 from "protobufjs/minimal";
import {
  IngestDocumentRequest,
  IngestDocumentResponse,
  ListDocumentsRequest,
  ListDocumentsResponse,
} from "./rpc_document";
import {
  CreateUserRequest,
  CreateUserResponse,
  LoginUserRequest,
  LoginUserResponse,
  UpdateUserRequest,
  UpdateUserResponse,
  VerifyEmailRequest,
  VerifyEmailResponse,
} from "./rpc_user";

export const protobufPackage = "apis";

export interface SimpleSedge {
  /** Ingest a new data object for processing */
  CreateUser(request: CreateUserRequest): Promise<CreateUserResponse>;
  LoginUser(request: LoginUserRequest): Promise<LoginUserResponse>;
  UpdateUser(request: UpdateUserRequest): Promise<UpdateUserResponse>;
  VerifyEmail(request: VerifyEmailRequest): Promise<VerifyEmailResponse>;
  ListDocuments(request: ListDocumentsRequest): Promise<ListDocumentsResponse>;
  IngestDocument(request: IngestDocumentRequest): Promise<IngestDocumentResponse>;
}

export class SimpleSedgeClientImpl implements SimpleSedge {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "apis.SimpleSedge";
    this.rpc = rpc;
    this.CreateUser = this.CreateUser.bind(this);
    this.LoginUser = this.LoginUser.bind(this);
    this.UpdateUser = this.UpdateUser.bind(this);
    this.VerifyEmail = this.VerifyEmail.bind(this);
    this.ListDocuments = this.ListDocuments.bind(this);
    this.IngestDocument = this.IngestDocument.bind(this);
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

  ListDocuments(request: ListDocumentsRequest): Promise<ListDocumentsResponse> {
    const data = ListDocumentsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListDocuments", data);
    return promise.then((data) => ListDocumentsResponse.decode(new _m0.Reader(data)));
  }

  IngestDocument(request: IngestDocumentRequest): Promise<IngestDocumentResponse> {
    const data = IngestDocumentRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "IngestDocument", data);
    return promise.then((data) => IngestDocumentResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
