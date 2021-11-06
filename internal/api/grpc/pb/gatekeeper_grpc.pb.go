// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package apipb

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GatekeeperServiceClient is the client API for GatekeeperService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatekeeperServiceClient interface {
	Allow(ctx context.Context, in *AllowRequest, opts ...grpc.CallOption) (*AllowResponse, error)
}

type gatekeeperServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGatekeeperServiceClient(cc grpc.ClientConnInterface) GatekeeperServiceClient {
	return &gatekeeperServiceClient{cc}
}

func (c *gatekeeperServiceClient) Allow(ctx context.Context, in *AllowRequest, opts ...grpc.CallOption) (*AllowResponse, error) {
	out := new(AllowResponse)
	err := c.cc.Invoke(ctx, "/api.GatekeeperService/Allow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatekeeperServiceServer is the server API for GatekeeperService service.
// All implementations must embed UnimplementedGatekeeperServiceServer
// for forward compatibility
type GatekeeperServiceServer interface {
	Allow(context.Context, *AllowRequest) (*AllowResponse, error)
	mustEmbedUnimplementedGatekeeperServiceServer()
}

// UnimplementedGatekeeperServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGatekeeperServiceServer struct{}

func (UnimplementedGatekeeperServiceServer) Allow(context.Context, *AllowRequest) (*AllowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Allow not implemented")
}
func (UnimplementedGatekeeperServiceServer) mustEmbedUnimplementedGatekeeperServiceServer() {}

// UnsafeGatekeeperServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatekeeperServiceServer will
// result in compilation errors.
type UnsafeGatekeeperServiceServer interface {
	mustEmbedUnimplementedGatekeeperServiceServer()
}

func RegisterGatekeeperServiceServer(s grpc.ServiceRegistrar, srv GatekeeperServiceServer) {
	s.RegisterService(&GatekeeperService_ServiceDesc, srv)
}

func _GatekeeperService_Allow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatekeeperServiceServer).Allow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatekeeperService/Allow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatekeeperServiceServer).Allow(ctx, req.(*AllowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GatekeeperService_ServiceDesc is the grpc.ServiceDesc for GatekeeperService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatekeeperService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.GatekeeperService",
	HandlerType: (*GatekeeperServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Allow",
			Handler:    _GatekeeperService_Allow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gatekeeper.proto",
}
