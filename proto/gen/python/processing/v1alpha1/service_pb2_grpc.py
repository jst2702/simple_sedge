# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from processing.v1alpha1 import rpc_document_pb2 as processing_dot_v1alpha1_dot_rpc__document__pb2


class IngestionServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.IngestDocument = channel.unary_unary(
                '/processing.IngestionService/IngestDocument',
                request_serializer=processing_dot_v1alpha1_dot_rpc__document__pb2.IngestDocumentRequest.SerializeToString,
                response_deserializer=processing_dot_v1alpha1_dot_rpc__document__pb2.IngestDocumentResponse.FromString,
                )
        self.ListDocuments = channel.unary_unary(
                '/processing.IngestionService/ListDocuments',
                request_serializer=processing_dot_v1alpha1_dot_rpc__document__pb2.ListDocumentsRequest.SerializeToString,
                response_deserializer=processing_dot_v1alpha1_dot_rpc__document__pb2.ListDocumentsResponse.FromString,
                )


class IngestionServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def IngestDocument(self, request, context):
        """Ingest a new data object for processing
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListDocuments(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_IngestionServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'IngestDocument': grpc.unary_unary_rpc_method_handler(
                    servicer.IngestDocument,
                    request_deserializer=processing_dot_v1alpha1_dot_rpc__document__pb2.IngestDocumentRequest.FromString,
                    response_serializer=processing_dot_v1alpha1_dot_rpc__document__pb2.IngestDocumentResponse.SerializeToString,
            ),
            'ListDocuments': grpc.unary_unary_rpc_method_handler(
                    servicer.ListDocuments,
                    request_deserializer=processing_dot_v1alpha1_dot_rpc__document__pb2.ListDocumentsRequest.FromString,
                    response_serializer=processing_dot_v1alpha1_dot_rpc__document__pb2.ListDocumentsResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'processing.IngestionService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class IngestionService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def IngestDocument(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/processing.IngestionService/IngestDocument',
            processing_dot_v1alpha1_dot_rpc__document__pb2.IngestDocumentRequest.SerializeToString,
            processing_dot_v1alpha1_dot_rpc__document__pb2.IngestDocumentResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListDocuments(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/processing.IngestionService/ListDocuments',
            processing_dot_v1alpha1_dot_rpc__document__pb2.ListDocumentsRequest.SerializeToString,
            processing_dot_v1alpha1_dot_rpc__document__pb2.ListDocumentsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
