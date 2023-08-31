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

func ListAccountOpt(ctx context.Context, req *ListAccountOptReq) (*ListAccountOptRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.ListAccountOpt(ctx, req)
}

func ListUserAccount(ctx context.Context, req *ListUserAccountReq) (*ListUserAccountRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.ListUserAccount(ctx, req)
}

func UpdateUserAccount(ctx context.Context, req *UpdateUserAccountReq) (*UpdateUserAccountRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.UpdateUserAccount(ctx, req)
}

func CreateExpense(ctx context.Context, req *CreateExpenseReq) (*CreateExpenseRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.CreateExpense(ctx, req)
}

func UpdateExpense(ctx context.Context, req *UpdateExpenseReq) (*UpdateExpenseRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.UpdateExpense(ctx, req)
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

func CreateIncome(ctx context.Context, req *CreateIncomeReq) (*CreateIncomeRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.CreateIncome(ctx, req)
}

func UpdateIncome(ctx context.Context, req *UpdateIncomeReq) (*UpdateIncomeRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.UpdateIncome(ctx, req)
}

func DeleteIncome(ctx context.Context, req *DeleteIncomeReq) (*DeleteIncomeRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.DeleteIncome(ctx, req)
}

func ListIncome(ctx context.Context, req *ListIncomeReq) (*ListIncomeRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.ListIncome(ctx, req)
}

func CreatePlan(ctx context.Context, req *CreatePlanReq) (*CreatePlanRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.CreatePlan(ctx, req)
}

func UpdatePlan(ctx context.Context, req *UpdatePlanReq) (*UpdatePlanRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.UpdatePlan(ctx, req)
}

func DeletePlan(ctx context.Context, req *DeletePlanReq) (*DeletePlanRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.DeletePlan(ctx, req)
}

func ListPlan(ctx context.Context, req *ListPlanReq) (*ListPlanRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.ListPlan(ctx, req)
}

func CreatePlanDetail(ctx context.Context, req *CreatePlanDetailReq) (*CreatePlanDetailRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.CreatePlanDetail(ctx, req)
}

func UpdatePlanDetail(ctx context.Context, req *UpdatePlanDetailReq) (*UpdatePlanDetailRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.UpdatePlanDetail(ctx, req)
}

func DeletePlanDetail(ctx context.Context, req *DeletePlanDetailReq) (*DeletePlanDetailRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.DeletePlanDetail(ctx, req)
}

func ListPlanDetail(ctx context.Context, req *ListPlanDetailReq) (*ListPlanDetailRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.ListPlanDetail(ctx, req)
}


