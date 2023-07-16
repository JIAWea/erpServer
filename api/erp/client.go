package erp

import (
	"context"
)

const ClientName = "erp"



func UserLogin(ctx context.Context, req *UserLoginReq) (*UserLoginRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.UserLogin(ctx, req)
}

func UserLogout(ctx context.Context, req *UserLogoutReq) (*UserLogoutRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.UserLogout(ctx, req)
}

func GetUserInfo(ctx context.Context, req *GetUserInfoReq) (*GetUserInfoRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.GetUserInfo(ctx, req)
}

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

func UpdateUserRole(ctx context.Context, req *UpdateUserRoleReq) (*UpdateUserRoleRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.UpdateUserRole(ctx, req)
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

func ListMenuTree(ctx context.Context, req *ListMenuTreeReq) (*ListMenuTreeRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.ListMenuTree(ctx, req)
}

func GetRoleMenuIdList(ctx context.Context, req *GetRoleMenuIdListReq) (*GetRoleMenuIdListRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.GetRoleMenuIdList(ctx, req)
}

func UpdateRoleMenu(ctx context.Context, req *UpdateRoleMenuReq) (*UpdateRoleMenuRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.UpdateRoleMenu(ctx, req)
}

func CreateAccount(ctx context.Context, req *CreateAccountReq) (*CreateAccountRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.CreateAccount(ctx, req)
}

func UpdateAccount(ctx context.Context, req *UpdateAccountReq) (*UpdateAccountRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.UpdateAccount(ctx, req)
}

func DeleteAccount(ctx context.Context, req *DeleteAccountReq) (*DeleteAccountRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.DeleteAccount(ctx, req)
}

func GetAccount(ctx context.Context, req *GetAccountReq) (*GetAccountRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.GetAccount(ctx, req)
}

func ListAccount(ctx context.Context, req *ListAccountReq) (*ListAccountRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.ListAccount(ctx, req)
}

func ImportExpense(ctx context.Context, req *ImportExpenseReq) (*ImportExpenseRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.ImportExpense(ctx, req)
}

func CreateExpense(ctx context.Context, req *CreateExpenseReq) (*CreateExpenseRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.CreateExpense(ctx, req)
}

func DeleteExpense(ctx context.Context, req *DeleteExpenseReq) (*DeleteExpenseRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.DeleteExpense(ctx, req)
}

func ListExpense(ctx context.Context, req *ListExpenseReq) (*ListExpenseRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.ListExpense(ctx, req)
}


