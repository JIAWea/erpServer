package erp


const (
        Success int32	=	0
        ErrInvalidParam int32	=	1
        ErrParamRequired int32	=	2
        ErrNotFoundUser int32	=	10000
        ErrNotFoundRole int32	=	10001
        ErrNotFoundMenu int32	=	10002
        ErrNotFoundUserRole int32	=	10003
        ErrNotFoundRoleMenu int32	=	10004
)


var ErrCodeMap = map[int32]string{
        Success: "Success",
        ErrInvalidParam: "非法参数",
        ErrParamRequired: "缺失参数",
        ErrNotFoundUser: "未找到相关记录",
        ErrNotFoundRole: "未找到相关记录",
        ErrNotFoundMenu: "未找到相关记录",
        ErrNotFoundUserRole: "未找到相关记录",
        ErrNotFoundRoleMenu: "未找到相关记录",
}

var ErrCode4StatusCodeMap = map[int32]int32{
        ErrInvalidParam: 400,
        ErrParamRequired: 400,
        ErrNotFoundUser: 404,
        ErrNotFoundRole: 404,
        ErrNotFoundMenu: 404,
        ErrNotFoundUserRole: 404,
        ErrNotFoundRoleMenu: 404,
}
