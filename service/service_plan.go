package app

import (
	"context"
	"fmt"

	"github.com/JIAWea/erpServer/internal/db"
	"github.com/JIAWea/erpServer/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/JIAWea/erpServer/api/erp"
	"github.com/ml444/gkit/core"
	"github.com/ml444/gkit/errorx"
	log "github.com/ml444/glog"
)

func (s ErpService) CreatePlan(ctx context.Context, req *erp.CreatePlanReq) (*erp.CreatePlanRsp, error) {
	var err error
	var rsp erp.CreatePlanRsp

	m := req.Data
	if m == nil {
		log.Error("the req of Model is nil")
		return nil, errorx.New(erp.ErrParamRequired)
	}

	m.UserId = core.GetUserId(ctx)
	if m.Uuid == "" {
		var prefix string
		switch m.Type {
		case uint32(erp.ModelPlan_TypePay):
			prefix = "YF"
		case uint32(erp.ModelPlan_TypeReceive):
			prefix = "YS"
		default:
		}
		m.Uuid = fmt.Sprintf("%s-%s", prefix, utils.RandomUUID(6))
	}

	m.TradeMoney = 0
	m.BalanceMoney = 0
	m.Status = uint32(erp.ModelPlan_StatusWaitConfirm)
	err = dbPlan.Create(ctx, m)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &rsp, nil
}

func (s ErpService) UpdatePlan(ctx context.Context, req *erp.UpdatePlanReq) (*erp.UpdatePlanRsp, error) {
	var err error
	var rsp erp.UpdatePlanRsp

	m := req.Data
	if m == nil || m.Id == 0 {
		log.Error("update request must have an Id")
		return nil, errorx.New(erp.ErrParamRequired)
	}
	err = dbPlan.Update(ctx, m.Id, map[string]interface{}{
		dbCustomer:   m.Customer,
		dbMark:       m.Mark,
		dbTotalMoney: m.TotalMoney,
		// dbTradeMoney:   m.TradeMoney,
		// dbBalanceMoney: m.TotalMoney - m.TradeMoney,
		dbComment:    m.Comment,
		dbAttachment: m.Attachment,
		dbAttName:    m.AttName,
		dbStatus:     m.Status,
	})
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}
	rsp.Data = m

	return &rsp, nil
}

func (s ErpService) DeletePlan(ctx context.Context, req *erp.DeletePlanReq) (*erp.DeletePlanRsp, error) {
	var err error
	var rsp erp.DeletePlanRsp

	err = dbPlan.DeleteByIdList(ctx, req.IdList)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &rsp, nil
}

// ListPlan      应收应付列表
// @Summary      应收应付列表
// @Description  应收应付列表
// @Tags         应收应付
// @Accept       json
// @Produce      json
// @Param        request body   erp.ListPlanReq true "options:[{type:1,value:'1'(1应付，2应收))},{type:2,value:'客户名称'},{type:3,value:'摘要'},{type:4,value:'1694957121,1694957127'},{type:5,value:'金额'},{type:6,value:'备注'},{type:7,value:'状态'}]"
// @Success      200  {object}  erp.ListPlanRsp
// @Router       /erp/ListPlan [post]
func (s ErpService) ListPlan(ctx context.Context, req *erp.ListPlanReq) (*erp.ListPlanRsp, error) {
	var err error
	var rsp erp.ListPlanRsp

	list, paginate, err := dbPlan.ListWithListOption(ctx, req.ListOption)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}
	rsp.Paginate = paginate
	rsp.List = list

	return &rsp, nil
}

// CreatePlanDetail 创建明细
// @Summary      创建明细
// @Description  创建明细
// @Tags         应收应付
// @Accept       json
// @Produce      json
// @Param        request body   erp.CreatePlanDetailReq true "JSON入参"
// @Success      200  {object}  erp.CreatePlanDetailRsp
// @Router       /erp/CreatePlanDetail [post]
func (s ErpService) CreatePlanDetail(ctx context.Context, req *erp.CreatePlanDetailReq) (*erp.CreatePlanDetailRsp, error) {
	var err error
	var rsp erp.CreatePlanDetailRsp

	m := req.Data
	if m == nil {
		log.Error("the req of Model is nil")
		return nil, errorx.New(erp.ErrParamRequired)
	}
	if req.Data.TradeMoney == 0 {
		return nil, errorx.New(erp.ErrTradeMoneyRequired)
	}
	m.UserId = core.GetUserId(ctx)

	plan, err := dbPlan.GetOne(ctx, req.Data.PlanId)
	if err != nil {
		return nil, errorx.New(erp.ErrNotFoundPlan)
	}
	if plan.BalanceMoney < req.Data.TradeMoney {
		return nil, errorx.New(erp.ErrMoneyBalance)
	}

	switch plan.Type {
	case uint32(erp.ModelPlan_TypePay):
		m.Uuid = fmt.Sprintf("YF-%s", utils.RandomUUID(6))
	case uint32(erp.ModelPlan_TypeReceive):
		m.Uuid = fmt.Sprintf("YS-%s", utils.RandomUUID(6))
	default:
	}

	err = db.Db().Transaction(func(tx *gorm.DB) error {
		// 排它锁
		var p *erp.ModelPlan
		err = tx.WithContext(ctx).Clauses(clause.Locking{Strength: "UPDATE"}).First(&p, req.Data.PlanId).Error
		if err != nil {
			return err
		}
		if p.BalanceMoney < req.Data.TradeMoney {
			return errorx.New(erp.ErrMoneyBalance)
		}

		err = tx.WithContext(ctx).Create(m).Error
		if err != nil {
			return err
		}

		// plan 更新余额
		err = tx.WithContext(ctx).Model(&erp.ModelPlan{}).
			Where(dbId, req.Data.PlanId).
			Updates(map[string]interface{}{
				dbTradeMoney:   gorm.Expr("trade_money + ?", req.Data.TradeMoney),
				dbBalanceMoney: gorm.Expr("balance_money - ?", req.Data.TradeMoney),
			}).Error
		if err != nil {
			return err
		}

		switch p.Type {
		case uint32(erp.ModelPlan_TypePay):
			err = tx.Model(&erp.ModelExpense{}).Create(&erp.ModelExpense{
				PayAt:     m.PlanAt,
				Uuid:      m.Uuid,
				Category:  0, // TODO 分类
				Mark:      m.Mark,
				PayMoney:  m.TradeMoney,
				AccountId: m.AccountId,
				Ticket:    "",
				HandleBy:  m.HandleBy,
				UserId:    m.UserId,
			}).Error
			if err != nil {
				return err
			}
		case uint32(erp.ModelPlan_TypeReceive):
			err = tx.Model(&erp.ModelIncome{}).Create(&erp.ModelIncome{
				IncomeAt:    m.PlanAt,
				Uuid:        m.Uuid,
				Category:    0, // TODO 分类
				Mark:        m.Mark,
				IncomeMoney: m.TradeMoney,
				AccountId:   m.AccountId,
				HandleBy:    m.HandleBy,
				UserId:      m.UserId,
				From:        "",
			}).Error
			if err != nil {
				return err
			}
		default:
			return errorx.New(erp.ErrModelPlanType)
		}

		return nil
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &rsp, nil
}

func (s ErpService) UpdatePlanDetail(ctx context.Context, req *erp.UpdatePlanDetailReq) (*erp.UpdatePlanDetailRsp, error) {
	var err error
	var rsp erp.UpdatePlanDetailRsp

	m := req.Data
	if m == nil || m.Id == 0 {
		log.Error("update request must have an Id")
		return nil, errorx.New(erp.ErrParamRequired)
	}
	err = dbPlanDetail.Update(ctx, m, map[string]interface{}{
		"id": m.Id,
	})
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}
	rsp.Data = m

	return &rsp, nil
}

func (s ErpService) DeletePlanDetail(ctx context.Context, req *erp.DeletePlanDetailReq) (*erp.DeletePlanDetailRsp, error) {
	var err error
	var rsp erp.DeletePlanDetailRsp

	err = dbPlanDetail.DeleteByIdList(ctx, req.IdList)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &rsp, nil
}

// ListPlanDetail 明细列表
// @Summary      明细列表
// @Description  明细列表
// @Tags         应收应付
// @Accept       json
// @Produce      json
// @Param        request body   erp.ListPlanDetailReq true "options: [{type:1,value:'plan_id')}]"
// @Success      200  {object}  erp.ListPlanDetailRsp
// @Router       /erp/ListPlanDetail [post]
func (s ErpService) ListPlanDetail(ctx context.Context, req *erp.ListPlanDetailReq) (*erp.ListPlanDetailRsp, error) {
	var err error
	var rsp erp.ListPlanDetailRsp
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

	list, paginate, err := dbPlanDetail.ListWithListOption(ctx, req.ListOption, accViewIdList)
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
