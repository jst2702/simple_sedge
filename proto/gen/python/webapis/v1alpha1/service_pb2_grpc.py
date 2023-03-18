# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from webapis.v1alpha1 import user_rpc_pb2 as webapis_dot_v1alpha1_dot_user__rpc__pb2


class SimpleSedgeStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateUser = channel.unary_unary(
                '/processing.v1alpha1.SimpleSedge/CreateUser',
                request_serializer=webapis_dot_v1alpha1_dot_user__rpc__pb2.CreateUserRequest.SerializeToString,
                response_deserializer=webapis_dot_v1alpha1_dot_user__rpc__pb2.CreateUserResponse.FromString,
                )
        self.LoginUser = channel.unary_unary(
                '/processing.v1alpha1.SimpleSedge/LoginUser',
                request_serializer=webapis_dot_v1alpha1_dot_user__rpc__pb2.LoginUserRequest.SerializeToString,
                response_deserializer=webapis_dot_v1alpha1_dot_user__rpc__pb2.LoginUserResponse.FromString,
                )


class SimpleSedgeServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreateUser(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def LoginUser(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SimpleSedgeServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateUser': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateUser,
                    request_deserializer=webapis_dot_v1alpha1_dot_user__rpc__pb2.CreateUserRequest.FromString,
                    response_serializer=webapis_dot_v1alpha1_dot_user__rpc__pb2.CreateUserResponse.SerializeToString,
            ),
            'LoginUser': grpc.unary_unary_rpc_method_handler(
                    servicer.LoginUser,
                    request_deserializer=webapis_dot_v1alpha1_dot_user__rpc__pb2.LoginUserRequest.FromString,
                    response_serializer=webapis_dot_v1alpha1_dot_user__rpc__pb2.LoginUserResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'processing.v1alpha1.SimpleSedge', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class SimpleSedge(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreateUser(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/processing.v1alpha1.SimpleSedge/CreateUser',
            webapis_dot_v1alpha1_dot_user__rpc__pb2.CreateUserRequest.SerializeToString,
            webapis_dot_v1alpha1_dot_user__rpc__pb2.CreateUserResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def LoginUser(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/processing.v1alpha1.SimpleSedge/LoginUser',
            webapis_dot_v1alpha1_dot_user__rpc__pb2.LoginUserRequest.SerializeToString,
            webapis_dot_v1alpha1_dot_user__rpc__pb2.LoginUserResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)