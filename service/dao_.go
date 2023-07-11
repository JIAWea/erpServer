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

var dbUser = NewTUser(db.Db())

type TUser struct {
	db    *gorm.DB
	model *erp.ModelUser
}

func NewTUser(db *gorm.DB) *TUser {
	return &TUser{
		db:    db,
		model: &erp.ModelUser{},
	}
}

func (d *TUser) newScope() *dbx.Scope {
	if d.db == nil {
		d.db = db.Db()
	}
	return dbx.NewScope(d.db, &erp.ModelUser{})
}

func (d *TUser) Create(ctx context.Context, m *erp.ModelUser) error {
	m.Id = 0
	return d.newScope().Create(ctx, &m)
}

func (d *TUser) Update(ctx context.Context, m, whereMap map[string]interface{}) error {
	return d.newScope().Where(whereMap).Update(m)
}

func (d *TUser) DeleteById(ctx context.Context, pk uint64) error {
	return d.newScope().Delete(&erp.ModelUser{}, pk)
}

func (d *TUser) DeleteByWhere(ctx context.Context, whereMap map[string]interface{}) error {
	return d.newScope().Delete(&erp.ModelUser{}, whereMap)
}

func (d *TUser) DeleteByIdList(ctx context.Context, pkList []uint64) error {
	if len(pkList) == 0 {
		return nil
	}
	return d.newScope().In(dbId, pkList).Delete(&erp.ModelUser{})
}

func (d *TUser) GetOne(ctx context.Context, pk uint64) (*erp.ModelUser, error) {
	var m erp.ModelUser
	err := d.newScope().SetNotFoundErr(erp.ErrNotFoundUser).Preload("RoleList").First(&m, pk)
	m.Password = ""
	return &m, err
}

func (d *TUser) ListWithListOption(ctx context.Context, listOption *listoption.ListOption, whereOpts interface{}) ([]*erp.ModelUser, *listoption.Paginate, error) {
	var err error
	scope := d.newScope().Eq(dbIsCreator, uint32(0)).Preload("RoleList")
	if listOption != nil {
		err = listoption.NewProcessor(listOption).
			AddString(erp.ListUserReq_ListOptName, func(val string) error {
				scope.Like(dbName, val)
				return nil
			}).
			AddUint32(erp.ListUserReq_ListOptStatus, func(val uint32) error {
				scope.Eq(dbStatus, val)
				return nil
			}).
			AddUint64(erp.ListUserReq_ListOptId, func(val uint64) error {
				scope.Eq(dbId, val)
				return nil
			}).
			Process()
		if err != nil {
			log.Error(err.Error())
			return nil, nil, err
		}
	}

	var userList []*erp.ModelUser
	var paginate *listoption.Paginate
	paginate, err = scope.PaginateQuery(listOption, &userList)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, nil, err
	}

	return userList, paginate, nil
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
	return d.newScope().Create(ctx, &m)
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
	scope := d.newScope()
	if listOption != nil {

		err = listoption.NewProcessor(listOption).
			AddUint64List(erp.ListRoleReq_ListOptIdList, func(valList []uint64) error {
				return nil
			}).
			AddString(erp.ListRoleReq_ListOptName, func(val string) error {
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

var dbMenu = NewTMenu(db.Db())

type TMenu struct {
	db    *gorm.DB
	model *erp.ModelMenu
}

func NewTMenu(db *gorm.DB) *TMenu {
	return &TMenu{
		db:    db,
		model: &erp.ModelMenu{},
	}
}

func (d *TMenu) newScope() *dbx.Scope {
	if d.db == nil {
		d.db = db.Db()
	}
	return dbx.NewScope(d.db, &erp.ModelMenu{})
}

func (d *TMenu) Create(ctx context.Context, m *erp.ModelMenu) error {
	return d.newScope().Create(ctx, &m)
}

func (d *TMenu) Update(ctx context.Context, m *erp.ModelMenu, whereMap map[string]interface{}) error {
	return d.newScope().Where(whereMap).Update(&m)
}

func (d *TMenu) DeleteById(ctx context.Context, pk uint64) error {
	return d.newScope().Delete(&erp.ModelMenu{}, pk)
}

func (d *TMenu) DeleteByWhere(ctx context.Context, whereMap map[string]interface{}) error {
	return d.newScope().Delete(&erp.ModelMenu{}, whereMap)
}

func (d *TMenu) GetOne(ctx context.Context, pk uint64) (*erp.ModelMenu, error) {
	var m erp.ModelMenu
	err := d.newScope().SetNotFoundErr(erp.ErrNotFoundMenu).First(&m, pk)
	return &m, err
}

func (d *TMenu) ListWithListOption(ctx context.Context, listOption *listoption.ListOption, whereOpts interface{}) ([]*erp.ModelMenu, *listoption.Paginate, error) {
	var err error
	scope := d.newScope().Where(whereOpts)
	if listOption != nil {

		err = listoption.NewProcessor(listOption).
			AddUint64List(erp.ListMenuReq_ListOptIdList, func(valList []uint64) error {
				return nil
			}).
			AddString(erp.ListMenuReq_ListOptName, func(val string) error {
				return nil
			}).
			Process()
		if err != nil {
			log.Error(err.Error())
			return nil, nil, err
		}

	}

	var menuList []*erp.ModelMenu
	var paginate *listoption.Paginate
	paginate, err = scope.PaginateQuery(listOption, &menuList)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, nil, err
	}

	return menuList, paginate, nil
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
	return d.newScope().Delete(&erp.ModelUserRole{}, whereMap)
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

func (d *TRoleMenu) Update(ctx context.Context, m *erp.ModelRoleMenu, whereMap map[string]interface{}) error {
	return d.newScope().Where(whereMap).Update(&m)
}

func (d *TRoleMenu) DeleteById(ctx context.Context, pk uint64) error {
	return d.newScope().Delete(&erp.ModelRoleMenu{}, pk)
}

func (d *TRoleMenu) DeleteByWhere(ctx context.Context, whereMap map[string]interface{}) error {
	return d.newScope().Delete(&erp.ModelRoleMenu{}, whereMap)
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
