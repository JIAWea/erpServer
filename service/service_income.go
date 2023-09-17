package app

import (
	"context"
	"fmt"

	"github.com/ml444/gkit/core"

	"github.com/JIAWea/erpServer/pkg/utils"
	"github.com/ml444/gkit/errorx"
	log "github.com/ml444/glog"

	"github.com/JIAWea/erpServer/api/erp"
)

func (s ErpService) CreateIncome(ctx context.Context, req *erp.CreateIncomeReq) (*erp.CreateIncomeRsp, error) {
	var err error
	var rsp erp.CreateIncomeRsp

	m := req.Data
	if m == nil {
		log.Error("the req of Model is nil")
		return nil, errorx.New(erp.ErrParamRequired)
	}

	m.UserId = core.GetUserId(ctx)
	if m.Uuid == "" {
		m.Uuid = fmt.Sprintf("SR-%s", utils.RandomUUID(6))
	}

	err = dbIncome.Create(ctx, m)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &rsp, nil
}

func (s ErpService) UpdateIncome(ctx context.Context, req *erp.UpdateIncomeReq) (*erp.UpdateIncomeRsp, error) {
	var err error
	var rsp erp.UpdateIncomeRsp

	m := req.Data
	if m == nil || m.Id == 0 {
		log.Error("update request must have an Id")
		return nil, errorx.New(erp.ErrParamRequired)
	}
	err = dbIncome.Update(ctx, m.Id, map[string]interface{}{
		dbIncomeAt:    m.IncomeAt,
		dbCategory:    m.Category,
		dbMark:        m.Mark,
		dbIncomeMoney: m.IncomeMoney,
		dbAccountId:   m.AccountId,
		dbFrom:        m.From,
		dbHandleBy:    m.HandleBy,
		dbAttachment:  m.Attachment,
		dbAttName:     m.AttName,
	})
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}
	rsp.Data = m

	return &rsp, nil
}

func (s ErpService) DeleteIncome(ctx context.Context, req *erp.DeleteIncomeReq) (*erp.DeleteIncomeRsp, error) {
	var err error
	var rsp erp.DeleteIncomeRsp

	err = dbIncome.DeleteByIdList(ctx, req.IdList)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &rsp, nil
}

func (s ErpService) ListIncome(ctx context.Context, req *erp.ListIncomeReq) (*erp.ListIncomeRsp, error) {
	var err error
	var rsp erp.ListIncomeRsp
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

	list, paginate, err := dbIncome.ListWithListOption(ctx, req.ListOption, accViewIdList)
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
