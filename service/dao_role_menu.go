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

var dbRoleMenu = NewTRoleMenu(db.Db())

type TRoleMenu struct {
	db    *gorm.DB
	model *erp.ModelRoleMenu
}

func NewTRoleMenu(db *gorm.DB) *TRoleMenu {
	return &TRoleMenu{
		db:    db,
		model: &erp.ModelRoleMenu{},
	}
}

func (d *TRoleMenu) newScope() *dbx.Scope {
	if d.db == nil {
		d.db = db.Db()
	}
	return dbx.NewScope(d.db, &erp.ModelRoleMenu{})
}

func (d *TRoleMenu) Create(ctx context.Context, m *erp.ModelRoleMenu) error {
	return d.newScope().Create(ctx, &m)
}

func (d *TRoleMenu) CreateInBatches(ctx context.Context, list []*erp.ModelRoleMenu) error {
	if len(list) == 0 {
		return nil
	}
	return d.newScope().CreateInBatches(list, 100)
}

func (d *TRoleMenu) Update(ctx context.Context, m *erp.ModelRoleMenu, whereMap map[string]interface{}) error {
	return d.newScope().Where(whereMap).Update(&m)
}

func (d *TRoleMenu) DeleteById(ctx context.Context, pk uint64) error {
	return d.newScope().Delete(&erp.ModelRoleMenu{}, pk)
}

func (d *TRoleMenu) DeleteByWhere(ctx context.Context, whereMap map[string]interface{}) error {
	return d.newScope().Where(whereMap).Delete(&erp.ModelRoleMenu{})
}

func (d *TRoleMenu) GetOne(ctx context.Context, pk uint64) (*erp.ModelRoleMenu, error) {
	var m erp.ModelRoleMenu
	err := d.newScope().SetNotFoundErr(erp.ErrNotFoundRoleMenu).First(&m, pk)
	return &m, err
}

func (d *TRoleMenu) ListWithListOption(ctx context.Context, listOption *listoption.ListOption, whereOpts interface{}) ([]*erp.ModelRoleMenu, *listoption.Paginate, error) {
	var err error
	scope := d.newScope().Where(whereOpts)
	if listOption != nil {

	}

	var roleMenuList []*erp.ModelRoleMenu
	var paginate *listoption.Paginate
	paginate, err = scope.PaginateQuery(listOption, &roleMenuList)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, nil, err
	}

	return roleMenuList, paginate, nil
}
