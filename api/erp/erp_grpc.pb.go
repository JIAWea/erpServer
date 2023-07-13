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
	Erp_UserLogin_FullMethodName         = "/erp.erp/UserLogin"
	Erp_UserLogout_FullMethodName        = "/erp.erp/UserLogout"
	Erp_GetUserInfo_FullMethodName       = "/erp.erp/GetUserInfo"
	Erp_CreateUser_FullMethodName        = "/erp.erp/CreateUser"
	Erp_UpdateUser_FullMethodName        = "/erp.erp/UpdateUser"
	Erp_UpdateUserRole_FullMethodName    = "/erp.erp/UpdateUserRole"
	Erp_UpdatePassword_FullMethodName    = "/erp.erp/UpdatePassword"
	Erp_DeleteUser_FullMethodName        = "/erp.erp/DeleteUser"
	Erp_GetUser_FullMethodName           = "/erp.erp/GetUser"
	Erp_ListUser_FullMethodName          = "/erp.erp/ListUser"
	Erp_CreateRole_FullMethodName        = "/erp.erp/CreateRole"
	Erp_UpdateRole_FullMethodName        = "/erp.erp/UpdateRole"
	Erp_DeleteRole_FullMethodName        = "/erp.erp/DeleteRole"
	Erp_GetRole_FullMethodName           = "/erp.erp/GetRole"
	Erp_ListRole_FullMethodName          = "/erp.erp/ListRole"
	Erp_ListMenu_FullMethodName          = "/erp.erp/ListMenu"
	Erp_ListMenuTree_FullMethodName      = "/erp.erp/ListMenuTree"
	Erp_GetRoleMenuIdList_FullMethodName = "/erp.erp/GetRoleMenuIdList"
	Erp_UpdateRoleMenu_FullMethodName    = "/erp.erp/UpdateRoleMenu"
	Erp_CreateAccount_FullMethodName     = "/erp.erp/CreateAccount"
	Erp_UpdateAccount_FullMethodName     = "/erp.erp/UpdateAccount"
	Erp_DeleteAccount_FullMethodName     = "/erp.erp/DeleteAccount"
	Erp_GetAccount_FullMethodName        = "/erp.erp/GetAccount"
	Erp_ListAccount_FullMethodName       = "/erp.erp/ListAccount"
)

// ErpClient is the client API for Erp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ErpClient interface {
	UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginRsp, error)
	UserLogout(ctx context.Context, in *UserLogoutReq, opts ...grpc.CallOption) (*UserLogoutRsp, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoRsp, error)
	CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserRsp, error)
	UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserRsp, error)
	UpdateUserRole(ctx context.Context, in *UpdateUserRoleReq, opts ...grpc.CallOption) (*UpdateUserRoleRsp, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*UpdatePasswordRsp, error)
	DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserRsp, error)
	GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserRsp, error)
	ListUser(ctx context.Context, in *ListUserReq, opts ...grpc.CallOption) (*ListUserRsp, error)
	CreateRole(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*CreateRoleRsp, error)
	UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*UpdateRoleRsp, error)
	DeleteRole(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*DeleteRoleRsp, error)
	GetRole(ctx context.Context, in *GetRoleReq, opts ...grpc.CallOption) (*GetRoleRsp, error)
	ListRole(ctx context.Context, in *ListRoleReq, opts ...grpc.CallOption) (*ListRoleRsp, error)
	ListMenu(ctx context.Context, in *ListMenuReq, opts ...grpc.CallOption) (*ListMenuRsp, error)
	ListMenuTree(ctx context.Context, in *ListMenuTreeReq, opts ...grpc.CallOption) (*ListMenuTreeRsp, error)
	GetRoleMenuIdList(ctx context.Context, in *GetRoleMenuIdListReq, opts ...grpc.CallOption) (*GetRoleMenuIdListRsp, error)
	UpdateRoleMenu(ctx context.Context, in *UpdateRoleMenuReq, opts ...grpc.CallOption) (*UpdateRoleMenuRsp, error)
	CreateAccount(ctx context.Context, in *CreateAccountReq, opts ...grpc.CallOption) (*CreateAccountRsp, error)
	UpdateAccount(ctx context.Context, in *UpdateAccountReq, opts ...grpc.CallOption) (*UpdateAccountRsp, error)
	DeleteAccount(ctx context.Context, in *DeleteAccountReq, opts ...grpc.CallOption) (*DeleteAccountRsp, error)
	GetAccount(ctx context.Context, in *GetAccountReq, opts ...grpc.CallOption) (*GetAccountRsp, error)
	ListAccount(ctx context.Context, in *ListAccountReq, opts ...grpc.CallOption) (*ListAccountRsp, error)
}

type erpClient struct {
	cc grpc.ClientConnInterface
}

func NewErpClient(cc grpc.ClientConnInterface) ErpClient {
	return &erpClient{cc}
}

func (c *erpClient) UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginRsp, error) {
	out := new(UserLoginRsp)
	err := c.cc.Invoke(ctx, Erp_UserLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) UserLogout(ctx context.Context, in *UserLogoutReq, opts ...grpc.CallOption) (*UserLogoutRsp, error) {
	out := new(UserLogoutRsp)
	err := c.cc.Invoke(ctx, Erp_UserLogout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoRsp, error) {
	out := new(GetUserInfoRsp)
	err := c.cc.Invoke(ctx, Erp_GetUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserRsp, error) {
	out := new(CreateUserRsp)
	err := c.cc.Invoke(ctx, Erp_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserRsp, error) {
	out := new(UpdateUserRsp)
	err := c.cc.Invoke(ctx, Erp_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) UpdateUserRole(ctx context.Context, in *UpdateUserRoleReq, opts ...grpc.CallOption) (*UpdateUserRoleRsp, error) {
	out := new(UpdateUserRoleRsp)
	err := c.cc.Invoke(ctx, Erp_UpdateUserRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*UpdatePasswordRsp, error) {
	out := new(UpdatePasswordRsp)
	err := c.cc.Invoke(ctx, Erp_UpdatePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserRsp, error) {
	out := new(DeleteUserRsp)
	err := c.cc.Invoke(ctx, Erp_DeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserRsp, error) {
	out := new(GetUserRsp)
	err := c.cc.Invoke(ctx, Erp_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) ListUser(ctx context.Context, in *ListUserReq, opts ...grpc.CallOption) (*ListUserRsp, error) {
	out := new(ListUserRsp)
	err := c.cc.Invoke(ctx, Erp_ListUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) CreateRole(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*CreateRoleRsp, error) {
	out := new(CreateRoleRsp)
	err := c.cc.Invoke(ctx, Erp_CreateRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*UpdateRoleRsp, error) {
	out := new(UpdateRoleRsp)
	err := c.cc.Invoke(ctx, Erp_UpdateRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) DeleteRole(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*DeleteRoleRsp, error) {
	out := new(DeleteRoleRsp)
	err := c.cc.Invoke(ctx, Erp_DeleteRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) GetRole(ctx context.Context, in *GetRoleReq, opts ...grpc.CallOption) (*GetRoleRsp, error) {
	out := new(GetRoleRsp)
	err := c.cc.Invoke(ctx, Erp_GetRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) ListRole(ctx context.Context, in *ListRoleReq, opts ...grpc.CallOption) (*ListRoleRsp, error) {
	out := new(ListRoleRsp)
	err := c.cc.Invoke(ctx, Erp_ListRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) ListMenu(ctx context.Context, in *ListMenuReq, opts ...grpc.CallOption) (*ListMenuRsp, error) {
	out := new(ListMenuRsp)
	err := c.cc.Invoke(ctx, Erp_ListMenu_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) ListMenuTree(ctx context.Context, in *ListMenuTreeReq, opts ...grpc.CallOption) (*ListMenuTreeRsp, error) {
	out := new(ListMenuTreeRsp)
	err := c.cc.Invoke(ctx, Erp_ListMenuTree_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) GetRoleMenuIdList(ctx context.Context, in *GetRoleMenuIdListReq, opts ...grpc.CallOption) (*GetRoleMenuIdListRsp, error) {
	out := new(GetRoleMenuIdListRsp)
	err := c.cc.Invoke(ctx, Erp_GetRoleMenuIdList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) UpdateRoleMenu(ctx context.Context, in *UpdateRoleMenuReq, opts ...grpc.CallOption) (*UpdateRoleMenuRsp, error) {
	out := new(UpdateRoleMenuRsp)
	err := c.cc.Invoke(ctx, Erp_UpdateRoleMenu_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) CreateAccount(ctx context.Context, in *CreateAccountReq, opts ...grpc.CallOption) (*CreateAccountRsp, error) {
	out := new(CreateAccountRsp)
	err := c.cc.Invoke(ctx, Erp_CreateAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) UpdateAccount(ctx context.Context, in *UpdateAccountReq, opts ...grpc.CallOption) (*UpdateAccountRsp, error) {
	out := new(UpdateAccountRsp)
	err := c.cc.Invoke(ctx, Erp_UpdateAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) DeleteAccount(ctx context.Context, in *DeleteAccountReq, opts ...grpc.CallOption) (*DeleteAccountRsp, error) {
	out := new(DeleteAccountRsp)
	err := c.cc.Invoke(ctx, Erp_DeleteAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) GetAccount(ctx context.Context, in *GetAccountReq, opts ...grpc.CallOption) (*GetAccountRsp, error) {
	out := new(GetAccountRsp)
	err := c.cc.Invoke(ctx, Erp_GetAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *erpClient) ListAccount(ctx context.Context, in *ListAccountReq, opts ...grpc.CallOption) (*ListAccountRsp, error) {
	out := new(ListAccountRsp)
	err := c.cc.Invoke(ctx, Erp_ListAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ErpServer is the server API for Erp service.
// All implementations must embed UnimplementedErpServer
// for forward compatibility
type ErpServer interface {
	UserLogin(context.Context, *UserLoginReq) (*UserLoginRsp, error)
	UserLogout(context.Context, *UserLogoutReq) (*UserLogoutRsp, error)
	GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoRsp, error)
	CreateUser(context.Context, *CreateUserReq) (*CreateUserRsp, error)
	UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserRsp, error)
	UpdateUserRole(context.Context, *UpdateUserRoleReq) (*UpdateUserRoleRsp, error)
	UpdatePassword(context.Context, *UpdatePasswordReq) (*UpdatePasswordRsp, error)
	DeleteUser(context.Context, *DeleteUserReq) (*DeleteUserRsp, error)
	GetUser(context.Context, *GetUserReq) (*GetUserRsp, error)
	ListUser(context.Context, *ListUserReq) (*ListUserRsp, error)
	CreateRole(context.Context, *CreateRoleReq) (*CreateRoleRsp, error)
	UpdateRole(context.Context, *UpdateRoleReq) (*UpdateRoleRsp, error)
	DeleteRole(context.Context, *DeleteRoleReq) (*DeleteRoleRsp, error)
	GetRole(context.Context, *GetRoleReq) (*GetRoleRsp, error)
	ListRole(context.Context, *ListRoleReq) (*ListRoleRsp, error)
	ListMenu(context.Context, *ListMenuReq) (*ListMenuRsp, error)
	ListMenuTree(context.Context, *ListMenuTreeReq) (*ListMenuTreeRsp, error)
	GetRoleMenuIdList(context.Context, *GetRoleMenuIdListReq) (*GetRoleMenuIdListRsp, error)
	UpdateRoleMenu(context.Context, *UpdateRoleMenuReq) (*UpdateRoleMenuRsp, error)
	CreateAccount(context.Context, *CreateAccountReq) (*CreateAccountRsp, error)
	UpdateAccount(context.Context, *UpdateAccountReq) (*UpdateAccountRsp, error)
	DeleteAccount(context.Context, *DeleteAccountReq) (*DeleteAccountRsp, error)
	GetAccount(context.Context, *GetAccountReq) (*GetAccountRsp, error)
	ListAccount(context.Context, *ListAccountReq) (*ListAccountRsp, error)
	mustEmbedUnimplementedErpServer()
}

// UnimplementedErpServer must be embedded to have forward compatible implementations.
type UnimplementedErpServer struct {
}

func (UnimplementedErpServer) UserLogin(context.Context, *UserLoginReq) (*UserLoginRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedErpServer) UserLogout(context.Context, *UserLogoutReq) (*UserLogoutRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogout not implemented")
}
func (UnimplementedErpServer) GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedErpServer) CreateUser(context.Context, *CreateUserReq) (*CreateUserRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedErpServer) UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedErpServer) UpdateUserRole(context.Context, *UpdateUserRoleReq) (*UpdateUserRoleRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserRole not implemented")
}
func (UnimplementedErpServer) UpdatePassword(context.Context, *UpdatePasswordReq) (*UpdatePasswordRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedErpServer) DeleteUser(context.Context, *DeleteUserReq) (*DeleteUserRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedErpServer) GetUser(context.Context, *GetUserReq) (*GetUserRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedErpServer) ListUser(context.Context, *ListUserReq) (*ListUserRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedErpServer) CreateRole(context.Context, *CreateRoleReq) (*CreateRoleRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedErpServer) UpdateRole(context.Context, *UpdateRoleReq) (*UpdateRoleRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedErpServer) DeleteRole(context.Context, *DeleteRoleReq) (*DeleteRoleRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedErpServer) GetRole(context.Context, *GetRoleReq) (*GetRoleRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (UnimplementedErpServer) ListRole(context.Context, *ListRoleReq) (*ListRoleRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRole not implemented")
}
func (UnimplementedErpServer) ListMenu(context.Context, *ListMenuReq) (*ListMenuRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMenu not implemented")
}
func (UnimplementedErpServer) ListMenuTree(context.Context, *ListMenuTreeReq) (*ListMenuTreeRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMenuTree not implemented")
}
func (UnimplementedErpServer) GetRoleMenuIdList(context.Context, *GetRoleMenuIdListReq) (*GetRoleMenuIdListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleMenuIdList not implemented")
}
func (UnimplementedErpServer) UpdateRoleMenu(context.Context, *UpdateRoleMenuReq) (*UpdateRoleMenuRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoleMenu not implemented")
}
func (UnimplementedErpServer) CreateAccount(context.Context, *CreateAccountReq) (*CreateAccountRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedErpServer) UpdateAccount(context.Context, *UpdateAccountReq) (*UpdateAccountRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedErpServer) DeleteAccount(context.Context, *DeleteAccountReq) (*DeleteAccountRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedErpServer) GetAccount(context.Context, *GetAccountReq) (*GetAccountRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedErpServer) ListAccount(context.Context, *ListAccountReq) (*ListAccountRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccount not implemented")
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

func _Erp_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).UserLogin(ctx, req.(*UserLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_UserLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLogoutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).UserLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_UserLogout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).UserLogout(ctx, req.(*UserLogoutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).GetUserInfo(ctx, req.(*GetUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).CreateUser(ctx, req.(*CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).UpdateUser(ctx, req.(*UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_UpdateUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).UpdateUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_UpdateUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).UpdateUserRole(ctx, req.(*UpdateUserRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_UpdatePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).UpdatePassword(ctx, req.(*UpdatePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).DeleteUser(ctx, req.(*DeleteUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).GetUser(ctx, req.(*GetUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_ListUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).ListUser(ctx, req.(*ListUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_CreateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).CreateRole(ctx, req.(*CreateRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_UpdateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).UpdateRole(ctx, req.(*UpdateRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_DeleteRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).DeleteRole(ctx, req.(*DeleteRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_GetRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).GetRole(ctx, req.(*GetRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_ListRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).ListRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_ListRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).ListRole(ctx, req.(*ListRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_ListMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMenuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).ListMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_ListMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).ListMenu(ctx, req.(*ListMenuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_ListMenuTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMenuTreeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).ListMenuTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_ListMenuTree_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).ListMenuTree(ctx, req.(*ListMenuTreeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_GetRoleMenuIdList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleMenuIdListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).GetRoleMenuIdList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_GetRoleMenuIdList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).GetRoleMenuIdList(ctx, req.(*GetRoleMenuIdListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_UpdateRoleMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleMenuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).UpdateRoleMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_UpdateRoleMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).UpdateRoleMenu(ctx, req.(*UpdateRoleMenuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).CreateAccount(ctx, req.(*CreateAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_UpdateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).UpdateAccount(ctx, req.(*UpdateAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_DeleteAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).DeleteAccount(ctx, req.(*DeleteAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_GetAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).GetAccount(ctx, req.(*GetAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Erp_ListAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpServer).ListAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Erp_ListAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpServer).ListAccount(ctx, req.(*ListAccountReq))
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
			MethodName: "UserLogin",
			Handler:    _Erp_UserLogin_Handler,
		},
		{
			MethodName: "UserLogout",
			Handler:    _Erp_UserLogout_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _Erp_GetUserInfo_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Erp_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Erp_UpdateUser_Handler,
		},
		{
			MethodName: "UpdateUserRole",
			Handler:    _Erp_UpdateUserRole_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _Erp_UpdatePassword_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Erp_DeleteUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Erp_GetUser_Handler,
		},
		{
			MethodName: "ListUser",
			Handler:    _Erp_ListUser_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _Erp_CreateRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _Erp_UpdateRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _Erp_DeleteRole_Handler,
		},
		{
			MethodName: "GetRole",
			Handler:    _Erp_GetRole_Handler,
		},
		{
			MethodName: "ListRole",
			Handler:    _Erp_ListRole_Handler,
		},
		{
			MethodName: "ListMenu",
			Handler:    _Erp_ListMenu_Handler,
		},
		{
			MethodName: "ListMenuTree",
			Handler:    _Erp_ListMenuTree_Handler,
		},
		{
			MethodName: "GetRoleMenuIdList",
			Handler:    _Erp_GetRoleMenuIdList_Handler,
		},
		{
			MethodName: "UpdateRoleMenu",
			Handler:    _Erp_UpdateRoleMenu_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _Erp_CreateAccount_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _Erp_UpdateAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _Erp_DeleteAccount_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _Erp_GetAccount_Handler,
		},
		{
			MethodName: "ListAccount",
			Handler:    _Erp_ListAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "erp.proto",
}
