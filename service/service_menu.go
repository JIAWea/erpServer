package app

import (
	"context"
	"github.com/ml444/gkit/errorx"

	log "github.com/ml444/glog"

	"github.com/JIAWea/erpServer/api/erp"
)

func (s ErpService) ListMenu(ctx context.Context, req *erp.ListMenuReq) (*erp.ListMenuRsp, error) {
	var err error
	var rsp erp.ListMenuRsp

	queryOpts := map[string]interface{}{
		dbStatus: uint32(erp.ModelMenu_StatusEnable),
	}
	list, paginate, err := dbMenu.ListWithListOption(ctx, req.ListOption, queryOpts)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}
	rsp.Paginate = paginate
	rsp.List = list

	return &rsp, nil
}

func (s ErpService) ListMenuTree(ctx context.Context, req *erp.ListMenuTreeReq) (*erp.ListMenuTreeRsp, error) {
	var err error
	var rsp erp.ListMenuTreeRsp

	list, err := dbMenu.GetMenuTree(ctx, req.ListOption)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}
	rsp.List = list

	return &rsp, nil
}

func (s ErpService) GetRoleMenuIdList(ctx context.Context, req *erp.GetRoleMenuIdListReq) (*erp.GetRoleMenuIdListRsp, error) {
	var rsp erp.GetRoleMenuIdListRsp

	list, err := dbRoleMenu.GetRoleMenuIdList(ctx, req.Type, req.RoleId)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}
	rsp.MenuIdList = list

	return &rsp, nil
}

func (s ErpService) UpdateRoleMenu(ctx context.Context, req *erp.UpdateRoleMenuReq) (*erp.UpdateRoleMenuRsp, error) {
	var rsp erp.UpdateRoleMenuRsp
	var err error

	if req.Id == 0 {
		return nil, errorx.New(erp.ErrParamRequired)
	}
	err = dbRoleMenu.DeleteByWhere(ctx,
		map[string]interface{}{
			"role_id": req.Id,
		})
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}

	existMap := make(map[uint64]struct{})
	var list []*erp.ModelRoleMenu
	for _, v := range req.MenuIdList {
		if _, ok := existMap[v]; !ok {
			list = append(list, &erp.ModelRoleMenu{
				RoleId: req.Id,
				MenuId: v,
			})
			existMap[v] = struct{}{}
		}
	}
	err = dbRoleMenu.CreateInBatches(ctx, list)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}

	return &rsp, nil
}
