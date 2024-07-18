// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: historical.proto

package genproto

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

// HistoricalServiceClient is the client API for HistoricalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HistoricalServiceClient interface {
	GetAllHistoricalEvents(ctx context.Context, in *GetAllHistoricalRequest, opts ...grpc.CallOption) (*GetAllHistoricalResponse, error)
	AddHistoricalEvent(ctx context.Context, in *AddHistoricalEventRequest, opts ...grpc.CallOption) (*AddHistoricalEventResponse, error)
	DeleteHistoricalEvent(ctx context.Context, in *DeleteHistoricalEventRequest, opts ...grpc.CallOption) (*DeleteHistoricalEventResponse, error)
}

type historicalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHistoricalServiceClient(cc grpc.ClientConnInterface) HistoricalServiceClient {
	return &historicalServiceClient{cc}
}

func (c *historicalServiceClient) GetAllHistoricalEvents(ctx context.Context, in *GetAllHistoricalRequest, opts ...grpc.CallOption) (*GetAllHistoricalResponse, error) {
	out := new(GetAllHistoricalResponse)
	err := c.cc.Invoke(ctx, "/timecapsule.HistoricalService/GetAllHistoricalEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historicalServiceClient) AddHistoricalEvent(ctx context.Context, in *AddHistoricalEventRequest, opts ...grpc.CallOption) (*AddHistoricalEventResponse, error) {
	out := new(AddHistoricalEventResponse)
	err := c.cc.Invoke(ctx, "/timecapsule.HistoricalService/AddHistoricalEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historicalServiceClient) DeleteHistoricalEvent(ctx context.Context, in *DeleteHistoricalEventRequest, opts ...grpc.CallOption) (*DeleteHistoricalEventResponse, error) {
	out := new(DeleteHistoricalEventResponse)
	err := c.cc.Invoke(ctx, "/timecapsule.HistoricalService/DeleteHistoricalEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HistoricalServiceServer is the server API for HistoricalService service.
// All implementations must embed UnimplementedHistoricalServiceServer
// for forward compatibility
type HistoricalServiceServer interface {
	GetAllHistoricalEvents(context.Context, *GetAllHistoricalRequest) (*GetAllHistoricalResponse, error)
	AddHistoricalEvent(context.Context, *AddHistoricalEventRequest) (*AddHistoricalEventResponse, error)
	DeleteHistoricalEvent(context.Context, *DeleteHistoricalEventRequest) (*DeleteHistoricalEventResponse, error)
	mustEmbedUnimplementedHistoricalServiceServer()
}

// UnimplementedHistoricalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHistoricalServiceServer struct {
}

func (UnimplementedHistoricalServiceServer) GetAllHistoricalEvents(context.Context, *GetAllHistoricalRequest) (*GetAllHistoricalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllHistoricalEvents not implemented")
}
func (UnimplementedHistoricalServiceServer) AddHistoricalEvent(context.Context, *AddHistoricalEventRequest) (*AddHistoricalEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddHistoricalEvent not implemented")
}
func (UnimplementedHistoricalServiceServer) DeleteHistoricalEvent(context.Context, *DeleteHistoricalEventRequest) (*DeleteHistoricalEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHistoricalEvent not implemented")
}
func (UnimplementedHistoricalServiceServer) mustEmbedUnimplementedHistoricalServiceServer() {}

// UnsafeHistoricalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HistoricalServiceServer will
// result in compilation errors.
type UnsafeHistoricalServiceServer interface {
	mustEmbedUnimplementedHistoricalServiceServer()
}

func RegisterHistoricalServiceServer(s grpc.ServiceRegistrar, srv HistoricalServiceServer) {
	s.RegisterService(&HistoricalService_ServiceDesc, srv)
}

func _HistoricalService_GetAllHistoricalEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllHistoricalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoricalServiceServer).GetAllHistoricalEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timecapsule.HistoricalService/GetAllHistoricalEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoricalServiceServer).GetAllHistoricalEvents(ctx, req.(*GetAllHistoricalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HistoricalService_AddHistoricalEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddHistoricalEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoricalServiceServer).AddHistoricalEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timecapsule.HistoricalService/AddHistoricalEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoricalServiceServer).AddHistoricalEvent(ctx, req.(*AddHistoricalEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HistoricalService_DeleteHistoricalEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHistoricalEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoricalServiceServer).DeleteHistoricalEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timecapsule.HistoricalService/DeleteHistoricalEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoricalServiceServer).DeleteHistoricalEvent(ctx, req.(*DeleteHistoricalEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HistoricalService_ServiceDesc is the grpc.ServiceDesc for HistoricalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HistoricalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "timecapsule.HistoricalService",
	HandlerType: (*HistoricalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllHistoricalEvents",
			Handler:    _HistoricalService_GetAllHistoricalEvents_Handler,
		},
		{
			MethodName: "AddHistoricalEvent",
			Handler:    _HistoricalService_AddHistoricalEvent_Handler,
		},
		{
			MethodName: "DeleteHistoricalEvent",
			Handler:    _HistoricalService_DeleteHistoricalEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "historical.proto",
}