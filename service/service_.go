package app

import (
	"context"

	"github.com/ml444/gkit/errorx"
	log "github.com/ml444/glog"

	"github.com/JIAWea/erpServer/api/erp"
)

type ErpService struct {
	erp.UnsafeErpServer
}

func NewErpService() ErpService {
	return ErpService{}
}

func (s ErpService) CreateRole(ctx context.Context, req *erp.CreateRoleReq) (*erp.CreateRoleRsp, error) {
	var err error
	var rsp erp.CreateRoleRsp

	m := req.Data
	if m == nil {
		log.Error("the req of Model is nil")
		return nil, errorx.New(erp.ErrParamRequired)
	}

	// do something

	err = dbRole.Create(ctx, m)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	rsp.Data = m

	return &rsp, nil
}

func (s ErpService) UpdateRole(ctx context.Context, req *erp.UpdateRoleReq) (*erp.UpdateRoleRsp, error) {
	var err error
	var rsp erp.UpdateRoleRsp

	m := req.Data
	if m == nil || m.Id == 0 {
		log.Error("update request must have an Id")
		return nil, errorx.New(erp.ErrParamRequired)
	}
	err = dbRole.Update(ctx, m, map[string]interface{}{
		"id": m.Id,
	})
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}
	rsp.Data = m

	return &rsp, nil
}

func (s ErpService) DeleteRole(ctx context.Context, req *erp.DeleteRoleReq) (*erp.DeleteRoleRsp, error) {
	var err error
	var rsp erp.DeleteRoleRsp

	err = dbRole.DeleteByWhere(ctx, map[string]interface{}{
		"id": req.Id,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &rsp, nil
}

func (s ErpService) GetRole(ctx context.Context, req *erp.GetRoleReq) (*erp.GetRoleRsp, error) {
	var err error
	var rsp erp.GetRoleRsp

	m, err := dbRole.GetOne(ctx, req.Id)
	if err != nil {
		log.Error(err)
		return nil, errorx.New(erp.ErrNotFoundRole)
	}
	rsp.Data = m

	return &rsp, nil
}

func (s ErpService) ListRole(ctx context.Context, req *erp.ListRoleReq) (*erp.ListRoleRsp, error) {
	var err error
	var rsp erp.ListRoleRsp

	queryOpts := make(map[string]interface{})
	list, paginate, err := dbRole.ListWithListOption(ctx, req.ListOption, queryOpts)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}
	rsp.Paginate = paginate
	rsp.List = list

	return &rsp, nil
}

func (s ErpService) ListMenu(ctx context.Context, req *erp.ListMenuReq) (*erp.ListMenuRsp, error) {
	var err error
	var rsp erp.ListMenuRsp

	queryOpts := make(map[string]interface{})
	// do something...
	list, paginate, err := dbMenu.ListWithListOption(ctx, req.ListOption, queryOpts)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}
	rsp.Paginate = paginate
	rsp.List = list

	return &rsp, nil
}

func (s ErpService) CreateAccount(ctx context.Context, req *erp.CreateAccountReq) (*erp.CreateAccountRsp, error) {
	var err error
	var rsp erp.CreateAccountRsp

	m := req.Data
	if m == nil {
		log.Error("the req of Model is nil")
		return nil, errorx.New(erp.ErrParamRequired)
	}

	// do something

	err = dbAccount.Create(ctx, m)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	rsp.Data = m

	return &rsp, nil
}

func (s ErpService) UpdateAccount(ctx context.Context, req *erp.UpdateAccountReq) (*erp.UpdateAccountRsp, error) {
	var err error
	var rsp erp.UpdateAccountRsp

	m := req.Data
	if m == nil || m.Id == 0 {
		log.Error("update request must have an Id")
		return nil, errorx.New(erp.ErrParamRequired)
	}
	err = dbAccount.Update(ctx, m, map[string]interface{}{
		"id": m.Id,
	})
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}
	rsp.Data = m

	return &rsp, nil
}

func (s ErpService) DeleteAccount(ctx context.Context, req *erp.DeleteAccountReq) (*erp.DeleteAccountRsp, error) {
	var err error
	var rsp erp.DeleteAccountRsp

	err = dbAccount.DeleteByWhere(ctx, map[string]interface{}{
		"id": req.Id,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &rsp, nil
}

func (s ErpService) GetAccount(ctx context.Context, req *erp.GetAccountReq) (*erp.GetAccountRsp, error) {
	var err error
	var rsp erp.GetAccountRsp

	m, err := dbAccount.GetOne(ctx, req.Id)
	if err != nil {
		log.Error(err)
		return nil, errorx.New(erp.ErrNotFoundAccount)
	}
	rsp.Data = m

	return &rsp, nil
}

func (s ErpService) ListAccount(ctx context.Context, req *erp.ListAccountReq) (*erp.ListAccountRsp, error) {
	var err error
	var rsp erp.ListAccountRsp

	queryOpts := make(map[string]interface{})
	// do something...
	list, paginate, err := dbAccount.ListWithListOption(ctx, req.ListOption, queryOpts)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}
	rsp.Paginate = paginate
	rsp.List = list

	return &rsp, nil
}
