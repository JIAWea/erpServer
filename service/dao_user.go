package app

import (
	"context"
	"github.com/JIAWea/erpServer/api/erp"
)

func (d *TUser) GetOneByName(ctx context.Context, name string) (*erp.ModelUser, error) {
	var m erp.ModelUser
	err := d.newScope().SetNotFoundErr(erp.ErrNotFoundUser).
		Eq(dbName, name).
		Eq(dbStatus, uint32(erp.ModelUser_StatusEnable)).
		First(&m)
	return &m, err
}
