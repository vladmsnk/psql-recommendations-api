// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: discovery/discovery.proto

package pb

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

const (
	Discovery_RegisterInstance_FullMethodName = "/discovery.Discovery/RegisterInstance"
	Discovery_GetInstanceInfo_FullMethodName  = "/discovery.Discovery/GetInstanceInfo"
)

// DiscoveryClient is the client API for Discovery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscoveryClient interface {
	RegisterInstance(ctx context.Context, in *RegisterInstanceRequest, opts ...grpc.CallOption) (*RegisterInstanceResponse, error)
	GetInstanceInfo(ctx context.Context, in *GetInstanceInfoRequest, opts ...grpc.CallOption) (*GetInstanceInfoResponse, error)
}

type discoveryClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscoveryClient(cc grpc.ClientConnInterface) DiscoveryClient {
	return &discoveryClient{cc}
}

func (c *discoveryClient) RegisterInstance(ctx context.Context, in *RegisterInstanceRequest, opts ...grpc.CallOption) (*RegisterInstanceResponse, error) {
	out := new(RegisterInstanceResponse)
	err := c.cc.Invoke(ctx, Discovery_RegisterInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) GetInstanceInfo(ctx context.Context, in *GetInstanceInfoRequest, opts ...grpc.CallOption) (*GetInstanceInfoResponse, error) {
	out := new(GetInstanceInfoResponse)
	err := c.cc.Invoke(ctx, Discovery_GetInstanceInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscoveryServer is the server API for Discovery service.
// All implementations must embed UnimplementedDiscoveryServer
// for forward compatibility
type DiscoveryServer interface {
	RegisterInstance(context.Context, *RegisterInstanceRequest) (*RegisterInstanceResponse, error)
	GetInstanceInfo(context.Context, *GetInstanceInfoRequest) (*GetInstanceInfoResponse, error)
	mustEmbedUnimplementedDiscoveryServer()
}

// UnimplementedDiscoveryServer must be embedded to have forward compatible implementations.
type UnimplementedDiscoveryServer struct {
}

func (UnimplementedDiscoveryServer) RegisterInstance(context.Context, *RegisterInstanceRequest) (*RegisterInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterInstance not implemented")
}
func (UnimplementedDiscoveryServer) GetInstanceInfo(context.Context, *GetInstanceInfoRequest) (*GetInstanceInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstanceInfo not implemented")
}
func (UnimplementedDiscoveryServer) mustEmbedUnimplementedDiscoveryServer() {}

// UnsafeDiscoveryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscoveryServer will
// result in compilation errors.
type UnsafeDiscoveryServer interface {
	mustEmbedUnimplementedDiscoveryServer()
}

func RegisterDiscoveryServer(s grpc.ServiceRegistrar, srv DiscoveryServer) {
	s.RegisterService(&Discovery_ServiceDesc, srv)
}

func _Discovery_RegisterInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).RegisterInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Discovery_RegisterInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).RegisterInstance(ctx, req.(*RegisterInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discovery_GetInstanceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstanceInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).GetInstanceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Discovery_GetInstanceInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).GetInstanceInfo(ctx, req.(*GetInstanceInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Discovery_ServiceDesc is the grpc.ServiceDesc for Discovery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Discovery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "discovery.Discovery",
	HandlerType: (*DiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterInstance",
			Handler:    _Discovery_RegisterInstance_Handler,
		},
		{
			MethodName: "GetInstanceInfo",
			Handler:    _Discovery_GetInstanceInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discovery/discovery.proto",
}
