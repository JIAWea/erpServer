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

var dbRole = NewTRole(db.Db())

type TRole struct {
	db    *gorm.DB
	model *erp.ModelRole
}

func NewTRole(db *gorm.DB) *TRole {
	return &TRole{
		db:    db,
		model: &erp.ModelRole{},
	}
}

func (d *TRole) newScope() *dbx.Scope {
	if d.db == nil {
		d.db = db.Db()
	}
	return dbx.NewScope(d.db, &erp.ModelRole{})
}

func (d *TRole) Create(ctx context.Context, m *erp.ModelRole) error {
	m.Id = 0
	return d.newScope().Create(ctx, &m)
}

func (d *TRole) Update(ctx context.Context, m *erp.ModelRole, whereMap map[string]interface{}) error {
	return d.newScope().Where(whereMap).Update(m)
}

func (d *TRole) DeleteById(ctx context.Context, pk uint64) error {
	return d.newScope().Delete(&erp.ModelRole{}, pk)
}

func (d *TRole) DeleteByWhere(ctx context.Context, whereMap map[string]interface{}) error {
	return d.newScope().Delete(&erp.ModelRole{}, whereMap)
}

func (d *TRole) GetOne(ctx context.Context, pk uint64) (*erp.ModelRole, error) {
	var m erp.ModelRole
	err := d.newScope().SetNotFoundErr(erp.ErrNotFoundRole).First(&m, pk)
	return &m, err
}

func (d *TRole) ListWithListOption(ctx context.Context, listOption *listoption.ListOption, whereOpts interface{}) ([]*erp.ModelRole, *listoption.Paginate, error) {
	var err error
	scope := d.newScope()
	if listOption != nil {

		err = listoption.NewProcessor(listOption).
			AddString(erp.ListRoleReq_ListOptName, func(val string) error {
				scope.Like(dbName, val)
				return nil
			}).
			AddUint32(erp.ListRoleReq_ListOptStatus, func(val uint32) error {
				scope.Eq(dbStatus, val)
				return nil
			}).
			AddUint64(erp.ListRoleReq_ListOptId, func(val uint64) error {
				scope.Eq(dbId, val)
				return nil
			}).
			Process()
		if err != nil {
			log.Error(err.Error())
			return nil, nil, err
		}

	}

	var roleList []*erp.ModelRole
	var paginate *listoption.Paginate
	paginate, err = scope.PaginateQuery(listOption, &roleList)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, nil, err
	}

	return roleList, paginate, nil
}
