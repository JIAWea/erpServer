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
	m.Id = 0
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
	scope := d.newScope().Order("created_at DESC")
	if listOption != nil {

		err = listoption.NewProcessor(listOption).
			AddString(erp.ListUserReq_ListOptName, func(val string) error {
				scope.Like(dbName, val)
				return nil
			}).
			AddUint32Range(erp.ListAccountReq_ListOptStatTimeRange, func(begin, end uint32) error {
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
