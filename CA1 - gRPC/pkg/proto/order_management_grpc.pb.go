// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: proto/order_management.proto

package proto

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
	OrderManagement_GetOrderUnary_FullMethodName        = "/OrderManagement/GetOrderUnary"
	OrderManagement_GetOrderServerStream_FullMethodName = "/OrderManagement/GetOrderServerStream"
	OrderManagement_GetOrderClientStream_FullMethodName = "/OrderManagement/GetOrderClientStream"
	OrderManagement_GetOrderBiDiStream_FullMethodName   = "/OrderManagement/GetOrderBiDiStream"
)

// OrderManagementClient is the client API for OrderManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderManagementClient interface {
	GetOrderUnary(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetOrderServerStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (OrderManagement_GetOrderServerStreamClient, error)
	GetOrderClientStream(ctx context.Context, opts ...grpc.CallOption) (OrderManagement_GetOrderClientStreamClient, error)
	GetOrderBiDiStream(ctx context.Context, opts ...grpc.CallOption) (OrderManagement_GetOrderBiDiStreamClient, error)
}

type orderManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderManagementClient(cc grpc.ClientConnInterface) OrderManagementClient {
	return &orderManagementClient{cc}
}

func (c *orderManagementClient) GetOrderUnary(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, OrderManagement_GetOrderUnary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementClient) GetOrderServerStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (OrderManagement_GetOrderServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderManagement_ServiceDesc.Streams[0], OrderManagement_GetOrderServerStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &orderManagementGetOrderServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderManagement_GetOrderServerStreamClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type orderManagementGetOrderServerStreamClient struct {
	grpc.ClientStream
}

func (x *orderManagementGetOrderServerStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderManagementClient) GetOrderClientStream(ctx context.Context, opts ...grpc.CallOption) (OrderManagement_GetOrderClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderManagement_ServiceDesc.Streams[1], OrderManagement_GetOrderClientStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &orderManagementGetOrderClientStreamClient{stream}
	return x, nil
}

type OrderManagement_GetOrderClientStreamClient interface {
	Send(*Request) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type orderManagementGetOrderClientStreamClient struct {
	grpc.ClientStream
}

func (x *orderManagementGetOrderClientStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderManagementGetOrderClientStreamClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderManagementClient) GetOrderBiDiStream(ctx context.Context, opts ...grpc.CallOption) (OrderManagement_GetOrderBiDiStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderManagement_ServiceDesc.Streams[2], OrderManagement_GetOrderBiDiStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &orderManagementGetOrderBiDiStreamClient{stream}
	return x, nil
}

type OrderManagement_GetOrderBiDiStreamClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type orderManagementGetOrderBiDiStreamClient struct {
	grpc.ClientStream
}

func (x *orderManagementGetOrderBiDiStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderManagementGetOrderBiDiStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderManagementServer is the server API for OrderManagement service.
// All implementations must embed UnimplementedOrderManagementServer
// for forward compatibility
type OrderManagementServer interface {
	GetOrderUnary(context.Context, *Request) (*Response, error)
	GetOrderServerStream(*Request, OrderManagement_GetOrderServerStreamServer) error
	GetOrderClientStream(OrderManagement_GetOrderClientStreamServer) error
	GetOrderBiDiStream(OrderManagement_GetOrderBiDiStreamServer) error
	mustEmbedUnimplementedOrderManagementServer()
}

// UnimplementedOrderManagementServer must be embedded to have forward compatible implementations.
type UnimplementedOrderManagementServer struct {
}

func (UnimplementedOrderManagementServer) GetOrderUnary(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderUnary not implemented")
}
func (UnimplementedOrderManagementServer) GetOrderServerStream(*Request, OrderManagement_GetOrderServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetOrderServerStream not implemented")
}
func (UnimplementedOrderManagementServer) GetOrderClientStream(OrderManagement_GetOrderClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetOrderClientStream not implemented")
}
func (UnimplementedOrderManagementServer) GetOrderBiDiStream(OrderManagement_GetOrderBiDiStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetOrderBiDiStream not implemented")
}
func (UnimplementedOrderManagementServer) mustEmbedUnimplementedOrderManagementServer() {}

// UnsafeOrderManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderManagementServer will
// result in compilation errors.
type UnsafeOrderManagementServer interface {
	mustEmbedUnimplementedOrderManagementServer()
}

func RegisterOrderManagementServer(s grpc.ServiceRegistrar, srv OrderManagementServer) {
	s.RegisterService(&OrderManagement_ServiceDesc, srv)
}

func _OrderManagement_GetOrderUnary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServer).GetOrderUnary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderManagement_GetOrderUnary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServer).GetOrderUnary(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagement_GetOrderServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderManagementServer).GetOrderServerStream(m, &orderManagementGetOrderServerStreamServer{stream})
}

type OrderManagement_GetOrderServerStreamServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type orderManagementGetOrderServerStreamServer struct {
	grpc.ServerStream
}

func (x *orderManagementGetOrderServerStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _OrderManagement_GetOrderClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderManagementServer).GetOrderClientStream(&orderManagementGetOrderClientStreamServer{stream})
}

type OrderManagement_GetOrderClientStreamServer interface {
	SendAndClose(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type orderManagementGetOrderClientStreamServer struct {
	grpc.ServerStream
}

func (x *orderManagementGetOrderClientStreamServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderManagementGetOrderClientStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _OrderManagement_GetOrderBiDiStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderManagementServer).GetOrderBiDiStream(&orderManagementGetOrderBiDiStreamServer{stream})
}

type OrderManagement_GetOrderBiDiStreamServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type orderManagementGetOrderBiDiStreamServer struct {
	grpc.ServerStream
}

func (x *orderManagementGetOrderBiDiStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderManagementGetOrderBiDiStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderManagement_ServiceDesc is the grpc.ServiceDesc for OrderManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OrderManagement",
	HandlerType: (*OrderManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrderUnary",
			Handler:    _OrderManagement_GetOrderUnary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetOrderServerStream",
			Handler:       _OrderManagement_GetOrderServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetOrderClientStream",
			Handler:       _OrderManagement_GetOrderClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetOrderBiDiStream",
			Handler:       _OrderManagement_GetOrderBiDiStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/order_management.proto",
}