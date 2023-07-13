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
	db.RegisterModel(&erp.ModelUser{}, &erp.ModelRole{}, &erp.ModelMenu{}, &erp.ModelUserRole{}, &erp.ModelRoleMenu{})
}

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

var dbAccount = NewTAccount(db.Db())

type TAccount struct {
	db    *gorm.DB
	model *erp.ModelAccount
}

func NewTAccount(db *gorm.DB) *TAccount {
	return &TAccount{
		db:    db,
		model: &erp.ModelAccount{},
	}
}

func (d *TAccount) newScope() *dbx.Scope {
	if d.db == nil {
		d.db = db.Db()
	}
	return dbx.NewScope(d.db, &erp.ModelAccount{})
}

func (d *TAccount) Create(ctx context.Context, m *erp.ModelAccount) error {
	return d.newScope().Create(ctx, &m)
}

func (d *TAccount) Update(ctx context.Context, m *erp.ModelAccount, whereMap map[string]interface{}) error {
	return d.newScope().Where(whereMap).Update(&m)
}

func (d *TAccount) DeleteById(ctx context.Context, pk uint64) error {
	return d.newScope().Delete(&erp.ModelAccount{}, pk)
}

func (d *TAccount) DeleteByWhere(ctx context.Context, whereMap map[string]interface{}) error {
	return d.newScope().Delete(&erp.ModelAccount{}, whereMap)
}

func (d *TAccount) GetOne(ctx context.Context, pk uint64) (*erp.ModelAccount, error) {
	var m erp.ModelAccount
	err := d.newScope().SetNotFoundErr(erp.ErrNotFoundAccount).First(&m, pk)
	return &m, err
}

func (d *TAccount) ListWithListOption(ctx context.Context, listOption *listoption.ListOption, whereOpts interface{}) ([]*erp.ModelAccount, *listoption.Paginate, error) {
	var err error
	scope := d.newScope().Where(whereOpts)
	if listOption != nil {

		err = listoption.NewProcessor(listOption).
			AddString(erp.ListAccountReq_ListOptName, func(val string) error {
				return nil
			}).
			AddUint32Range(erp.ListAccountReq_ListOptTimeRange, func(begin, end uint32) error {
				return nil
			}).
			Process()
		if err != nil {
			log.Error(err.Error())
			return nil, nil, err
		}

	}

	var accountList []*erp.ModelAccount
	var paginate *listoption.Paginate
	paginate, err = scope.PaginateQuery(listOption, &accountList)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, nil, err
	}

	return accountList, paginate, nil
}
