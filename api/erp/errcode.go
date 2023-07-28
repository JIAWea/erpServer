package erp


const (
        Success int32	=	0
        ErrInternal int32	=	99999
        ErrInvalidParam int32	=	1
        ErrParamRequired int32	=	2
        ErrUserNotExist int32	=	3998
        ErrUserDisable int32	=	3999
        ErrAuthContext int32	=	4000
        ErrTokenUnverifiable int32	=	4001
        ErrInvalidToken int32	=	4002
        ErrTokenExpired int32	=	4003
        ErrNotFoundUser int32	=	10000
        ErrNotFoundRole int32	=	10001
        ErrNotFoundMenu int32	=	10002
        ErrNotFoundUserRole int32	=	10003
        ErrNotFoundRoleMenu int32	=	10004
        ErrPasswordFormatInvalid int32	=	10005
        ErrPassword int32	=	10006
        ErrNotFoundAccount int32	=	10007
        ErrUserExist int32	=	10008
        ErrNotFoundExpense int32	=	10009
        ErrNotFoundIncome int32	=	10010
        ErrExpenseCategoryInvalid int32	=	10011
        ErrNotFoundUserAccount int32	=	10012
)


var ErrCodeMap = map[int32]string{
        Success: "Success",
        ErrInternal: "系统错误",
        ErrInvalidParam: "非法参数",
        ErrParamRequired: "缺失参数",
        ErrUserNotExist: "该用户不存在",
        ErrUserDisable: "账号已禁用",
        ErrAuthContext: "认证失败",
        ErrTokenUnverifiable: "认证异常",
        ErrInvalidToken: "认证错误",
        ErrTokenExpired: "认证过期",
        ErrNotFoundUser: "未找到相关记录",
        ErrNotFoundRole: "未找到相关记录",
        ErrNotFoundMenu: "未找到相关记录",
        ErrNotFoundUserRole: "未找到相关记录",
        ErrNotFoundRoleMenu: "未找到相关记录",
        ErrPasswordFormatInvalid: "密码不合法",
        ErrPassword: "密码错误",
        ErrNotFoundAccount: "未找到相关记录",
        ErrUserExist: "用户名已存在",
        ErrNotFoundExpense: "记录不存在",
        ErrNotFoundIncome: "记录不存在",
        ErrExpenseCategoryInvalid: "支出类目错误",
        ErrNotFoundUserAccount: "记录不存在",
}

var ErrCode4StatusCodeMap = map[int32]int32{
        ErrInternal: 500,
        ErrInvalidParam: 400,
        ErrParamRequired: 400,
        ErrUserNotExist: 401,
        ErrUserDisable: 401,
        ErrAuthContext: 401,
        ErrTokenUnverifiable: 401,
        ErrInvalidToken: 401,
        ErrTokenExpired: 401,
        ErrNotFoundUser: 404,
        ErrNotFoundRole: 404,
        ErrNotFoundMenu: 404,
        ErrNotFoundUserRole: 404,
        ErrNotFoundRoleMenu: 404,
        ErrPasswordFormatInvalid: 400,
        ErrPassword: 400,
        ErrNotFoundAccount: 404,
        ErrUserExist: 400,
        ErrNotFoundExpense: 404,
        ErrNotFoundIncome: 404,
        ErrExpenseCategoryInvalid: 400,
        ErrNotFoundUserAccount: 404,
}
