package app

import (
	"context"
	"fmt"
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

func (d *TExpense) Update(ctx context.Context, id uint64, data map[string]interface{}) error {
	return d.newScope().Where(dbId, id).Update(data)
}

func (d *TExpense) DeleteById(ctx context.Context, pk uint64) error {
	return d.newScope().Delete(&erp.ModelExpense{}, pk)
}

func (d *TExpense) DeleteByIdList(ctx context.Context, idList []uint64) error {
	if len(idList) == 0 {
		return nil
	}
	return d.newScope().In(dbId, idList).Delete(&erp.ModelExpense{})
}

func (d *TExpense) DeleteByWhere(ctx context.Context, whereMap map[string]interface{}) error {
	return d.newScope().Delete(&erp.ModelExpense{}, whereMap)
}

func (d *TExpense) GetOne(ctx context.Context, pk uint64) (*erp.ModelExpense, error) {
	var m erp.ModelExpense
	err := d.newScope().SetNotFoundErr(erp.ErrNotFoundExpense).First(&m, pk)
	return &m, err
}

func (d *TExpense) ListWithListOption(ctx context.Context, listOption *listoption.ListOption, accIdList []uint64) ([]*erp.ModelExpense, *listoption.Paginate, error) {
	var err error
	scope := d.newScope().Order("created_at DESC")

	if len(accIdList) > 0 {
		scope.In(dbAccountId, accIdList)
	}

	if listOption != nil {
		err = listoption.NewProcessor(listOption).
			AddString(erp.ListExpenseReq_ListOptAccountName, func(val string) error {
				var accList []*erp.ModelAccount
				err = dbAccount.newScope().Like(dbName, val).Find(&accList)
				if err != nil {
					return err
				}
				if len(accList) == 0 {
					scope.Eq("1", 0)
					return nil
				}
				var idList []uint64
				for _, v := range accList {
					idList = append(idList, v.Id)
				}
				scope.In(dbAccountId, idList)
				return nil
			}).
			AddString(erp.ListExpenseReq_ListOptMark, func(val string) error {
				scope.Like(dbMark, val)
				return nil
			}).
			AddUint32(erp.ListExpenseReq_ListOptPayMoney, func(val uint32) error {
				scope.Eq(dbPayMoney, val)
				return nil
			}).
			AddUint32Range(erp.ListExpenseReq_ListOptStatTimeRange, func(begin, end uint32) error {
				scope.Where(fmt.Sprintf("%s >= ? AND %s <= ?", dbPayAt, dbPayAt), begin, end)
				return nil
			}).
			Process()
		if err != nil {
			log.Error(err.Error())
			return nil, nil, err
		}
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
