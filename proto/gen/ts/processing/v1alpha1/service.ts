/* eslint-disable */
import _m0 from "protobufjs/minimal";
import {
  IngestDocumentRequest,
  IngestDocumentResponse,
  ListDocumentsRequest,
  ListDocumentsResponse,
} from "./rpc_document";

export const protobufPackage = "processing.v1alpha1";

export interface IngestionService {
  /** Ingest a new data object for processing */
  IngestDocument(request: IngestDocumentRequest): Promise<IngestDocumentResponse>;
  ListDocuments(request: ListDocumentsRequest): Promise<ListDocumentsResponse>;
}

export class IngestionServiceClientImpl implements IngestionService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "processing.v1alpha1.IngestionService";
    this.rpc = rpc;
    this.IngestDocument = this.IngestDocument.bind(this);
    this.ListDocuments = this.ListDocuments.bind(this);
  }
  IngestDocument(request: IngestDocumentRequest): Promise<IngestDocumentResponse> {
    const data = IngestDocumentRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "IngestDocument", data);
    return promise.then((data) => IngestDocumentResponse.decode(new _m0.Reader(data)));
  }

  ListDocuments(request: ListDocumentsRequest): Promise<ListDocumentsResponse> {
    const data = ListDocumentsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListDocuments", data);
    return promise.then((data) => ListDocumentsResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
