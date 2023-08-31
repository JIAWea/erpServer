package app

import (
	"context"
	"github.com/JIAWea/erpServer/api/erp"
	"github.com/ml444/gkit/core"
	"github.com/ml444/gkit/errorx"
)

// CheckPerm 权限校验
func CheckPerm(ctx context.Context, req interface{}) (context.Context, interface{}, error) {
	val := ctx.Value(core.HeadersKey)
	header, ok := val.(core.Header)
	if !ok {
		return ctx, nil, errorx.New(erp.ErrInternal)
	}

	path := header[core.HttpHeaderPath]

	var menu erp.ModelMenu
	err := dbMenu.newScope().Eq(dbPath, path).SetNotFoundErr(erp.ErrNotFoundMenu).First(&menu)
	if err != nil {
		return ctx, req, err
	}

	userId, err := core.MustUserId(ctx)
	if err != nil {
		return ctx, req, errorx.New(erp.ErrAuthContext)
	}
	var userRoleList []*erp.ModelUserRole
	err = dbUserRole.newScope().Eq(dbUserId, userId).Find(&userRoleList)
	if err != nil {
		return ctx, req, err
	}

	var roleIdList []uint64
	for _, v := range userRoleList {
		roleIdList = append(roleIdList, v.RoleId)
	}
	if len(roleIdList) == 0 {
		return ctx, req, errorx.New(erp.Err403)
	}

	var total int64
	err = dbRoleMenu.newScope().In(dbRoleId, roleIdList).Eq(dbMenuId, menu.Id).Count(&total)
	if err != nil {
		return ctx, req, err
	}
	if total == 0 {
		return ctx, req, errorx.New(erp.Err403)
	}

	return ctx, req, nil
}
