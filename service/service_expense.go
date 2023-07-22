package app

import (
	"context"
	"github.com/ml444/gkit/core"

	"github.com/JIAWea/erpServer/pkg/utils"
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

	m.UserId = core.GetUserId(ctx)
	if m.Uuid == "" {
		m.Uuid = utils.GenUUID()
	}

	err = dbExpense.Create(ctx, m)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &rsp, nil
}

func (s ErpService) UpdateExpense(ctx context.Context, req *erp.UpdateExpenseReq) (*erp.UpdateExpenseRsp, error) {
	var err error
	var rsp erp.UpdateExpenseRsp

	m := req.Data
	if m == nil || m.Id == 0 {
		log.Error("update request must have an Id")
		return nil, errorx.New(erp.ErrParamRequired)
	}
	err = dbExpense.Update(ctx, m.Id, map[string]interface{}{
		dbPayAt:      m.PayAt,
		dbCategory:   m.Category,
		dbMark:       m.Mark,
		dbPayMoney:   m.PayMoney,
		dbAccountId:  m.AccountId,
		dbTicket:     m.Ticket,
		dbHandleBy:   m.HandleBy,
		dbAttachment: m.Attachment,
		dbAttName:    m.AttName,
	})
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}
	rsp.Data = m

	return &rsp, nil
}

func (s ErpService) DeleteExpense(ctx context.Context, req *erp.DeleteExpenseReq) (*erp.DeleteExpenseRsp, error) {
	var err error
	var rsp erp.DeleteExpenseRsp

	err = dbExpense.DeleteByIdList(ctx, req.IdList)
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

	var accIdList []uint64
	for _, v := range list {
		if v.AccountId != 0 {
			accIdList = append(accIdList, v.AccountId)
		}
	}
	if len(accIdList) > 0 {
		var accList []*erp.ModelAccount
		err = dbAccount.newScope().In(dbId, accIdList).Find(&accList)
		if err != nil {
			return nil, err
		}
		accMap := make(map[uint64]*erp.ModelAccount)
		for _, v := range accList {
			accMap[v.Id] = v
		}
		rsp.AccountMap = accMap
	}

	return &rsp, nil
}
