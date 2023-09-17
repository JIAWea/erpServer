package app

import (
	"context"
	"fmt"

	"github.com/JIAWea/erpServer/api/erp"
	"github.com/JIAWea/erpServer/internal/db"
	"github.com/ml444/gkit/dbx"
	"github.com/ml444/gkit/listoption"
	log "github.com/ml444/glog"
	"gorm.io/gorm"
)

var dbPlan = NewTPlan(db.Db())

type TPlan struct {
	db    *gorm.DB
	model *erp.ModelPlan
}

func NewTPlan(db *gorm.DB) *TPlan {
	return &TPlan{
		db:    db,
		model: &erp.ModelPlan{},
	}
}

func (d *TPlan) newScope() *dbx.Scope {
	if d.db == nil {
		d.db = db.Db()
	}
	return dbx.NewScope(d.db, &erp.ModelPlan{})
}

func (d *TPlan) Create(ctx context.Context, m *erp.ModelPlan) error {
	return d.newScope().Create(ctx, &m)
}

func (d *TPlan) Update(ctx context.Context, id uint64, data map[string]interface{}) error {
	return d.newScope().Where(dbId, id).Update(data)
}

func (d *TPlan) DeleteById(ctx context.Context, pk uint64) error {
	return d.newScope().Delete(&erp.ModelPlan{}, pk)
}

func (d *TPlan) DeleteByIdList(ctx context.Context, idList []uint64) error {
	if len(idList) == 0 {
		return nil
	}
	return d.newScope().In(dbId, idList).Delete(&erp.ModelPlan{})
}

func (d *TPlan) DeleteByWhere(ctx context.Context, whereMap map[string]interface{}) error {
	return d.newScope().Delete(&erp.ModelPlan{}, whereMap)
}

func (d *TPlan) GetOne(ctx context.Context, pk uint64) (*erp.ModelPlan, error) {
	var m erp.ModelPlan
	err := d.newScope().SetNotFoundErr(erp.ErrNotFoundIncome).First(&m, pk)
	return &m, err
}

func (d *TPlan) ListWithListOption(ctx context.Context, listOption *listoption.ListOption) ([]*erp.ModelPlan, *listoption.Paginate, error) {
	var err error
	scope := d.newScope().Order("created_at DESC")

	if listOption != nil {
		err = listoption.NewProcessor(listOption).
			AddUint32(erp.ListPlanReq_ListOptType, func(val uint32) error {
				scope.Eq(dbType, val)
				return nil
			}).
			AddString(erp.ListPlanReq_ListOptCustomer, func(val string) error {
				scope.Like(dbCustomer, val)
				return nil
			}).
			AddString(erp.ListPlanReq_ListOptMark, func(val string) error {
				scope.Like(dbMark, val)
				return nil
			}).
			AddUint32Range(erp.ListPlanReq_ListOptStatTimeRange, func(begin, end uint32) error {
				scope.Where(fmt.Sprintf("%s >= ? AND %s <= ?", dbPlanAt, dbPlanAt), begin, end)
				return nil
			}).
			Process()
		if err != nil {
			log.Error(err.Error())
			return nil, nil, err
		}
	}

	var incomeList []*erp.ModelPlan
	var paginate *listoption.Paginate
	paginate, err = scope.PaginateQuery(listOption, &incomeList)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, nil, err
	}

	return incomeList, paginate, nil
}

var dbPlanDetail = NewTPlanDetail(db.Db())

type TPlanDetail struct {
	db    *gorm.DB
	model *erp.ModelPlanDetail
}

func NewTPlanDetail(db *gorm.DB) *TPlanDetail {
	return &TPlanDetail{
		db:    db,
		model: &erp.ModelPlanDetail{},
	}
}

func (d *TPlanDetail) newScope() *dbx.Scope {
	if d.db == nil {
		d.db = db.Db()
	}
	return dbx.NewScope(d.db, &erp.ModelPlanDetail{})
}

func (d *TPlanDetail) Create(ctx context.Context, m *erp.ModelPlanDetail) error {
	return d.newScope().Create(ctx, &m)
}

func (d *TPlanDetail) Update(ctx context.Context, m *erp.ModelPlanDetail, whereMap map[string]interface{}) error {
	return d.newScope().Where(whereMap).Update(&m)
}

func (d *TPlanDetail) DeleteById(ctx context.Context, pk uint64) error {
	return d.newScope().Delete(&erp.ModelPlanDetail{}, pk)
}

func (d *TPlanDetail) DeleteByWhere(ctx context.Context, whereMap map[string]interface{}) error {
	return d.newScope().Delete(&erp.ModelPlanDetail{}, whereMap)
}

func (d *TPlanDetail) DeleteByIdList(ctx context.Context, idList []uint64) error {
	if len(idList) == 0 {
		return nil
	}
	return d.newScope().In(dbId, idList).Delete(&erp.ModelPlanDetail{})
}

func (d *TPlanDetail) GetOne(ctx context.Context, pk uint64) (*erp.ModelPlanDetail, error) {
	var m erp.ModelPlanDetail
	err := d.newScope().SetNotFoundErr(erp.ErrNotFoundPlanDetail).First(&m, pk)
	return &m, err
}

func (d *TPlanDetail) ListWithListOption(ctx context.Context, listOption *listoption.ListOption, accIdList []uint64) ([]*erp.ModelPlanDetail, *listoption.Paginate, error) {
	var err error
	scope := d.newScope().Order("created_at DESC")

	if len(accIdList) > 0 {
		scope.In(dbAccountId, accIdList)
	}

	if listOption != nil {

		err = listoption.NewProcessor(listOption).
			AddUint64(erp.ListPlanDetailReq_ListOptPlanId, func(val uint64) error {
				scope.Where(dbPlanId, val)
				return nil
			}).
			Process()
		if err != nil {
			log.Error(err.Error())
			return nil, nil, err
		}

	}

	var planDetailList []*erp.ModelPlanDetail
	var paginate *listoption.Paginate
	paginate, err = scope.PaginateQuery(listOption, &planDetailList)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, nil, err
	}

	return planDetailList, paginate, nil
}
