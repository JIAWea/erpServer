package erp

import (
	"context"
)

const ClientName = "erp"



func CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.CreateUser(ctx, req)
}

func UpdateUser(ctx context.Context, req *UpdateUserReq) (*UpdateUserRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.UpdateUser(ctx, req)
}

func UpdatePassword(ctx context.Context, req *UpdatePasswordReq) (*UpdatePasswordRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.UpdatePassword(ctx, req)
}

func DeleteUser(ctx context.Context, req *DeleteUserReq) (*DeleteUserRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.DeleteUser(ctx, req)
}

func GetUser(ctx context.Context, req *GetUserReq) (*GetUserRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.GetUser(ctx, req)
}

func ListUser(ctx context.Context, req *ListUserReq) (*ListUserRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.ListUser(ctx, req)
}

func CreateRole(ctx context.Context, req *CreateRoleReq) (*CreateRoleRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.CreateRole(ctx, req)
}

func UpdateRole(ctx context.Context, req *UpdateRoleReq) (*UpdateRoleRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.UpdateRole(ctx, req)
}

func DeleteRole(ctx context.Context, req *DeleteRoleReq) (*DeleteRoleRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.DeleteRole(ctx, req)
}

func GetRole(ctx context.Context, req *GetRoleReq) (*GetRoleRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.GetRole(ctx, req)
}

func ListRole(ctx context.Context, req *ListRoleReq) (*ListRoleRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.ListRole(ctx, req)
}

func ListMenu(ctx context.Context, req *ListMenuReq) (*ListMenuRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.ListMenu(ctx, req)
}


