// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: proto/metrics.proto

package metrics

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	MetricsCollector_GetMetric_FullMethodName  = "/proto.MetricsCollector/GetMetric"
	MetricsCollector_GetMetrics_FullMethodName = "/proto.MetricsCollector/GetMetrics"
	MetricsCollector_AddMetric_FullMethodName  = "/proto.MetricsCollector/AddMetric"
	MetricsCollector_AddMetrics_FullMethodName = "/proto.MetricsCollector/AddMetrics"
)

// MetricsCollectorClient is the client API for MetricsCollector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsCollectorClient interface {
	GetMetric(ctx context.Context, in *GetMetricRequest, opts ...grpc.CallOption) (*GetMetricResponse, error)
	GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error)
	AddMetric(ctx context.Context, in *AddMetricRequest, opts ...grpc.CallOption) (*AddMetricResponse, error)
	AddMetrics(ctx context.Context, in *AddMetricsRequest, opts ...grpc.CallOption) (*AddMetricsResponse, error)
}

type metricsCollectorClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsCollectorClient(cc grpc.ClientConnInterface) MetricsCollectorClient {
	return &metricsCollectorClient{cc}
}

func (c *metricsCollectorClient) GetMetric(ctx context.Context, in *GetMetricRequest, opts ...grpc.CallOption) (*GetMetricResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMetricResponse)
	err := c.cc.Invoke(ctx, MetricsCollector_GetMetric_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsCollectorClient) GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMetricsResponse)
	err := c.cc.Invoke(ctx, MetricsCollector_GetMetrics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsCollectorClient) AddMetric(ctx context.Context, in *AddMetricRequest, opts ...grpc.CallOption) (*AddMetricResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddMetricResponse)
	err := c.cc.Invoke(ctx, MetricsCollector_AddMetric_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsCollectorClient) AddMetrics(ctx context.Context, in *AddMetricsRequest, opts ...grpc.CallOption) (*AddMetricsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddMetricsResponse)
	err := c.cc.Invoke(ctx, MetricsCollector_AddMetrics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsCollectorServer is the server API for MetricsCollector service.
// All implementations must embed UnimplementedMetricsCollectorServer
// for forward compatibility
type MetricsCollectorServer interface {
	GetMetric(context.Context, *GetMetricRequest) (*GetMetricResponse, error)
	GetMetrics(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error)
	AddMetric(context.Context, *AddMetricRequest) (*AddMetricResponse, error)
	AddMetrics(context.Context, *AddMetricsRequest) (*AddMetricsResponse, error)
	mustEmbedUnimplementedMetricsCollectorServer()
}

// UnimplementedMetricsCollectorServer must be embedded to have forward compatible implementations.
type UnimplementedMetricsCollectorServer struct {
}

func (UnimplementedMetricsCollectorServer) GetMetric(context.Context, *GetMetricRequest) (*GetMetricResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetric not implemented")
}
func (UnimplementedMetricsCollectorServer) GetMetrics(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetrics not implemented")
}
func (UnimplementedMetricsCollectorServer) AddMetric(context.Context, *AddMetricRequest) (*AddMetricResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMetric not implemented")
}
func (UnimplementedMetricsCollectorServer) AddMetrics(context.Context, *AddMetricsRequest) (*AddMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMetrics not implemented")
}
func (UnimplementedMetricsCollectorServer) mustEmbedUnimplementedMetricsCollectorServer() {}

// UnsafeMetricsCollectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsCollectorServer will
// result in compilation errors.
type UnsafeMetricsCollectorServer interface {
	mustEmbedUnimplementedMetricsCollectorServer()
}

func RegisterMetricsCollectorServer(s grpc.ServiceRegistrar, srv MetricsCollectorServer) {
	s.RegisterService(&MetricsCollector_ServiceDesc, srv)
}

func _MetricsCollector_GetMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsCollectorServer).GetMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsCollector_GetMetric_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsCollectorServer).GetMetric(ctx, req.(*GetMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsCollector_GetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsCollectorServer).GetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsCollector_GetMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsCollectorServer).GetMetrics(ctx, req.(*GetMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsCollector_AddMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsCollectorServer).AddMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsCollector_AddMetric_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsCollectorServer).AddMetric(ctx, req.(*AddMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsCollector_AddMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsCollectorServer).AddMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsCollector_AddMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsCollectorServer).AddMetrics(ctx, req.(*AddMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetricsCollector_ServiceDesc is the grpc.ServiceDesc for MetricsCollector service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricsCollector_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MetricsCollector",
	HandlerType: (*MetricsCollectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMetric",
			Handler:    _MetricsCollector_GetMetric_Handler,
		},
		{
			MethodName: "GetMetrics",
			Handler:    _MetricsCollector_GetMetrics_Handler,
		},
		{
			MethodName: "AddMetric",
			Handler:    _MetricsCollector_AddMetric_Handler,
		},
		{
			MethodName: "AddMetrics",
			Handler:    _MetricsCollector_AddMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/metrics.proto",
}
