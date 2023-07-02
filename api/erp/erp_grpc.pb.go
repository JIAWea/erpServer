// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.3
// source: erp.proto

package erp

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
	Erp_CreateErp_FullMethodName = "/erp.erp/CreateErp"
	Erp_UpdateErp_FullMethodName = "/erp.erp/UpdateErp"
	Erp_DeleteErp_FullMethodName = "/erp.erp/DeleteErp"
	Erp_GetErp_FullMethodName    = "/erp.erp/GetErp"
	Erp_ListErp_FullMethodName   = "/erp.erp/ListErp"
	Erp_DDListErp_FullMethodName = "/erp.erp/DDListErp"
)

// ErpClient is the client API for Erp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ErpClient interface {
	CreateErp(ctx context.Context, in *CreateErpReq, opts ...grpc.CallOption) (*CreateErpRsp, error)
	UpdateErp(ctx context.Context, in *UpdateErpReq, opts ...grpc.CallOption) (*UpdateErpRsp, error)
	DeleteErp(ctx context.Context, in *DeleteErpReq, opts ...grpc.CallOption) (*DeleteErpRsp, error)
	GetErp(ctx context.Context, in *GetErpReq, opts ...grpc.CallOption) (*GetErpRsp, error)
	ListErp(ctx context.Context, in *ListErpReq, opts ...grpc.CallOption) (*ListErpRsp, error)
	DDListErp(ctx context.Context, in *ListErpReq, opts ...grpc.CallOption) (*ListErpRsp, error)
}

type erpClient struct {
	cc grpc.ClientConnInterface
}

func NewErpClient(cc grpc.ClientConnInterface) ErpClient {
	return &erpClient{cc}
}

func (c *erpClient) CreateErp(ctx context.Context, in *CreateErpReq, opts ...grpc.CallOption) (*CreateErpRsp, error) {
	out := new(CreateErpRsp)
	err := c.cc.Invoke(ctx, Erp_CreateErp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) UpdateErp(ctx context.Context, in *UpdateErpReq, opts ...grpc.CallOption) (*UpdateErpRsp, error) {
	out := new(UpdateErpRsp)
	err := c.cc.Invoke(ctx, Erp_UpdateErp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) DeleteErp(ctx context.Context, in *DeleteErpReq, opts ...grpc.CallOption) (*DeleteErpRsp, error) {
	out := new(DeleteErpRsp)
	err := c.cc.Invoke(ctx, Erp_DeleteErp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) GetErp(ctx context.Context, in *GetErpReq, opts ...grpc.CallOption) (*GetErpRsp, error) {
	out := new(GetErpRsp)
	err := c.cc.Invoke(ctx, Erp_GetErp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) ListErp(ctx context.Context, in *ListErpReq, opts ...grpc.CallOption) (*ListErpRsp, error) {
	out := new(ListErpRsp)
	err := c.cc.Invoke(ctx, Erp_ListErp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) DDListErp(ctx context.Context, in *ListErpReq, opts ...grpc.CallOption) (*ListErpRsp, error) {
	out := new(ListErpRsp)
	err := c.cc.Invoke(ctx, Erp_DDListErp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ErpServer is the server API for Erp service.
// All implementations must embed UnimplementedErpServer
// for forward compatibility
type ErpServer interface {
	CreateErp(context.Context, *CreateErpReq) (*CreateErpRsp, error)
	UpdateErp(context.Context, *UpdateErpReq) (*UpdateErpRsp, error)
	DeleteErp(context.Context, *DeleteErpReq) (*DeleteErpRsp, error)
	GetErp(context.Context, *GetErpReq) (*GetErpRsp, error)
	ListErp(context.Context, *ListErpReq) (*ListErpRsp, error)
	DDListErp(context.Context, *ListErpReq) (*ListErpRsp, error)
	mustEmbedUnimplementedErpServer()
}

// UnimplementedErpServer must be embedded to have forward compatible implementations.
type UnimplementedErpServer struct {
}

func (UnimplementedErpServer) CreateErp(context.Context, *CreateErpReq) (*CreateErpRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateErp not implemented")
}
func (UnimplementedErpServer) UpdateErp(context.Context, *UpdateErpReq) (*UpdateErpRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateErp not implemented")
}
func (UnimplementedErpServer) DeleteErp(context.Context, *DeleteErpReq) (*DeleteErpRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteErp not implemented")
}
func (UnimplementedErpServer) GetErp(context.Context, *GetErpReq) (*GetErpRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetErp not implemented")
}
func (UnimplementedErpServer) ListErp(context.Context, *ListErpReq) (*ListErpRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListErp not implemented")
}
func (UnimplementedErpServer) DDListErp(context.Context, *ListErpReq) (*ListErpRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DDListErp not implemented")
}
func (UnimplementedErpServer) mustEmbedUnimplementedErpServer() {}

// UnsafeErpServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ErpServer will
// result in compilation errors.
type UnsafeErpServer interface {
	mustEmbedUnimplementedErpServer()
}

func RegisterErpServer(s grpc.ServiceRegistrar, srv ErpServer) {
	s.RegisterService(&Erp_ServiceDesc, srv)
}

func _Erp_CreateErp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateErpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).CreateErp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_CreateErp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).CreateErp(ctx, req.(*CreateErpReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_UpdateErp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateErpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).UpdateErp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_UpdateErp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).UpdateErp(ctx, req.(*UpdateErpReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_DeleteErp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteErpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).DeleteErp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_DeleteErp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).DeleteErp(ctx, req.(*DeleteErpReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_GetErp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetErpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).GetErp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_GetErp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).GetErp(ctx, req.(*GetErpReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_ListErp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListErpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).ListErp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_ListErp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).ListErp(ctx, req.(*ListErpReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_DDListErp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListErpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).DDListErp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_DDListErp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).DDListErp(ctx, req.(*ListErpReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Erp_ServiceDesc is the grpc.ServiceDesc for Erp service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Erp_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "erp.erp",
	HandlerType: (*ErpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateErp",
			Handler:    _Erp_CreateErp_Handler,
		},
		{
			MethodName: "UpdateErp",
			Handler:    _Erp_UpdateErp_Handler,
		},
		{
			MethodName: "DeleteErp",
			Handler:    _Erp_DeleteErp_Handler,
		},
		{
			MethodName: "GetErp",
			Handler:    _Erp_GetErp_Handler,
		},
		{
			MethodName: "ListErp",
			Handler:    _Erp_ListErp_Handler,
		},
		{
			MethodName: "DDListErp",
			Handler:    _Erp_DDListErp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "erp.proto",
}
