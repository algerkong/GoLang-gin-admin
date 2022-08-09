package erromsg

const (
	SUCCESS = 200
	ERROR   = 500
	// code = 1000  用户模块错误
	ERROR_USER_USED        = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	// code = 2000  文章模块错误
	ARTICLE_NOT_EXIST = 2001
	// code = 3000  分类模块错误
	CATEGORY_NOT_EXIST = 3001
)

var codeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",
	// code = 1000  用户模块错误
	ERROR_USER_USED:        "用户已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "token不存在",
	ERROR_TOKEN_RUNTIME:    "token过期",
	ERROR_TOKEN_WRONG:      "token错误",
	ERROR_TOKEN_TYPE_WRONG: "token格式错误",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
