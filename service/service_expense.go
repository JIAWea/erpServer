package app

import (
	"context"

	"github.com/ml444/gkit/errorx"
	log "github.com/ml444/glog"

	"github.com/JIAWea/erpServer/api/erp"
)

func (s ErpService) ImportExpense(ctx context.Context, req *erp.ImportExpenseReq) (*erp.ImportExpenseRsp, error) {
	var rsp erp.ImportExpenseRsp

	return &rsp, nil
}

func (s ErpService) CreateExpense(ctx context.Context, req *erp.CreateExpenseReq) (*erp.CreateExpenseRsp, error) {
	var err error
	var rsp erp.CreateExpenseRsp

	m := req.Data
	if m == nil {
		log.Error("the req of Model is nil")
		return nil, errorx.New(erp.ErrParamRequired)
	}

	// do something

	err = dbExpense.Create(ctx, m)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &rsp, nil
}

func (s ErpService) DeleteExpense(ctx context.Context, req *erp.DeleteExpenseReq) (*erp.DeleteExpenseRsp, error) {
	var err error
	var rsp erp.DeleteExpenseRsp

	err = dbExpense.DeleteByWhere(ctx, map[string]interface{}{
		"id": req.Id,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &rsp, nil
}

func (s ErpService) ListExpense(ctx context.Context, req *erp.ListExpenseReq) (*erp.ListExpenseRsp, error) {
	var err error
	var rsp erp.ListExpenseRsp

	queryOpts := make(map[string]interface{})
	list, paginate, err := dbExpense.ListWithListOption(ctx, req.ListOption, queryOpts)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}
	rsp.Paginate = paginate
	rsp.List = list

	return &rsp, nil
}
