package erp


const (
        Success int32	=	0
        ErrInvalidParam int32	=	1
        ErrParamRequired int32	=	2
        ErrNotFoundErp int32	=	3
        ErrNotFoundRole int32	=	4
)


var ErrCodeMap = map[int32]string{
        Success: "Success",
        ErrInvalidParam: "非法参数",
        ErrParamRequired: "缺失参数",
        ErrNotFoundErp: "未找到相关记录",
        ErrNotFoundRole: "未找到相关记录",
}

var ErrCode4StatusCodeMap = map[int32]int32{
        ErrInvalidParam: 400,
        ErrParamRequired: 400,
        ErrNotFoundErp: 404,
        ErrNotFoundRole: 404,
}
