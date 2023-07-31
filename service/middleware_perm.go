package app

import (
	"context"
	log "github.com/ml444/glog"
)

// CheckPerm 权限校验
func CheckPerm(ctx context.Context, req interface{}) (context.Context, interface{}, error) {
	log.Info("hhh...")
	return ctx, req, nil
}
