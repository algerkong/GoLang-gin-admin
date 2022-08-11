package v1

import (
	"gin-gorm/middleware"
	"gin-gorm/model"
	"gin-gorm/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context){
    var data model.User
    c.ShouldBindJSON(&data)
    code := model.CheckLogin(data.Username,data.Password)
    var token string
    if code == errmsg.SUCCESS {
        token,_ = middleware.SetToken(data.Username)
    }
    c.JSON(errmsg.GetHttpReq(code), gin.H{
        "code": code,
        "msg":  errmsg.GetErrMsg(code),
        "token": token,
    })

}