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


func init() {
	db.RegisterModel(&erp.ModelErp{},&erp.ModelRole{},)
}




var dbErp = NewTErp(db.Db())

type TErp struct {
	db    *gorm.DB
	model *erp.ModelErp
}

func NewTErp(db *gorm.DB) *TErp {
	return &TErp{
		db:    db,
		model: &erp.ModelErp{},
	}
}

func (d *TErp) newScope() *dbx.Scope {
    if d.db == nil {
		d.db = db.Db()
	}
	return dbx.NewScope(d.db, &erp.ModelErp{})
}

func (d *TErp) Create(ctx context.Context, m *erp.ModelErp) error {
	return d.newScope().Create(m)
}

func (d *TErp) Update(ctx context.Context, m *erp.ModelErp, whereMap map[string]interface{}) error {
	return d.newScope().Where(whereMap).Update(&m)
}

func (d *TErp) DeleteById(ctx context.Context, pk uint64) error {
	return d.newScope().Delete(&erp.ModelErp{}, pk)
}

func (d *TErp) DeleteByWhere(ctx context.Context, whereMap map[string]interface{}) error {
	return d.newScope().Delete(&erp.ModelErp{}, whereMap)
}

func (d *TErp) GetOne(ctx context.Context, pk uint64) (*erp.ModelErp, error) {
	var m erp.ModelErp
	err := d.newScope().SetNotFoundErr(erp.ErrNotFoundErp).First(&m, pk)
	return &m, err
}


func (d *TErp) ListWithListOption(ctx context.Context, listOption *listoption.ListOption, whereOpts interface{}) ([]*erp.ModelErp, *listoption.Paginate, error) {
	var err error
	scope := d.newScope().Where(whereOpts)
	if listOption != nil {





		err = listoption.NewProcessor(listOption).


            AddUint64List(erp.ListErpReq_ListOptIdList, func(valList []uint64) error {
            	return nil
            }).



            AddString(erp.ListErpReq_ListOptName, func(val string) error {
            	return nil
            }).



			Process()
		if err != nil {
			log.Error(err.Error())
			return nil, nil, err
		}


	}

    var erpList []*erp.ModelErp
	var paginate *listoption.Paginate
	paginate, err = scope.PaginateQuery(listOption, &erpList )
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, nil, err
	}

	return erpList , paginate, nil
}




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
	return d.newScope().Create(m)
}

func (d *TRole) Update(ctx context.Context, m *erp.ModelRole, whereMap map[string]interface{}) error {
	return d.newScope().Where(whereMap).Update(&m)
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
	scope := d.newScope().Where(whereOpts)
	if listOption != nil {



	}

    var roleList []*erp.ModelRole
	var paginate *listoption.Paginate
	paginate, err = scope.PaginateQuery(listOption, &roleList )
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, nil, err
	}

	return roleList , paginate, nil
}


