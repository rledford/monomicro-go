// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/v1/randint.proto

package apiv1

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

// RandintServiceClient is the client API for RandintService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RandintServiceClient interface {
	GetRandint(ctx context.Context, in *GetRandintRequest, opts ...grpc.CallOption) (*GetRandintResponse, error)
}

type randintServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRandintServiceClient(cc grpc.ClientConnInterface) RandintServiceClient {
	return &randintServiceClient{cc}
}

func (c *randintServiceClient) GetRandint(ctx context.Context, in *GetRandintRequest, opts ...grpc.CallOption) (*GetRandintResponse, error) {
	out := new(GetRandintResponse)
	err := c.cc.Invoke(ctx, "/api.v1.RandintService/GetRandint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RandintServiceServer is the server API for RandintService service.
// All implementations should embed UnimplementedRandintServiceServer
// for forward compatibility
type RandintServiceServer interface {
	GetRandint(context.Context, *GetRandintRequest) (*GetRandintResponse, error)
}

// UnimplementedRandintServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRandintServiceServer struct {
}

func (UnimplementedRandintServiceServer) GetRandint(context.Context, *GetRandintRequest) (*GetRandintResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandint not implemented")
}

// UnsafeRandintServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RandintServiceServer will
// result in compilation errors.
type UnsafeRandintServiceServer interface {
	mustEmbedUnimplementedRandintServiceServer()
}

func RegisterRandintServiceServer(s grpc.ServiceRegistrar, srv RandintServiceServer) {
	s.RegisterService(&RandintService_ServiceDesc, srv)
}

func _RandintService_GetRandint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRandintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandintServiceServer).GetRandint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.RandintService/GetRandint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandintServiceServer).GetRandint(ctx, req.(*GetRandintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RandintService_ServiceDesc is the grpc.ServiceDesc for RandintService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RandintService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.RandintService",
	HandlerType: (*RandintServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRandint",
			Handler:    _RandintService_GetRandint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/randint.proto",
}