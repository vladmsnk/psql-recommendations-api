// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: collector/colelctor.proto

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
	Collector_CollectKnobs_FullMethodName           = "/collector.Collector/CollectKnobs"
	Collector_CollectInternalMetrics_FullMethodName = "/collector.Collector/CollectInternalMetrics"
	Collector_CollectExternalMetrics_FullMethodName = "/collector.Collector/CollectExternalMetrics"
	Collector_InitLoad_FullMethodName               = "/collector.Collector/InitLoad"
	Collector_SetKnobs_FullMethodName               = "/collector.Collector/SetKnobs"
)

// CollectorClient is the client API for Collector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollectorClient interface {
	// Collects PostgreSQL knobs
	CollectKnobs(ctx context.Context, in *CollectKnobsRequest, opts ...grpc.CallOption) (*CollectKnobsResponse, error)
	CollectInternalMetrics(ctx context.Context, in *CollectInternalMetricsRequest, opts ...grpc.CallOption) (*CollectInternalMetricsResponse, error)
	CollectExternalMetrics(ctx context.Context, in *CollectExternalMetricsRequest, opts ...grpc.CallOption) (*CollectExternalMetricsResponse, error)
	InitLoad(ctx context.Context, in *InitLoadRequest, opts ...grpc.CallOption) (*InitLoadResponse, error)
	// Применяет параметры конфигурации
	SetKnobs(ctx context.Context, in *SetKnobsRequest, opts ...grpc.CallOption) (*SetKnobsResponse, error)
}

type collectorClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectorClient(cc grpc.ClientConnInterface) CollectorClient {
	return &collectorClient{cc}
}

func (c *collectorClient) CollectKnobs(ctx context.Context, in *CollectKnobsRequest, opts ...grpc.CallOption) (*CollectKnobsResponse, error) {
	out := new(CollectKnobsResponse)
	err := c.cc.Invoke(ctx, Collector_CollectKnobs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectorClient) CollectInternalMetrics(ctx context.Context, in *CollectInternalMetricsRequest, opts ...grpc.CallOption) (*CollectInternalMetricsResponse, error) {
	out := new(CollectInternalMetricsResponse)
	err := c.cc.Invoke(ctx, Collector_CollectInternalMetrics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectorClient) CollectExternalMetrics(ctx context.Context, in *CollectExternalMetricsRequest, opts ...grpc.CallOption) (*CollectExternalMetricsResponse, error) {
	out := new(CollectExternalMetricsResponse)
	err := c.cc.Invoke(ctx, Collector_CollectExternalMetrics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectorClient) InitLoad(ctx context.Context, in *InitLoadRequest, opts ...grpc.CallOption) (*InitLoadResponse, error) {
	out := new(InitLoadResponse)
	err := c.cc.Invoke(ctx, Collector_InitLoad_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectorClient) SetKnobs(ctx context.Context, in *SetKnobsRequest, opts ...grpc.CallOption) (*SetKnobsResponse, error) {
	out := new(SetKnobsResponse)
	err := c.cc.Invoke(ctx, Collector_SetKnobs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectorServer is the server API for Collector service.
// All implementations must embed UnimplementedCollectorServer
// for forward compatibility
type CollectorServer interface {
	// Collects PostgreSQL knobs
	CollectKnobs(context.Context, *CollectKnobsRequest) (*CollectKnobsResponse, error)
	CollectInternalMetrics(context.Context, *CollectInternalMetricsRequest) (*CollectInternalMetricsResponse, error)
	CollectExternalMetrics(context.Context, *CollectExternalMetricsRequest) (*CollectExternalMetricsResponse, error)
	InitLoad(context.Context, *InitLoadRequest) (*InitLoadResponse, error)
	// Применяет параметры конфигурации
	SetKnobs(context.Context, *SetKnobsRequest) (*SetKnobsResponse, error)
	mustEmbedUnimplementedCollectorServer()
}

// UnimplementedCollectorServer must be embedded to have forward compatible implementations.
type UnimplementedCollectorServer struct {
}

func (UnimplementedCollectorServer) CollectKnobs(context.Context, *CollectKnobsRequest) (*CollectKnobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectKnobs not implemented")
}
func (UnimplementedCollectorServer) CollectInternalMetrics(context.Context, *CollectInternalMetricsRequest) (*CollectInternalMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectInternalMetrics not implemented")
}
func (UnimplementedCollectorServer) CollectExternalMetrics(context.Context, *CollectExternalMetricsRequest) (*CollectExternalMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectExternalMetrics not implemented")
}
func (UnimplementedCollectorServer) InitLoad(context.Context, *InitLoadRequest) (*InitLoadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitLoad not implemented")
}
func (UnimplementedCollectorServer) SetKnobs(context.Context, *SetKnobsRequest) (*SetKnobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetKnobs not implemented")
}
func (UnimplementedCollectorServer) mustEmbedUnimplementedCollectorServer() {}

// UnsafeCollectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollectorServer will
// result in compilation errors.
type UnsafeCollectorServer interface {
	mustEmbedUnimplementedCollectorServer()
}

func RegisterCollectorServer(s grpc.ServiceRegistrar, srv CollectorServer) {
	s.RegisterService(&Collector_ServiceDesc, srv)
}

func _Collector_CollectKnobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectKnobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorServer).CollectKnobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Collector_CollectKnobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorServer).CollectKnobs(ctx, req.(*CollectKnobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collector_CollectInternalMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectInternalMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorServer).CollectInternalMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Collector_CollectInternalMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorServer).CollectInternalMetrics(ctx, req.(*CollectInternalMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collector_CollectExternalMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectExternalMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorServer).CollectExternalMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Collector_CollectExternalMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorServer).CollectExternalMetrics(ctx, req.(*CollectExternalMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collector_InitLoad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitLoadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorServer).InitLoad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Collector_InitLoad_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorServer).InitLoad(ctx, req.(*InitLoadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collector_SetKnobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetKnobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorServer).SetKnobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Collector_SetKnobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorServer).SetKnobs(ctx, req.(*SetKnobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Collector_ServiceDesc is the grpc.ServiceDesc for Collector service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Collector_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "collector.Collector",
	HandlerType: (*CollectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CollectKnobs",
			Handler:    _Collector_CollectKnobs_Handler,
		},
		{
			MethodName: "CollectInternalMetrics",
			Handler:    _Collector_CollectInternalMetrics_Handler,
		},
		{
			MethodName: "CollectExternalMetrics",
			Handler:    _Collector_CollectExternalMetrics_Handler,
		},
		{
			MethodName: "InitLoad",
			Handler:    _Collector_InitLoad_Handler,
		},
		{
			MethodName: "SetKnobs",
			Handler:    _Collector_SetKnobs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collector/colelctor.proto",
}