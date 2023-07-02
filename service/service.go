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

func (s ErpService) CreateErp(ctx context.Context, req *erp.CreateErpReq) (*erp.CreateErpRsp, error) {
	var err error
	var rsp erp.CreateErpRsp

	m := req.Data
	if m == nil {
		log.Error("the req of Model is nil")
		return nil, errorx.New(erp.ErrParamRequired)
	}

	// do something

	err = dbErp.Create(ctx, m)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	rsp.Data = m

	return &rsp, nil
}

func (s ErpService) UpdateErp(ctx context.Context, req *erp.UpdateErpReq) (*erp.UpdateErpRsp, error) {
	var err error
	var rsp erp.UpdateErpRsp

	m := req.Data
	if m == nil || m.Id == 0 {
		log.Error("update request must have an Id")
		return nil, errorx.New(erp.ErrParamRequired)
	}
	err = dbErp.Update(ctx, m, map[string]interface{}{
		"id": m.Id,
	})
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}
	rsp.Data = m

	return &rsp, nil
}

func (s ErpService) DeleteErp(ctx context.Context, req *erp.DeleteErpReq) (*erp.DeleteErpRsp, error) {
	var err error
	var rsp erp.DeleteErpRsp

	err = dbErp.DeleteByWhere(ctx, map[string]interface{}{
		"id": req.Id,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &rsp, nil
}

func (s ErpService) GetErp(ctx context.Context, req *erp.GetErpReq) (*erp.GetErpRsp, error) {
	var err error
	var rsp erp.GetErpRsp

	m, err := dbErp.GetOne(ctx, req.Id)
	if err != nil {
		log.Error(err)
		return nil, errorx.New(erp.ErrNotFoundErp)
	}
	rsp.Data = m

	return &rsp, nil
}

func (s ErpService) ListErp(ctx context.Context, req *erp.ListErpReq) (*erp.ListErpRsp, error) {
	var err error
	var rsp erp.ListErpRsp

	queryOpts := make(map[string]interface{})
	// do something...
	list, paginate, err := dbErp.ListWithListOption(ctx, req.ListOption, queryOpts)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}
	rsp.Paginate = paginate
	rsp.List = list

	return &rsp, nil
}

func (s ErpService) DDListErp(ctx context.Context, req *erp.ListErpReq) (*erp.ListErpRsp, error) {
	var err error
	var rsp erp.ListErpRsp

	queryOpts := make(map[string]interface{})
	// do something...
	list, paginate, err := dbErp.ListWithListOption(ctx, req.ListOption, queryOpts)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}
	rsp.Paginate = paginate
	rsp.List = list

	return &rsp, nil
}
