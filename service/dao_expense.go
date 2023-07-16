package app

import (
	"context"
	"github.com/ml444/gkit/dbx"
	"github.com/ml444/gkit/listoption"
	log "github.com/ml444/glog"
	"gorm.io/gorm"

	"github.com/JIAWea/erpServer/api/erp"
	"github.com/JIAWea/erpServer/internal/db"
)

var dbExpense = NewTExpense(db.Db())

type TExpense struct {
	db    *gorm.DB
	model *erp.ModelExpense
}

func NewTExpense(db *gorm.DB) *TExpense {
	return &TExpense{
		db:    db,
		model: &erp.ModelExpense{},
	}
}

func (d *TExpense) newScope() *dbx.Scope {
	if d.db == nil {
		d.db = db.Db()
	}
	return dbx.NewScope(d.db, &erp.ModelExpense{})
}

func (d *TExpense) Create(ctx context.Context, m *erp.ModelExpense) error {
	return d.newScope().Create(ctx, &m)
}

func (d *TExpense) Update(ctx context.Context, m *erp.ModelExpense, whereMap map[string]interface{}) error {
	return d.newScope().Where(whereMap).Update(&m)
}

func (d *TExpense) DeleteById(ctx context.Context, pk uint64) error {
	return d.newScope().Delete(&erp.ModelExpense{}, pk)
}

func (d *TExpense) DeleteByWhere(ctx context.Context, whereMap map[string]interface{}) error {
	return d.newScope().Delete(&erp.ModelExpense{}, whereMap)
}

func (d *TExpense) GetOne(ctx context.Context, pk uint64) (*erp.ModelExpense, error) {
	var m erp.ModelExpense
	err := d.newScope().SetNotFoundErr(erp.ErrNotFoundExpense).First(&m, pk)
	return &m, err
}

func (d *TExpense) ListWithListOption(ctx context.Context, listOption *listoption.ListOption, whereOpts interface{}) ([]*erp.ModelExpense, *listoption.Paginate, error) {
	var err error
	scope := d.newScope()
	if listOption != nil {

	}

	var expenseList []*erp.ModelExpense
	var paginate *listoption.Paginate
	paginate, err = scope.PaginateQuery(listOption, &expenseList)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, nil, err
	}

	return expenseList, paginate, nil
}
