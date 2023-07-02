package erp

import (
	"context"
)




func CreateErp(ctx context.Context, req *CreateErpReq) (*CreateErpRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.CreateErp(ctx, req)
}

func UpdateErp(ctx context.Context, req *UpdateErpReq) (*UpdateErpRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.UpdateErp(ctx, req)
}

func DeleteErp(ctx context.Context, req *DeleteErpReq) (*DeleteErpRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.DeleteErp(ctx, req)
}

func GetErp(ctx context.Context, req *GetErpReq) (*GetErpRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.GetErp(ctx, req)
}

func ListErp(ctx context.Context, req *ListErpReq) (*ListErpRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.ListErp(ctx, req)
}

func DDListErp(ctx context.Context, req *ListErpReq) (*ListErpRsp, error) {
	if cliMgr.conn == nil {
		return nil, cliMgr.initErr
	}
    return cliMgr.cli.DDListErp(ctx, req)
}


