package app

import (
	"context"
	"fmt"
	"github.com/JIAWea/erpServer/api/erp"
	"github.com/JIAWea/erpServer/pkg/utils"
	"github.com/ml444/gkit/core"
	"github.com/ml444/gkit/errorx"
	log "github.com/ml444/glog"
)

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
		m.Uuid = fmt.Sprintf("ZC-%s", utils.RandomUUID(6))
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
	var accViewIdList []uint64

	uid := core.GetUserId(ctx)
	if !dbUser.IsCreator(ctx, uid) {
		accViewIdList, err = dbUserAccount.GetIdListByUserId(ctx, uid)
		if err != nil {
			log.Errorf("err: %v", err)
			return nil, err
		}
		if len(accViewIdList) == 0 {
			return &rsp, nil
		}
	}

	list, paginate, err := dbExpense.ListWithListOption(ctx, req.ListOption, accViewIdList)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}
	rsp.Paginate = paginate
	rsp.List = list

	var (
		accIdList     []uint64
		accIdExistMap = make(map[uint64]struct{})
	)
	for _, v := range list {
		if _, ok := accIdExistMap[v.AccountId]; !ok {
			accIdList = append(accIdList, v.AccountId)
			accIdExistMap[v.AccountId] = struct{}{}
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
