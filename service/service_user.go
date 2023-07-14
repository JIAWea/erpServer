package app

import (
	"context"
	"regexp"
	"strings"

	"github.com/JIAWea/erpServer/api/erp"
	"github.com/JIAWea/erpServer/pkg/utils"
	"github.com/ml444/gkit/auth"
	"github.com/ml444/gkit/core"
	"github.com/ml444/gkit/errorx"
	log "github.com/ml444/glog"
)

func (s ErpService) UserLogin(ctx context.Context, req *erp.UserLoginReq) (*erp.UserLoginRsp, error) {
	var err error
	var rsp erp.UserLoginRsp

	if ok, _ := regexp.MatchString(utils.DefaultPasswordRegex, req.Password); !ok {
		return nil, errorx.New(erp.ErrPassword)
	}

	// 密码判断
	user, err := dbUser.GetOneByName(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if err = utils.ComparePasswd(user.Password, req.Password); err != nil {
		return nil, errorx.New(erp.ErrPassword)
	}

	tk, err := auth.GenerateJWT(user.Id, 0, "manager-web")
	if err != nil {
		return nil, errorx.New(erp.ErrInternal)
	}

	rsp.Token = tk

	return &rsp, nil
}

func (s ErpService) UserLogout(ctx context.Context, req *erp.UserLogoutReq) (*erp.UserLogoutRsp, error) {
	var rsp erp.UserLogoutRsp
	auth.DelCacheAuthDataBySign(ctx, strings.TrimPrefix(req.Token, "Bearer "))
	return &rsp, nil
}

func (s ErpService) GetUserInfo(ctx context.Context, req *erp.GetUserInfoReq) (*erp.GetUserInfoRsp, error) {
	var err error
	var rsp erp.GetUserInfoRsp

	user, err := dbUser.GetOne(ctx, core.GetUserId(ctx))
	if err != nil {
		return nil, err
	}

	rsp.Data = user

	return &rsp, nil
}

func (s ErpService) CreateUser(ctx context.Context, req *erp.CreateUserReq) (*erp.CreateUserRsp, error) {
	var err error
	var rsp erp.CreateUserRsp

	m := req.Data
	if m == nil {
		log.Error("the req of Model is nil")
		return nil, errorx.New(erp.ErrParamRequired)
	}

	if ok, _ := regexp.MatchString(utils.DefaultPasswordRegex, m.Password); !ok {
		return nil, errorx.New(erp.ErrPasswordFormatInvalid)
	}

	m.Password = utils.GenPasswd(m.Password)
	err = dbUser.Create(ctx, m)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	m.Password = ""
	rsp.Data = m

	return &rsp, nil
}

func (s ErpService) UpdateUser(ctx context.Context, req *erp.UpdateUserReq) (*erp.UpdateUserRsp, error) {
	var err error
	var rsp erp.UpdateUserRsp

	m := req.Data
	if m == nil || m.Id == 0 {
		log.Error("update request must have an Id")
		return nil, errorx.New(erp.ErrParamRequired)
	}
	err = dbUser.Update(ctx,
		map[string]interface{}{
			dbName:     m.Name,
			dbNickName: m.NickName,
			dbStatus:   m.Status,
		},
		map[string]interface{}{
			"id": m.Id,
		})
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}
	rsp.Data = m

	return &rsp, nil
}

func (s ErpService) UpdateUserRole(ctx context.Context, req *erp.UpdateUserRoleReq) (*erp.UpdateUserRoleRsp, error) {
	var err error
	var rsp erp.UpdateUserRoleRsp

	if req.Id == 0 {
		return nil, errorx.New(erp.ErrParamRequired)
	}
	err = dbUserRole.DeleteByWhere(ctx,
		map[string]interface{}{
			"user_id": req.Id,
		})
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}

	existMap := make(map[uint64]struct{})
	var list []*erp.ModelUserRole
	for _, v := range req.RoleIdList {
		if _, ok := existMap[v]; !ok {
			list = append(list, &erp.ModelUserRole{
				UserId: req.Id,
				RoleId: v,
			})
			existMap[v] = struct{}{}
		}
	}
	err = dbUserRole.CreateInBatches(ctx, list)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}

	return &rsp, nil
}

func (s ErpService) UpdatePassword(ctx context.Context, req *erp.UpdatePasswordReq) (*erp.UpdatePasswordRsp, error) {
	var err error
	var rsp erp.UpdatePasswordRsp

	m := req.Data
	m.Id = core.GetUserId(ctx)
	if m == nil || m.Id == 0 {
		log.Error("update request must have an Id")
		return nil, errorx.New(erp.ErrParamRequired)
	}

	// 密码合法性校验
	if m.NewPassword != m.NewPasswordAgain {
		return nil, errorx.New(erp.ErrPasswordFormatInvalid)
	}
	if ok, _ := regexp.MatchString(utils.DefaultPasswordRegex, m.NewPassword); !ok {
		return nil, errorx.New(erp.ErrPasswordFormatInvalid)
	}

	if err = dbUser.CheckPassword(ctx, m.Id, m.OldPassword); err != nil {
		return nil, errorx.New(erp.ErrPassword)
	}

	err = dbUser.Update(ctx,
		map[string]interface{}{
			dbPassword: utils.GenPasswd(m.NewPassword),
		},
		map[string]interface{}{
			"id": m.Id,
		})
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}

	return &rsp, nil
}

func (s ErpService) DeleteUser(ctx context.Context, req *erp.DeleteUserReq) (*erp.DeleteUserRsp, error) {
	var err error
	var rsp erp.DeleteUserRsp

	err = dbUser.DeleteByIdList(ctx, req.IdList)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &rsp, nil
}

func (s ErpService) GetUser(ctx context.Context, req *erp.GetUserReq) (*erp.GetUserRsp, error) {
	var err error
	var rsp erp.GetUserRsp

	m, err := dbUser.GetOne(ctx, req.Id)
	if err != nil {
		log.Error(err)
		return nil, errorx.New(erp.ErrNotFoundUser)
	}
	rsp.Data = m

	return &rsp, nil
}

func (s ErpService) ListUser(ctx context.Context, req *erp.ListUserReq) (*erp.ListUserRsp, error) {
	var err error
	var rsp erp.ListUserRsp

	queryOpts := make(map[string]interface{})
	list, paginate, err := dbUser.ListWithListOption(ctx, req.ListOption, queryOpts)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}
	rsp.Paginate = paginate
	rsp.List = list

	return &rsp, nil
}
