package app

import (
	"context"
	"github.com/ml444/gkit/errorx"
	log "github.com/ml444/glog"

	"github.com/JIAWea/erpServer/api/erp"
)

func (s ErpService) CreateAccount(ctx context.Context, req *erp.CreateAccountReq) (*erp.CreateAccountRsp, error) {
	var err error
	var rsp erp.CreateAccountRsp

	m := req.Data
	if m == nil {
		log.Error("the req of Model is nil")
		return nil, errorx.New(erp.ErrParamRequired)
	}

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
	list, paginate, err := dbAccount.ListWithListOption(ctx, req.ListOption, queryOpts)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}

	rsp.AccountStat = make(map[uint64]*erp.ListAccountRsp_AccountStat, len(list))

	for _, v := range list {
		var exp erp.ModelExpense
		err = dbExpense.newScope().
			Select("SUM(pay_money) AS pay_money").
			Eq(dbAccountId, v.Id).
			First(&exp)
		if err != nil {
			log.Errorf("err: %v", err)
			return nil, err
		}

		rsp.AccountStat[v.Id] = &erp.ListAccountRsp_AccountStat{
			TotalExpense: exp.PayMoney,
			//TotalIncome:  totalIncome,
		}
	}

	rsp.Paginate = paginate
	rsp.List = list

	return &rsp, nil
}
