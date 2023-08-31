package app

import (
	"context"

	"github.com/ml444/gkit/core"

	"github.com/JIAWea/erpServer/pkg/utils"
	"github.com/ml444/gkit/errorx"
	log "github.com/ml444/glog"

	"github.com/JIAWea/erpServer/api/erp"
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
		m.Uuid = utils.GenUUID()
	}

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
		dbCustomer:     m.Customer,
		dbMark:         m.Mark,
		dbTotalMoney:   m.TotalMoney,
		dbTradeMoney:   m.TradeMoney,
		dbBalanceMoney: m.TotalMoney - m.TradeMoney,
		dbComment:      m.Comment,
		dbAttachment:   m.Attachment,
		dbAttName:      m.AttName,
		dbStatus:       m.Status,
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

func (s ErpService) CreatePlanDetail(ctx context.Context, req *erp.CreatePlanDetailReq) (*erp.CreatePlanDetailRsp, error) {
	var err error
	var rsp erp.CreatePlanDetailRsp

	m := req.Data
	if m == nil {
		log.Error("the req of Model is nil")
		return nil, errorx.New(erp.ErrParamRequired)
	}

	err = dbPlanDetail.Create(ctx, m)
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

func (s ErpService) ListPlanDetail(ctx context.Context, req *erp.ListPlanDetailReq) (*erp.ListPlanDetailRsp, error) {
	var err error
	var rsp erp.ListPlanDetailRsp

	queryOpts := make(map[string]interface{})
	// do something...
	list, paginate, err := dbPlanDetail.ListWithListOption(ctx, req.ListOption, queryOpts)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}
	rsp.Paginate = paginate
	rsp.List = list

	return &rsp, nil
}
