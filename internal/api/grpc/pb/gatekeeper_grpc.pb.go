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
	ResetLogin(ctx context.Context, in *ResetLoginRequest, opts ...grpc.CallOption) (*ResetLoginResponse, error)
	ResetIP(ctx context.Context, in *ResetIPRequest, opts ...grpc.CallOption) (*ResetIPResponse, error)
	WhitelistAdd(ctx context.Context, in *WhitelistAddRequest, opts ...grpc.CallOption) (*WhitelistAddResponse, error)
	WhitelistRemove(ctx context.Context, in *WhitelistRemoveRequest, opts ...grpc.CallOption) (*WhitelistRemoveResponse, error)
	BlacklistAdd(ctx context.Context, in *BlacklistAddRequest, opts ...grpc.CallOption) (*BlacklistAddResponse, error)
	BlacklistRemove(ctx context.Context, in *BlacklistRemoveRequest, opts ...grpc.CallOption) (*BlacklistRemoveResponse, error)
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

func (c *gatekeeperServiceClient) ResetLogin(ctx context.Context, in *ResetLoginRequest, opts ...grpc.CallOption) (*ResetLoginResponse, error) {
	out := new(ResetLoginResponse)
	err := c.cc.Invoke(ctx, "/api.GatekeeperService/ResetLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatekeeperServiceClient) ResetIP(ctx context.Context, in *ResetIPRequest, opts ...grpc.CallOption) (*ResetIPResponse, error) {
	out := new(ResetIPResponse)
	err := c.cc.Invoke(ctx, "/api.GatekeeperService/ResetIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatekeeperServiceClient) WhitelistAdd(ctx context.Context, in *WhitelistAddRequest, opts ...grpc.CallOption) (*WhitelistAddResponse, error) {
	out := new(WhitelistAddResponse)
	err := c.cc.Invoke(ctx, "/api.GatekeeperService/WhitelistAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatekeeperServiceClient) WhitelistRemove(ctx context.Context, in *WhitelistRemoveRequest, opts ...grpc.CallOption) (*WhitelistRemoveResponse, error) {
	out := new(WhitelistRemoveResponse)
	err := c.cc.Invoke(ctx, "/api.GatekeeperService/WhitelistRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatekeeperServiceClient) BlacklistAdd(ctx context.Context, in *BlacklistAddRequest, opts ...grpc.CallOption) (*BlacklistAddResponse, error) {
	out := new(BlacklistAddResponse)
	err := c.cc.Invoke(ctx, "/api.GatekeeperService/BlacklistAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatekeeperServiceClient) BlacklistRemove(ctx context.Context, in *BlacklistRemoveRequest, opts ...grpc.CallOption) (*BlacklistRemoveResponse, error) {
	out := new(BlacklistRemoveResponse)
	err := c.cc.Invoke(ctx, "/api.GatekeeperService/BlacklistRemove", in, out, opts...)
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
	ResetLogin(context.Context, *ResetLoginRequest) (*ResetLoginResponse, error)
	ResetIP(context.Context, *ResetIPRequest) (*ResetIPResponse, error)
	WhitelistAdd(context.Context, *WhitelistAddRequest) (*WhitelistAddResponse, error)
	WhitelistRemove(context.Context, *WhitelistRemoveRequest) (*WhitelistRemoveResponse, error)
	BlacklistAdd(context.Context, *BlacklistAddRequest) (*BlacklistAddResponse, error)
	BlacklistRemove(context.Context, *BlacklistRemoveRequest) (*BlacklistRemoveResponse, error)
	mustEmbedUnimplementedGatekeeperServiceServer()
}

// UnimplementedGatekeeperServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGatekeeperServiceServer struct {
}

func (UnimplementedGatekeeperServiceServer) Allow(context.Context, *AllowRequest) (*AllowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Allow not implemented")
}
func (UnimplementedGatekeeperServiceServer) ResetLogin(context.Context, *ResetLoginRequest) (*ResetLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetLogin not implemented")
}
func (UnimplementedGatekeeperServiceServer) ResetIP(context.Context, *ResetIPRequest) (*ResetIPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetIP not implemented")
}
func (UnimplementedGatekeeperServiceServer) WhitelistAdd(context.Context, *WhitelistAddRequest) (*WhitelistAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhitelistAdd not implemented")
}
func (UnimplementedGatekeeperServiceServer) WhitelistRemove(context.Context, *WhitelistRemoveRequest) (*WhitelistRemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhitelistRemove not implemented")
}
func (UnimplementedGatekeeperServiceServer) BlacklistAdd(context.Context, *BlacklistAddRequest) (*BlacklistAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlacklistAdd not implemented")
}
func (UnimplementedGatekeeperServiceServer) BlacklistRemove(context.Context, *BlacklistRemoveRequest) (*BlacklistRemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlacklistRemove not implemented")
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

func _GatekeeperService_ResetLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatekeeperServiceServer).ResetLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatekeeperService/ResetLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatekeeperServiceServer).ResetLogin(ctx, req.(*ResetLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatekeeperService_ResetIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatekeeperServiceServer).ResetIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatekeeperService/ResetIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatekeeperServiceServer).ResetIP(ctx, req.(*ResetIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatekeeperService_WhitelistAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhitelistAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatekeeperServiceServer).WhitelistAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatekeeperService/WhitelistAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatekeeperServiceServer).WhitelistAdd(ctx, req.(*WhitelistAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatekeeperService_WhitelistRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhitelistRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatekeeperServiceServer).WhitelistRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatekeeperService/WhitelistRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatekeeperServiceServer).WhitelistRemove(ctx, req.(*WhitelistRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatekeeperService_BlacklistAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlacklistAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatekeeperServiceServer).BlacklistAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatekeeperService/BlacklistAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatekeeperServiceServer).BlacklistAdd(ctx, req.(*BlacklistAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatekeeperService_BlacklistRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlacklistRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatekeeperServiceServer).BlacklistRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatekeeperService/BlacklistRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatekeeperServiceServer).BlacklistRemove(ctx, req.(*BlacklistRemoveRequest))
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
		{
			MethodName: "ResetLogin",
			Handler:    _GatekeeperService_ResetLogin_Handler,
		},
		{
			MethodName: "ResetIP",
			Handler:    _GatekeeperService_ResetIP_Handler,
		},
		{
			MethodName: "WhitelistAdd",
			Handler:    _GatekeeperService_WhitelistAdd_Handler,
		},
		{
			MethodName: "WhitelistRemove",
			Handler:    _GatekeeperService_WhitelistRemove_Handler,
		},
		{
			MethodName: "BlacklistAdd",
			Handler:    _GatekeeperService_BlacklistAdd_Handler,
		},
		{
			MethodName: "BlacklistRemove",
			Handler:    _GatekeeperService_BlacklistRemove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gatekeeper.proto",
}
