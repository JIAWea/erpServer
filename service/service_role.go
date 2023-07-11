package app

import (
	"context"
	"github.com/JIAWea/erpServer/api/erp"
	"github.com/ml444/gkit/errorx"
	log "github.com/ml444/glog"
)

func (s ErpService) CreateRole(ctx context.Context, req *erp.CreateRoleReq) (*erp.CreateRoleRsp, error) {
	var err error
	var rsp erp.CreateRoleRsp

	m := req.Data
	if m == nil {
		log.Error("the req of Model is nil")
		return nil, errorx.New(erp.ErrParamRequired)
	}

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
