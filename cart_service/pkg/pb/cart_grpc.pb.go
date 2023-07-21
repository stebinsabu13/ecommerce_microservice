// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/proto/cart.proto

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

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartServiceClient interface {
	ViewCart(ctx context.Context, in *ViewCartRequest, opts ...grpc.CallOption) (*ViewCartResponse, error)
	AddtoCart(ctx context.Context, in *AddCartRequest, opts ...grpc.CallOption) (*AddCartResponse, error)
	CreateCart(ctx context.Context, in *CreateCartRequest, opts ...grpc.CallOption) (*CreateCartResponse, error)
	FindCart(ctx context.Context, in *CartRequest, opts ...grpc.CallOption) (*CartResponse, error)
	CartItems(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*CartItemResponse, error)
	UpdateCart(ctx context.Context, in *CartResponse, opts ...grpc.CallOption) (*UpdateResponse, error)
	DeleteCart(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
}

type cartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartServiceClient(cc grpc.ClientConnInterface) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) ViewCart(ctx context.Context, in *ViewCartRequest, opts ...grpc.CallOption) (*ViewCartResponse, error) {
	out := new(ViewCartResponse)
	err := c.cc.Invoke(ctx, "/pb.CartService/ViewCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) AddtoCart(ctx context.Context, in *AddCartRequest, opts ...grpc.CallOption) (*AddCartResponse, error) {
	out := new(AddCartResponse)
	err := c.cc.Invoke(ctx, "/pb.CartService/AddtoCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) CreateCart(ctx context.Context, in *CreateCartRequest, opts ...grpc.CallOption) (*CreateCartResponse, error) {
	out := new(CreateCartResponse)
	err := c.cc.Invoke(ctx, "/pb.CartService/CreateCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) FindCart(ctx context.Context, in *CartRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	out := new(CartResponse)
	err := c.cc.Invoke(ctx, "/pb.CartService/FindCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) CartItems(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*CartItemResponse, error) {
	out := new(CartItemResponse)
	err := c.cc.Invoke(ctx, "/pb.CartService/CartItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) UpdateCart(ctx context.Context, in *CartResponse, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/pb.CartService/UpdateCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) DeleteCart(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/pb.CartService/DeleteCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
// All implementations must embed UnimplementedCartServiceServer
// for forward compatibility
type CartServiceServer interface {
	ViewCart(context.Context, *ViewCartRequest) (*ViewCartResponse, error)
	AddtoCart(context.Context, *AddCartRequest) (*AddCartResponse, error)
	CreateCart(context.Context, *CreateCartRequest) (*CreateCartResponse, error)
	FindCart(context.Context, *CartRequest) (*CartResponse, error)
	CartItems(context.Context, *CartItemRequest) (*CartItemResponse, error)
	UpdateCart(context.Context, *CartResponse) (*UpdateResponse, error)
	DeleteCart(context.Context, *CartItemRequest) (*UpdateResponse, error)
	mustEmbedUnimplementedCartServiceServer()
}

// UnimplementedCartServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCartServiceServer struct {
}

func (UnimplementedCartServiceServer) ViewCart(context.Context, *ViewCartRequest) (*ViewCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewCart not implemented")
}
func (UnimplementedCartServiceServer) AddtoCart(context.Context, *AddCartRequest) (*AddCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddtoCart not implemented")
}
func (UnimplementedCartServiceServer) CreateCart(context.Context, *CreateCartRequest) (*CreateCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCart not implemented")
}
func (UnimplementedCartServiceServer) FindCart(context.Context, *CartRequest) (*CartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCart not implemented")
}
func (UnimplementedCartServiceServer) CartItems(context.Context, *CartItemRequest) (*CartItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CartItems not implemented")
}
func (UnimplementedCartServiceServer) UpdateCart(context.Context, *CartResponse) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCart not implemented")
}
func (UnimplementedCartServiceServer) DeleteCart(context.Context, *CartItemRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCart not implemented")
}
func (UnimplementedCartServiceServer) mustEmbedUnimplementedCartServiceServer() {}

// UnsafeCartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServiceServer will
// result in compilation errors.
type UnsafeCartServiceServer interface {
	mustEmbedUnimplementedCartServiceServer()
}

func RegisterCartServiceServer(s grpc.ServiceRegistrar, srv CartServiceServer) {
	s.RegisterService(&CartService_ServiceDesc, srv)
}

func _CartService_ViewCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).ViewCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CartService/ViewCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).ViewCart(ctx, req.(*ViewCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_AddtoCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).AddtoCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CartService/AddtoCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).AddtoCart(ctx, req.(*AddCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_CreateCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).CreateCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CartService/CreateCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).CreateCart(ctx, req.(*CreateCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_FindCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).FindCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CartService/FindCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).FindCart(ctx, req.(*CartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_CartItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).CartItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CartService/CartItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).CartItems(ctx, req.(*CartItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_UpdateCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).UpdateCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CartService/UpdateCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).UpdateCart(ctx, req.(*CartResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_DeleteCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).DeleteCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CartService/DeleteCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).DeleteCart(ctx, req.(*CartItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CartService_ServiceDesc is the grpc.ServiceDesc for CartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ViewCart",
			Handler:    _CartService_ViewCart_Handler,
		},
		{
			MethodName: "AddtoCart",
			Handler:    _CartService_AddtoCart_Handler,
		},
		{
			MethodName: "CreateCart",
			Handler:    _CartService_CreateCart_Handler,
		},
		{
			MethodName: "FindCart",
			Handler:    _CartService_FindCart_Handler,
		},
		{
			MethodName: "CartItems",
			Handler:    _CartService_CartItems_Handler,
		},
		{
			MethodName: "UpdateCart",
			Handler:    _CartService_UpdateCart_Handler,
		},
		{
			MethodName: "DeleteCart",
			Handler:    _CartService_DeleteCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/cart.proto",
}
