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

var dbUserAccount = NewTUserAccount(db.Db())

type TUserAccount struct {
	db    *gorm.DB
	model *erp.ModelUserAccount
}

func NewTUserAccount(db *gorm.DB) *TUserAccount {
	return &TUserAccount{
		db:    db,
		model: &erp.ModelUserAccount{},
	}
}

func (d *TUserAccount) newScope() *dbx.Scope {
	if d.db == nil {
		d.db = db.Db()
	}
	return dbx.NewScope(d.db, &erp.ModelUserAccount{})
}

func (d *TUserAccount) Create(ctx context.Context, m *erp.ModelUserAccount) error {
	return d.newScope().Create(ctx, &m)
}

func (d *TUserAccount) CreateInBatches(ctx context.Context, list []*erp.ModelUserAccount) error {
	if len(list) == 0 {
		return nil
	}
	return d.newScope().CreateInBatches(list, 100)
}

func (d *TUserAccount) Update(ctx context.Context, m *erp.ModelUserAccount, whereMap map[string]interface{}) error {
	return d.newScope().Where(whereMap).Update(&m)
}

func (d *TUserAccount) DeleteById(ctx context.Context, pk uint64) error {
	return d.newScope().Delete(&erp.ModelUserAccount{}, pk)
}

func (d *TUserAccount) DeleteByWhere(ctx context.Context, whereMap map[string]interface{}) error {
	if len(whereMap) == 0 {
		return nil
	}
	return d.newScope().Where(whereMap).Delete(&erp.ModelUserAccount{})
}

func (d *TUserAccount) GetOne(ctx context.Context, pk uint64) (*erp.ModelUserAccount, error) {
	var m erp.ModelUserAccount
	err := d.newScope().SetNotFoundErr(erp.ErrNotFoundUserAccount).First(&m, pk)
	return &m, err
}

func (d *TUserAccount) ListWithListOption(ctx context.Context, listOption *listoption.ListOption, whereOpts interface{}) ([]*erp.ModelUserAccount, *listoption.Paginate, error) {
	var err error
	scope := d.newScope().Where(whereOpts)
	if listOption != nil {

	}

	var userAccountList []*erp.ModelUserAccount
	var paginate *listoption.Paginate
	paginate, err = scope.PaginateQuery(listOption, &userAccountList)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, nil, err
	}

	return userAccountList, paginate, nil
}

func (d *TUserAccount) GetIdListByUserId(ctx context.Context, userId uint64) ([]uint64, error) {
	var list []*erp.ModelUserAccount
	err := d.newScope().Eq(dbUserId, userId).Find(&list)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}

	var accIdList []uint64
	for _, v := range list {
		accIdList = append(accIdList, v.AccountId)
	}
	return accIdList, nil
}
