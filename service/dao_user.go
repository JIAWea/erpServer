package app

import (
	"context"
	"errors"
	"github.com/JIAWea/erpServer/api/erp"
	"github.com/JIAWea/erpServer/internal/db"
	"github.com/JIAWea/erpServer/pkg/utils"
	"github.com/ml444/gkit/dbx"
	"github.com/ml444/gkit/errorx"
	"github.com/ml444/gkit/listoption"
	log "github.com/ml444/glog"
	"gorm.io/gorm"
)

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

func (d *TUser) IsCreator(ctx context.Context, pk uint64) bool {
	var m erp.ModelUser
	err := d.newScope().SetNotFoundErr(erp.ErrNotFoundUser).First(&m, pk)
	if err != nil {
		return false
	}
	if m.IsCreator {
		return true
	}

	return false
}

func (d *TUser) CheckPassword(ctx context.Context, pk uint64, pwd string) error {
	var m erp.ModelUser
	err := d.newScope().SetNotFoundErr(erp.ErrNotFoundUser).First(&m, pk)
	if err != nil {
		return err
	}
	return utils.ComparePasswd(m.Password, pwd)
}

func (d *TUser) ListWithListOption(ctx context.Context, listOption *listoption.ListOption, whereOpts interface{}) ([]*erp.ModelUser, *listoption.Paginate, error) {
	var err error
	scope := d.newScope().Eq(dbIsCreator, uint32(0)).Preload("RoleList").Order("created_at DESC")
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

	for _, user := range userList {
		user.Password = ""
	}

	return userList, paginate, nil
}

func (d *TUser) GetOneByName(ctx context.Context, name string) (*erp.ModelUser, error) {
	var m erp.ModelUser
	err := d.newScope().SetNotFoundErr(erp.ErrNotFoundUser).
		Eq(dbName, name).
		First(&m)
	return &m, err
}

func (d *TUser) IsNameExist(ctx context.Context, name string) (bool, error) {
	var m erp.ModelUser
	err := d.newScope().SetNotFoundErr(erp.ErrNotFoundUser).Eq(dbName, name).First(&m)
	if errors.Is(err, errorx.New(erp.ErrNotFoundUser)) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
