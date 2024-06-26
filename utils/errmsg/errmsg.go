package errmsg

import "net/http"

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
	ERROR_USER_DELTED      = 1008
    ERROR_USER_NO_RIGHT    = 1009
	// code = 2000  文章模块错误
	ARTICLE_NOT_EXIST       = 2001
	ERROR_ARTICLE_NOT_EXIST = 2003
	// code = 3000  分类模块错误
	ERROR_CATEGORY_USED      = 3001
	ERROR_CATEGORY_NOT_EXIST = 3003
)

var codeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",
	// code = 1000  用户模块错误
	ERROR_USER_USED:        "用户已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_USER_DELTED:      "用户已被删除",
	ERROR_TOKEN_EXIST:      "token不存在",
	ERROR_TOKEN_RUNTIME:    "token过期",
	ERROR_TOKEN_WRONG:      "token错误",
	ERROR_TOKEN_TYPE_WRONG: "token格式错误",
    ERROR_USER_NO_RIGHT : "用户没有权限",
	// code = 3000  分类模块错误
	ERROR_CATEGORY_USED:      "分类已存在",
	ERROR_CATEGORY_NOT_EXIST: "分类不存在",

	// code = 2000  文章模块错误
	ERROR_ARTICLE_NOT_EXIST: "文章不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}

func GetHttpReq(code int) int {
	if code != SUCCESS {
		return http.StatusBadRequest
	}
	return http.StatusOK
}
