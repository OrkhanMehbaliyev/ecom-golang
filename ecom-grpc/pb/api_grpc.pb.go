// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.0
// source: api.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Ecom_CreateProduct_FullMethodName           = "/pb.ecom/CreateProduct"
	Ecom_GetProduct_FullMethodName              = "/pb.ecom/GetProduct"
	Ecom_ListProducts_FullMethodName            = "/pb.ecom/ListProducts"
	Ecom_UpdateProduct_FullMethodName           = "/pb.ecom/UpdateProduct"
	Ecom_DeleteProduct_FullMethodName           = "/pb.ecom/DeleteProduct"
	Ecom_CreateOrder_FullMethodName             = "/pb.ecom/CreateOrder"
	Ecom_GetOrder_FullMethodName                = "/pb.ecom/GetOrder"
	Ecom_ListOrders_FullMethodName              = "/pb.ecom/ListOrders"
	Ecom_UpdateOrderStatus_FullMethodName       = "/pb.ecom/UpdateOrderStatus"
	Ecom_DeleteOrder_FullMethodName             = "/pb.ecom/DeleteOrder"
	Ecom_CreateUser_FullMethodName              = "/pb.ecom/CreateUser"
	Ecom_GetUser_FullMethodName                 = "/pb.ecom/GetUser"
	Ecom_ListUsers_FullMethodName               = "/pb.ecom/ListUsers"
	Ecom_UpdateUser_FullMethodName              = "/pb.ecom/UpdateUser"
	Ecom_DeleteUser_FullMethodName              = "/pb.ecom/DeleteUser"
	Ecom_CreateSession_FullMethodName           = "/pb.ecom/CreateSession"
	Ecom_GetSession_FullMethodName              = "/pb.ecom/GetSession"
	Ecom_RevokeSession_FullMethodName           = "/pb.ecom/RevokeSession"
	Ecom_DeleteSession_FullMethodName           = "/pb.ecom/DeleteSession"
	Ecom_ListNotificationEvents_FullMethodName  = "/pb.ecom/ListNotificationEvents"
	Ecom_UpdateNotificationEvent_FullMethodName = "/pb.ecom/UpdateNotificationEvent"
)

// EcomClient is the client API for Ecom service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EcomClient interface {
	CreateProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error)
	GetProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error)
	ListProducts(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ListProductRes, error)
	UpdateProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error)
	DeleteProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error)
	CreateOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error)
	GetOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error)
	ListOrders(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*ListOrderRes, error)
	UpdateOrderStatus(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error)
	DeleteOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error)
	CreateUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error)
	GetUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error)
	ListUsers(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*ListUserRes, error)
	UpdateUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error)
	DeleteUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error)
	CreateSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error)
	GetSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error)
	RevokeSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error)
	DeleteSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error)
	ListNotificationEvents(ctx context.Context, in *ListNotificationEventsReq, opts ...grpc.CallOption) (*ListNotificationEventsRes, error)
	UpdateNotificationEvent(ctx context.Context, in *UpdateNotificationEventReq, opts ...grpc.CallOption) (*UpdateNotificationEventRes, error)
}

type ecomClient struct {
	cc grpc.ClientConnInterface
}

func NewEcomClient(cc grpc.ClientConnInterface) EcomClient {
	return &ecomClient{cc}
}

func (c *ecomClient) CreateProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductRes)
	err := c.cc.Invoke(ctx, Ecom_CreateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomClient) GetProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductRes)
	err := c.cc.Invoke(ctx, Ecom_GetProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomClient) ListProducts(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ListProductRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProductRes)
	err := c.cc.Invoke(ctx, Ecom_ListProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomClient) UpdateProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductRes)
	err := c.cc.Invoke(ctx, Ecom_UpdateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomClient) DeleteProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductRes)
	err := c.cc.Invoke(ctx, Ecom_DeleteProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomClient) CreateOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderRes)
	err := c.cc.Invoke(ctx, Ecom_CreateOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomClient) GetOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderRes)
	err := c.cc.Invoke(ctx, Ecom_GetOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomClient) ListOrders(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*ListOrderRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListOrderRes)
	err := c.cc.Invoke(ctx, Ecom_ListOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomClient) UpdateOrderStatus(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderRes)
	err := c.cc.Invoke(ctx, Ecom_UpdateOrderStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomClient) DeleteOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderRes)
	err := c.cc.Invoke(ctx, Ecom_DeleteOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomClient) CreateUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserRes)
	err := c.cc.Invoke(ctx, Ecom_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomClient) GetUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserRes)
	err := c.cc.Invoke(ctx, Ecom_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomClient) ListUsers(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*ListUserRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUserRes)
	err := c.cc.Invoke(ctx, Ecom_ListUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomClient) UpdateUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserRes)
	err := c.cc.Invoke(ctx, Ecom_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomClient) DeleteUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserRes)
	err := c.cc.Invoke(ctx, Ecom_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomClient) CreateSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SessionRes)
	err := c.cc.Invoke(ctx, Ecom_CreateSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomClient) GetSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SessionRes)
	err := c.cc.Invoke(ctx, Ecom_GetSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomClient) RevokeSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SessionRes)
	err := c.cc.Invoke(ctx, Ecom_RevokeSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomClient) DeleteSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SessionRes)
	err := c.cc.Invoke(ctx, Ecom_DeleteSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomClient) ListNotificationEvents(ctx context.Context, in *ListNotificationEventsReq, opts ...grpc.CallOption) (*ListNotificationEventsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNotificationEventsRes)
	err := c.cc.Invoke(ctx, Ecom_ListNotificationEvents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomClient) UpdateNotificationEvent(ctx context.Context, in *UpdateNotificationEventReq, opts ...grpc.CallOption) (*UpdateNotificationEventRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateNotificationEventRes)
	err := c.cc.Invoke(ctx, Ecom_UpdateNotificationEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EcomServer is the server API for Ecom service.
// All implementations must embed UnimplementedEcomServer
// for forward compatibility.
type EcomServer interface {
	CreateProduct(context.Context, *ProductReq) (*ProductRes, error)
	GetProduct(context.Context, *ProductReq) (*ProductRes, error)
	ListProducts(context.Context, *ProductReq) (*ListProductRes, error)
	UpdateProduct(context.Context, *ProductReq) (*ProductRes, error)
	DeleteProduct(context.Context, *ProductReq) (*ProductRes, error)
	CreateOrder(context.Context, *OrderReq) (*OrderRes, error)
	GetOrder(context.Context, *OrderReq) (*OrderRes, error)
	ListOrders(context.Context, *OrderReq) (*ListOrderRes, error)
	UpdateOrderStatus(context.Context, *OrderReq) (*OrderRes, error)
	DeleteOrder(context.Context, *OrderReq) (*OrderRes, error)
	CreateUser(context.Context, *UserReq) (*UserRes, error)
	GetUser(context.Context, *UserReq) (*UserRes, error)
	ListUsers(context.Context, *UserReq) (*ListUserRes, error)
	UpdateUser(context.Context, *UserReq) (*UserRes, error)
	DeleteUser(context.Context, *UserReq) (*UserRes, error)
	CreateSession(context.Context, *SessionReq) (*SessionRes, error)
	GetSession(context.Context, *SessionReq) (*SessionRes, error)
	RevokeSession(context.Context, *SessionReq) (*SessionRes, error)
	DeleteSession(context.Context, *SessionReq) (*SessionRes, error)
	ListNotificationEvents(context.Context, *ListNotificationEventsReq) (*ListNotificationEventsRes, error)
	UpdateNotificationEvent(context.Context, *UpdateNotificationEventReq) (*UpdateNotificationEventRes, error)
	mustEmbedUnimplementedEcomServer()
}

// UnimplementedEcomServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEcomServer struct{}

func (UnimplementedEcomServer) CreateProduct(context.Context, *ProductReq) (*ProductRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedEcomServer) GetProduct(context.Context, *ProductReq) (*ProductRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedEcomServer) ListProducts(context.Context, *ProductReq) (*ListProductRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
}
func (UnimplementedEcomServer) UpdateProduct(context.Context, *ProductReq) (*ProductRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedEcomServer) DeleteProduct(context.Context, *ProductReq) (*ProductRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedEcomServer) CreateOrder(context.Context, *OrderReq) (*OrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedEcomServer) GetOrder(context.Context, *OrderReq) (*OrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedEcomServer) ListOrders(context.Context, *OrderReq) (*ListOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}
func (UnimplementedEcomServer) UpdateOrderStatus(context.Context, *OrderReq) (*OrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderStatus not implemented")
}
func (UnimplementedEcomServer) DeleteOrder(context.Context, *OrderReq) (*OrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
func (UnimplementedEcomServer) CreateUser(context.Context, *UserReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedEcomServer) GetUser(context.Context, *UserReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedEcomServer) ListUsers(context.Context, *UserReq) (*ListUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedEcomServer) UpdateUser(context.Context, *UserReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedEcomServer) DeleteUser(context.Context, *UserReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedEcomServer) CreateSession(context.Context, *SessionReq) (*SessionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedEcomServer) GetSession(context.Context, *SessionReq) (*SessionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
}
func (UnimplementedEcomServer) RevokeSession(context.Context, *SessionReq) (*SessionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeSession not implemented")
}
func (UnimplementedEcomServer) DeleteSession(context.Context, *SessionReq) (*SessionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSession not implemented")
}
func (UnimplementedEcomServer) ListNotificationEvents(context.Context, *ListNotificationEventsReq) (*ListNotificationEventsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNotificationEvents not implemented")
}
func (UnimplementedEcomServer) UpdateNotificationEvent(context.Context, *UpdateNotificationEventReq) (*UpdateNotificationEventRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNotificationEvent not implemented")
}
func (UnimplementedEcomServer) mustEmbedUnimplementedEcomServer() {}
func (UnimplementedEcomServer) testEmbeddedByValue()              {}

// UnsafeEcomServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EcomServer will
// result in compilation errors.
type UnsafeEcomServer interface {
	mustEmbedUnimplementedEcomServer()
}

func RegisterEcomServer(s grpc.ServiceRegistrar, srv EcomServer) {
	// If the following call pancis, it indicates UnimplementedEcomServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Ecom_ServiceDesc, srv)
}

func _Ecom_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecom_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).CreateProduct(ctx, req.(*ProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecom_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecom_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).GetProduct(ctx, req.(*ProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecom_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecom_ListProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).ListProducts(ctx, req.(*ProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecom_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecom_UpdateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).UpdateProduct(ctx, req.(*ProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecom_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecom_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).DeleteProduct(ctx, req.(*ProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecom_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecom_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).CreateOrder(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecom_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecom_GetOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).GetOrder(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecom_ListOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).ListOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecom_ListOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).ListOrders(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecom_UpdateOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).UpdateOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecom_UpdateOrderStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).UpdateOrderStatus(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecom_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecom_DeleteOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).DeleteOrder(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecom_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecom_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).CreateUser(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecom_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecom_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).GetUser(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecom_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecom_ListUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).ListUsers(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecom_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecom_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).UpdateUser(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecom_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecom_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).DeleteUser(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecom_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecom_CreateSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).CreateSession(ctx, req.(*SessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecom_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecom_GetSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).GetSession(ctx, req.(*SessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecom_RevokeSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).RevokeSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecom_RevokeSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).RevokeSession(ctx, req.(*SessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecom_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecom_DeleteSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).DeleteSession(ctx, req.(*SessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecom_ListNotificationEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNotificationEventsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).ListNotificationEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecom_ListNotificationEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).ListNotificationEvents(ctx, req.(*ListNotificationEventsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecom_UpdateNotificationEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNotificationEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomServer).UpdateNotificationEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ecom_UpdateNotificationEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomServer).UpdateNotificationEvent(ctx, req.(*UpdateNotificationEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Ecom_ServiceDesc is the grpc.ServiceDesc for Ecom service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ecom_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ecom",
	HandlerType: (*EcomServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _Ecom_CreateProduct_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _Ecom_GetProduct_Handler,
		},
		{
			MethodName: "ListProducts",
			Handler:    _Ecom_ListProducts_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _Ecom_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _Ecom_DeleteProduct_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _Ecom_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _Ecom_GetOrder_Handler,
		},
		{
			MethodName: "ListOrders",
			Handler:    _Ecom_ListOrders_Handler,
		},
		{
			MethodName: "UpdateOrderStatus",
			Handler:    _Ecom_UpdateOrderStatus_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _Ecom_DeleteOrder_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Ecom_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Ecom_GetUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _Ecom_ListUsers_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Ecom_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Ecom_DeleteUser_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _Ecom_CreateSession_Handler,
		},
		{
			MethodName: "GetSession",
			Handler:    _Ecom_GetSession_Handler,
		},
		{
			MethodName: "RevokeSession",
			Handler:    _Ecom_RevokeSession_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _Ecom_DeleteSession_Handler,
		},
		{
			MethodName: "ListNotificationEvents",
			Handler:    _Ecom_ListNotificationEvents_Handler,
		},
		{
			MethodName: "UpdateNotificationEvent",
			Handler:    _Ecom_UpdateNotificationEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
