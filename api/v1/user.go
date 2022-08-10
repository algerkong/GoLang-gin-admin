package v1

import (
	"gin-gorm/model"
	"gin-gorm/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 查询用户是否存在
func UserExist(c *gin.Context) {

}

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		code = model.CreateUser(&data)
	}
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": errmsg.SUCCESS,
		"msg":  errmsg.GetErrMsg(errmsg.SUCCESS),
		"data": data,
	})

}

// 查询单个用户
func GetUser(c *gin.Context) {
    
}

// 查询所有用户
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetUsers(pageSize, pageNum)
	if data == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errmsg.ERROR,
			"msg":  errmsg.GetErrMsg(errmsg.ERROR),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": errmsg.SUCCESS,
		"msg":  errmsg.GetErrMsg(errmsg.SUCCESS),
		"data": data,
	})
}

// 编辑用户
func EditUser(c *gin.Context) {}

// 删除用户
func DeleteUser(c *gin.Context) {}
