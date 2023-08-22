package app

import (
	"context"
	"fmt"

	"github.com/JIAWea/erpServer/api/erp"
	"github.com/JIAWea/erpServer/internal/db"
	"github.com/ml444/gkit/dbx"
	"github.com/ml444/gkit/listoption"
	log "github.com/ml444/glog"
	"gorm.io/gorm"
)

var dbIncome = NewTIncome(db.Db())

type TIncome struct {
	db    *gorm.DB
	model *erp.ModelIncome
}

func NewTIncome(db *gorm.DB) *TIncome {
	return &TIncome{
		db:    db,
		model: &erp.ModelIncome{},
	}
}

func (d *TIncome) newScope() *dbx.Scope {
	if d.db == nil {
		d.db = db.Db()
	}
	return dbx.NewScope(d.db, &erp.ModelIncome{})
}

func (d *TIncome) Create(ctx context.Context, m *erp.ModelIncome) error {
	return d.newScope().Create(ctx, &m)
}

func (d *TIncome) Update(ctx context.Context, id uint64, data map[string]interface{}) error {
	return d.newScope().Where(dbId, id).Update(data)
}

func (d *TIncome) DeleteById(ctx context.Context, pk uint64) error {
	return d.newScope().Delete(&erp.ModelIncome{}, pk)
}

func (d *TIncome) DeleteByIdList(ctx context.Context, idList []uint64) error {
	if len(idList) == 0 {
		return nil
	}
	return d.newScope().In(dbId, idList).Delete(&erp.ModelExpense{})
}

func (d *TIncome) DeleteByWhere(ctx context.Context, whereMap map[string]interface{}) error {
	return d.newScope().Delete(&erp.ModelIncome{}, whereMap)
}

func (d *TIncome) GetOne(ctx context.Context, pk uint64) (*erp.ModelIncome, error) {
	var m erp.ModelIncome
	err := d.newScope().SetNotFoundErr(erp.ErrNotFoundIncome).First(&m, pk)
	return &m, err
}

func (d *TIncome) ListWithListOption(ctx context.Context, listOption *listoption.ListOption, accIdList []uint64) ([]*erp.ModelIncome, *listoption.Paginate, error) {
	var err error
	scope := d.newScope().Order("created_at DESC")

	if len(accIdList) > 0 {
		scope.In(dbAccountId, accIdList)
	}

	if listOption != nil {
		err = listoption.NewProcessor(listOption).
			AddString(erp.ListIncomeReq_ListOptAccountName, func(val string) error {
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
			AddString(erp.ListIncomeReq_ListOptMark, func(val string) error {
				scope.Like(dbMark, val)
				return nil
			}).
			AddUint32(erp.ListIncomeReq_ListOptPayMoney, func(val uint32) error {
				scope.Eq(dbIncomeMoney, val)
				return nil
			}).
			AddUint32Range(erp.ListIncomeReq_ListOptStatTimeRange, func(begin, end uint32) error {
				scope.Where(fmt.Sprintf("%s >= ? AND %s <= ?", dbIncomeAt, dbIncomeAt), begin, end)
				return nil
			}).
			Process()
		if err != nil {
			log.Error(err.Error())
			return nil, nil, err
		}
	}

	var incomeList []*erp.ModelIncome
	var paginate *listoption.Paginate
	paginate, err = scope.PaginateQuery(listOption, &incomeList)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, nil, err
	}

	return incomeList, paginate, nil
}
