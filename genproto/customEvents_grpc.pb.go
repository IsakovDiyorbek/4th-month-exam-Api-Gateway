// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: customEvents.proto

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

// CustomEventsServiceClient is the client API for CustomEventsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomEventsServiceClient interface {
	AddCustomEvent(ctx context.Context, in *AddCustomEventRequest, opts ...grpc.CallOption) (*AddCustomEventResponse, error)
	UpdateCustomEvent(ctx context.Context, in *UpdateCustomEventsRequest, opts ...grpc.CallOption) (*UpdateCustomEventsResponse, error)
	DeleteCustomEvent(ctx context.Context, in *DeleteCustomEventsRequest, opts ...grpc.CallOption) (*DeleteCustomEventsResponse, error)
	GetAllCustomEvents(ctx context.Context, in *GetAllEventsRequest, opts ...grpc.CallOption) (*GetAllEventsResponse, error)
	GetByIdCustomEvent(ctx context.Context, in *GetByIdEvetsRequest, opts ...grpc.CallOption) (*GetByIdEvetsResponse, error)
}

type customEventsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomEventsServiceClient(cc grpc.ClientConnInterface) CustomEventsServiceClient {
	return &customEventsServiceClient{cc}
}

func (c *customEventsServiceClient) AddCustomEvent(ctx context.Context, in *AddCustomEventRequest, opts ...grpc.CallOption) (*AddCustomEventResponse, error) {
	out := new(AddCustomEventResponse)
	err := c.cc.Invoke(ctx, "/timecapsule.CustomEventsService/AddCustomEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customEventsServiceClient) UpdateCustomEvent(ctx context.Context, in *UpdateCustomEventsRequest, opts ...grpc.CallOption) (*UpdateCustomEventsResponse, error) {
	out := new(UpdateCustomEventsResponse)
	err := c.cc.Invoke(ctx, "/timecapsule.CustomEventsService/UpdateCustomEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customEventsServiceClient) DeleteCustomEvent(ctx context.Context, in *DeleteCustomEventsRequest, opts ...grpc.CallOption) (*DeleteCustomEventsResponse, error) {
	out := new(DeleteCustomEventsResponse)
	err := c.cc.Invoke(ctx, "/timecapsule.CustomEventsService/DeleteCustomEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customEventsServiceClient) GetAllCustomEvents(ctx context.Context, in *GetAllEventsRequest, opts ...grpc.CallOption) (*GetAllEventsResponse, error) {
	out := new(GetAllEventsResponse)
	err := c.cc.Invoke(ctx, "/timecapsule.CustomEventsService/GetAllCustomEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customEventsServiceClient) GetByIdCustomEvent(ctx context.Context, in *GetByIdEvetsRequest, opts ...grpc.CallOption) (*GetByIdEvetsResponse, error) {
	out := new(GetByIdEvetsResponse)
	err := c.cc.Invoke(ctx, "/timecapsule.CustomEventsService/GetByIdCustomEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomEventsServiceServer is the server API for CustomEventsService service.
// All implementations must embed UnimplementedCustomEventsServiceServer
// for forward compatibility
type CustomEventsServiceServer interface {
	AddCustomEvent(context.Context, *AddCustomEventRequest) (*AddCustomEventResponse, error)
	UpdateCustomEvent(context.Context, *UpdateCustomEventsRequest) (*UpdateCustomEventsResponse, error)
	DeleteCustomEvent(context.Context, *DeleteCustomEventsRequest) (*DeleteCustomEventsResponse, error)
	GetAllCustomEvents(context.Context, *GetAllEventsRequest) (*GetAllEventsResponse, error)
	GetByIdCustomEvent(context.Context, *GetByIdEvetsRequest) (*GetByIdEvetsResponse, error)
	mustEmbedUnimplementedCustomEventsServiceServer()
}

// UnimplementedCustomEventsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomEventsServiceServer struct {
}

func (UnimplementedCustomEventsServiceServer) AddCustomEvent(context.Context, *AddCustomEventRequest) (*AddCustomEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCustomEvent not implemented")
}
func (UnimplementedCustomEventsServiceServer) UpdateCustomEvent(context.Context, *UpdateCustomEventsRequest) (*UpdateCustomEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomEvent not implemented")
}
func (UnimplementedCustomEventsServiceServer) DeleteCustomEvent(context.Context, *DeleteCustomEventsRequest) (*DeleteCustomEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCustomEvent not implemented")
}
func (UnimplementedCustomEventsServiceServer) GetAllCustomEvents(context.Context, *GetAllEventsRequest) (*GetAllEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCustomEvents not implemented")
}
func (UnimplementedCustomEventsServiceServer) GetByIdCustomEvent(context.Context, *GetByIdEvetsRequest) (*GetByIdEvetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdCustomEvent not implemented")
}
func (UnimplementedCustomEventsServiceServer) mustEmbedUnimplementedCustomEventsServiceServer() {}

// UnsafeCustomEventsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomEventsServiceServer will
// result in compilation errors.
type UnsafeCustomEventsServiceServer interface {
	mustEmbedUnimplementedCustomEventsServiceServer()
}

func RegisterCustomEventsServiceServer(s grpc.ServiceRegistrar, srv CustomEventsServiceServer) {
	s.RegisterService(&CustomEventsService_ServiceDesc, srv)
}

func _CustomEventsService_AddCustomEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCustomEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomEventsServiceServer).AddCustomEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timecapsule.CustomEventsService/AddCustomEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomEventsServiceServer).AddCustomEvent(ctx, req.(*AddCustomEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomEventsService_UpdateCustomEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCustomEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomEventsServiceServer).UpdateCustomEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timecapsule.CustomEventsService/UpdateCustomEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomEventsServiceServer).UpdateCustomEvent(ctx, req.(*UpdateCustomEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomEventsService_DeleteCustomEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCustomEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomEventsServiceServer).DeleteCustomEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timecapsule.CustomEventsService/DeleteCustomEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomEventsServiceServer).DeleteCustomEvent(ctx, req.(*DeleteCustomEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomEventsService_GetAllCustomEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomEventsServiceServer).GetAllCustomEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timecapsule.CustomEventsService/GetAllCustomEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomEventsServiceServer).GetAllCustomEvents(ctx, req.(*GetAllEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomEventsService_GetByIdCustomEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdEvetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomEventsServiceServer).GetByIdCustomEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timecapsule.CustomEventsService/GetByIdCustomEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomEventsServiceServer).GetByIdCustomEvent(ctx, req.(*GetByIdEvetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomEventsService_ServiceDesc is the grpc.ServiceDesc for CustomEventsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomEventsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "timecapsule.CustomEventsService",
	HandlerType: (*CustomEventsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCustomEvent",
			Handler:    _CustomEventsService_AddCustomEvent_Handler,
		},
		{
			MethodName: "UpdateCustomEvent",
			Handler:    _CustomEventsService_UpdateCustomEvent_Handler,
		},
		{
			MethodName: "DeleteCustomEvent",
			Handler:    _CustomEventsService_DeleteCustomEvent_Handler,
		},
		{
			MethodName: "GetAllCustomEvents",
			Handler:    _CustomEventsService_GetAllCustomEvents_Handler,
		},
		{
			MethodName: "GetByIdCustomEvent",
			Handler:    _CustomEventsService_GetByIdCustomEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customEvents.proto",
}
