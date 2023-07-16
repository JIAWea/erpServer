package app

import (
	"github.com/JIAWea/erpServer/api/erp"
)

type ErpService struct {
	erp.UnsafeErpServer
}

func NewErpService() ErpService {
	return ErpService{}
}
