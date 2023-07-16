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

var dbUserRole = NewTUserRole(db.Db())

type TUserRole struct {
	db    *gorm.DB
	model *erp.ModelUserRole
}

func NewTUserRole(db *gorm.DB) *TUserRole {
	return &TUserRole{
		db:    db,
		model: &erp.ModelUserRole{},
	}
}

func (d *TUserRole) newScope() *dbx.Scope {
	if d.db == nil {
		d.db = db.Db()
	}
	return dbx.NewScope(d.db, &erp.ModelUserRole{})
}

func (d *TUserRole) Create(ctx context.Context, m *erp.ModelUserRole) error {
	return d.newScope().Create(ctx, &m)
}

func (d *TUserRole) CreateInBatches(ctx context.Context, list []*erp.ModelUserRole) error {
	if len(list) == 0 {
		return nil
	}
	return d.newScope().CreateInBatches(list, 100)
}

func (d *TUserRole) Update(ctx context.Context, m *erp.ModelUserRole, whereMap map[string]interface{}) error {
	return d.newScope().Where(whereMap).Update(&m)
}

func (d *TUserRole) DeleteById(ctx context.Context, pk uint64) error {
	return d.newScope().Delete(&erp.ModelUserRole{}, pk)
}

func (d *TUserRole) DeleteByWhere(ctx context.Context, whereMap map[string]interface{}) error {
	if len(whereMap) == 0 {
		return nil
	}
	return d.newScope().Where(whereMap).Delete(&erp.ModelUserRole{})
}

func (d *TUserRole) GetOne(ctx context.Context, pk uint64) (*erp.ModelUserRole, error) {
	var m erp.ModelUserRole
	err := d.newScope().SetNotFoundErr(erp.ErrNotFoundUserRole).First(&m, pk)
	return &m, err
}

func (d *TUserRole) ListWithListOption(ctx context.Context, listOption *listoption.ListOption, whereOpts interface{}) ([]*erp.ModelUserRole, *listoption.Paginate, error) {
	var err error
	scope := d.newScope().Where(whereOpts)
	if listOption != nil {

	}

	var userRoleList []*erp.ModelUserRole
	var paginate *listoption.Paginate
	paginate, err = scope.PaginateQuery(listOption, &userRoleList)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, nil, err
	}

	return userRoleList, paginate, nil
}
