package app

import (
	"context"

	"github.com/JIAWea/erpServer/api/erp"
	"github.com/JIAWea/erpServer/internal/db"
	"github.com/ml444/gkit/dbx"
	"github.com/ml444/gkit/listoption"
	log "github.com/ml444/glog"
	"gorm.io/gorm"
)

var dbIncome = NewTIncome(db.Db())

type TIncome struct {
	db    *gorm.DB
	model *erp.ModelIncome
}

func NewTIncome(db *gorm.DB) *TIncome {
	return &TIncome{
		db:    db,
		model: &erp.ModelIncome{},
	}
}

func (d *TIncome) newScope() *dbx.Scope {
	if d.db == nil {
		d.db = db.Db()
	}
	return dbx.NewScope(d.db, &erp.ModelIncome{})
}

func (d *TIncome) Create(ctx context.Context, m *erp.ModelIncome) error {
	return d.newScope().Create(ctx, &m)
}

func (d *TIncome) Update(ctx context.Context, m *erp.ModelIncome, whereMap map[string]interface{}) error {
	return d.newScope().Where(whereMap).Update(&m)
}

func (d *TIncome) DeleteById(ctx context.Context, pk uint64) error {
	return d.newScope().Delete(&erp.ModelIncome{}, pk)
}

func (d *TIncome) DeleteByWhere(ctx context.Context, whereMap map[string]interface{}) error {
	return d.newScope().Delete(&erp.ModelIncome{}, whereMap)
}

func (d *TIncome) GetOne(ctx context.Context, pk uint64) (*erp.ModelIncome, error) {
	var m erp.ModelIncome
	err := d.newScope().SetNotFoundErr(erp.ErrNotFoundIncome).First(&m, pk)
	return &m, err
}

func (d *TIncome) ListWithListOption(ctx context.Context, listOption *listoption.ListOption, whereOpts interface{}) ([]*erp.ModelIncome, *listoption.Paginate, error) {
	var err error
	scope := d.newScope().Where(whereOpts)
	if listOption != nil {

	}

	var incomeList []*erp.ModelIncome
	var paginate *listoption.Paginate
	paginate, err = scope.PaginateQuery(listOption, &incomeList)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, nil, err
	}

	return incomeList, paginate, nil
}
