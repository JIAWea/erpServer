package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/ml444/gkit/core"
	"github.com/ml444/gkit/errorx"
	"github.com/ml444/gkit/listoption"
	log "github.com/ml444/glog"
	"strconv"
	"strings"

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

	uid := core.GetUserId(ctx)

	var idList []uint64
	if !dbUser.IsCreator(ctx, uid) {
		idList, err = dbUserAccount.GetIdListByUserId(ctx, uid)
		if err != nil {
			log.Errorf("err: %v", err)
			return nil, err
		}
		if len(idList) == 0 {
			return &rsp, nil
		}
	}

	list, paginate, err := dbAccount.ListWithListOption(ctx, req.ListOption, idList)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}

	statStartAt := uint32(0)
	statEndAt := uint32(0)
	if req.ListOption != nil {
		for _, option := range req.ListOption.Options {
			if option.Type == int32(erp.ListAccountReq_ListOptStatTimeRange) {
				valList := strings.Split(option.Value, ",")
				if len(valList) != 2 {
					return nil, errorx.New(erp.ErrInvalidParam)
				}
				s1, _ := strconv.ParseUint(valList[0], 10, 64)
				s2, _ := strconv.ParseUint(valList[1], 10, 64)
				statStartAt, statEndAt = uint32(s1), uint32(s2)
			}
		}
	}

	rsp.AccountStat = make(map[uint64]*erp.ListAccountRsp_AccountStat, len(list))

	for _, v := range list {
		var exp erp.ModelExpense
		expScope := dbExpense.newScope().
			Select("SUM(pay_money) AS pay_money").
			SetNotFoundErr(erp.ErrNotFoundExpense).
			Eq(dbAccountId, v.Id)
		if statStartAt > 0 && statEndAt > 0 {
			expScope.Where(fmt.Sprintf("%s >= ? AND %s <= ?", dbPayAt, dbPayAt), statStartAt, statEndAt)
		}
		err = expScope.First(&exp)
		if err != nil {
			if !errors.Is(err, errorx.New(erp.ErrNotFoundExpense)) {
				log.Errorf("err: %v", err)
				return nil, err
			}
		}

		var inc erp.ModelIncome
		incScope := dbIncome.newScope().
			Select("SUM(income_money) AS income_money").
			SetNotFoundErr(erp.ErrNotFoundIncome).
			Eq(dbAccountId, v.Id)
		if statStartAt > 0 && statEndAt > 0 {
			expScope.Where(fmt.Sprintf("%s >= ? AND %s <= ?", dbIncomeAt, dbIncomeAt), statStartAt, statEndAt)
		}
		err = incScope.First(&inc)
		if err != nil {
			if !errors.Is(err, errorx.New(erp.ErrNotFoundIncome)) {
				log.Errorf("err: %v", err)
				return nil, err
			}
		}

		var balance uint32
		if exp.PayMoney <= v.InitialMoney {
			balance = v.InitialMoney + inc.IncomeMoney - exp.PayMoney
		}

		rsp.AccountStat[v.Id] = &erp.ListAccountRsp_AccountStat{
			TotalExpense: exp.PayMoney,
			TotalIncome:  inc.IncomeMoney,
			Balance:      balance,
		}
	}

	rsp.Paginate = paginate
	rsp.List = list

	return &rsp, nil
}

func (s ErpService) ListAccountOpt(ctx context.Context, req *erp.ListAccountOptReq) (*erp.ListAccountOptRsp, error) {
	var err error
	var rsp erp.ListAccountOptRsp

	list, _, err := dbAccount.ListWithListOption(ctx, &listoption.ListOption{SkipCount: true}, nil)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}

	rsp.List = list

	return &rsp, nil
}
