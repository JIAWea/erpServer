package app

import (
	"github.com/JIAWea/erpServer/api/erp"
	"github.com/JIAWea/erpServer/internal/db"
)

func init() {
	db.RegisterModel(
		&erp.ModelUser{},
		&erp.ModelRole{},
		&erp.ModelMenu{},
		&erp.ModelUserRole{},
		&erp.ModelRoleMenu{},
		&erp.ModelAccount{},
		&erp.ModelExpense{},
	)
}
