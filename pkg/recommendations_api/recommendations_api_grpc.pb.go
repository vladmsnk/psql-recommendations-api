// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: recommendations_api/recommendations_api.proto

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
	RecommendationsAPI_GetRecommendations_FullMethodName = "/collector.RecommendationsAPI/GetRecommendations"
)

// RecommendationsAPIClient is the client API for RecommendationsAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecommendationsAPIClient interface {
	GetRecommendations(ctx context.Context, in *GetRecommendationsRequest, opts ...grpc.CallOption) (*GetRecommendationsResponse, error)
}

type recommendationsAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewRecommendationsAPIClient(cc grpc.ClientConnInterface) RecommendationsAPIClient {
	return &recommendationsAPIClient{cc}
}

func (c *recommendationsAPIClient) GetRecommendations(ctx context.Context, in *GetRecommendationsRequest, opts ...grpc.CallOption) (*GetRecommendationsResponse, error) {
	out := new(GetRecommendationsResponse)
	err := c.cc.Invoke(ctx, RecommendationsAPI_GetRecommendations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecommendationsAPIServer is the server API for RecommendationsAPI service.
// All implementations must embed UnimplementedRecommendationsAPIServer
// for forward compatibility
type RecommendationsAPIServer interface {
	GetRecommendations(context.Context, *GetRecommendationsRequest) (*GetRecommendationsResponse, error)
	mustEmbedUnimplementedRecommendationsAPIServer()
}

// UnimplementedRecommendationsAPIServer must be embedded to have forward compatible implementations.
type UnimplementedRecommendationsAPIServer struct {
}

func (UnimplementedRecommendationsAPIServer) GetRecommendations(context.Context, *GetRecommendationsRequest) (*GetRecommendationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecommendations not implemented")
}
func (UnimplementedRecommendationsAPIServer) mustEmbedUnimplementedRecommendationsAPIServer() {}

// UnsafeRecommendationsAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecommendationsAPIServer will
// result in compilation errors.
type UnsafeRecommendationsAPIServer interface {
	mustEmbedUnimplementedRecommendationsAPIServer()
}

func RegisterRecommendationsAPIServer(s grpc.ServiceRegistrar, srv RecommendationsAPIServer) {
	s.RegisterService(&RecommendationsAPI_ServiceDesc, srv)
}

func _RecommendationsAPI_GetRecommendations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecommendationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationsAPIServer).GetRecommendations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecommendationsAPI_GetRecommendations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationsAPIServer).GetRecommendations(ctx, req.(*GetRecommendationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecommendationsAPI_ServiceDesc is the grpc.ServiceDesc for RecommendationsAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecommendationsAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "collector.RecommendationsAPI",
	HandlerType: (*RecommendationsAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRecommendations",
			Handler:    _RecommendationsAPI_GetRecommendations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recommendations_api/recommendations_api.proto",
}
