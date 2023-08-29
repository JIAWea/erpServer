package app

import (
	"context"
	"fmt"

	"github.com/JIAWea/erpServer/api/erp"
	"github.com/ml444/gkit/core"
	"github.com/ml444/gkit/errorx"
)

// CheckPerm 权限校验
func CheckPerm(ctx context.Context, req interface{}) (context.Context, interface{}, error) {
	val := ctx.Value(core.HeadersKey)
	header, ok := val.(core.Header)
	if !ok {
		return ctx, nil, errorx.New(erp.ErrInternal)
	}

	path := header[core.HttpHeaderPath]
	fmt.Println("path===> ", path)

	return ctx, req, nil
}
