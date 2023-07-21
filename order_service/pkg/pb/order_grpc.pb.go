// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/proto/order.proto

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

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	AddtoOrders(ctx context.Context, in *AddOrderRequest, opts ...grpc.CallOption) (*AddOrderResponse, error)
	GetOrders(ctx context.Context, in *ShowOrdersRequest, opts ...grpc.CallOption) (*ShowOrdersResponse, error)
	GetOrderDetails(ctx context.Context, in *ViewOrderDetailResquest, opts ...grpc.CallOption) (*ShowOrderDetailResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) AddtoOrders(ctx context.Context, in *AddOrderRequest, opts ...grpc.CallOption) (*AddOrderResponse, error) {
	out := new(AddOrderResponse)
	err := c.cc.Invoke(ctx, "/pb.OrderService/AddtoOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrders(ctx context.Context, in *ShowOrdersRequest, opts ...grpc.CallOption) (*ShowOrdersResponse, error) {
	out := new(ShowOrdersResponse)
	err := c.cc.Invoke(ctx, "/pb.OrderService/GetOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrderDetails(ctx context.Context, in *ViewOrderDetailResquest, opts ...grpc.CallOption) (*ShowOrderDetailResponse, error) {
	out := new(ShowOrderDetailResponse)
	err := c.cc.Invoke(ctx, "/pb.OrderService/GetOrderDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	AddtoOrders(context.Context, *AddOrderRequest) (*AddOrderResponse, error)
	GetOrders(context.Context, *ShowOrdersRequest) (*ShowOrdersResponse, error)
	GetOrderDetails(context.Context, *ViewOrderDetailResquest) (*ShowOrderDetailResponse, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) AddtoOrders(context.Context, *AddOrderRequest) (*AddOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddtoOrders not implemented")
}
func (UnimplementedOrderServiceServer) GetOrders(context.Context, *ShowOrdersRequest) (*ShowOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedOrderServiceServer) GetOrderDetails(context.Context, *ViewOrderDetailResquest) (*ShowOrderDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderDetails not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_AddtoOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).AddtoOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OrderService/AddtoOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).AddtoOrders(ctx, req.(*AddOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OrderService/GetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrders(ctx, req.(*ShowOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrderDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewOrderDetailResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrderDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OrderService/GetOrderDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrderDetails(ctx, req.(*ViewOrderDetailResquest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddtoOrders",
			Handler:    _OrderService_AddtoOrders_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _OrderService_GetOrders_Handler,
		},
		{
			MethodName: "GetOrderDetails",
			Handler:    _OrderService_GetOrderDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/order.proto",
}
