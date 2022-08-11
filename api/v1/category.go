package v1

import (
	"gin-gorm/model"
	"gin-gorm/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	c.ShouldBindJSON(&data)
	code := model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		code = model.AddCategory(&data)
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

// 查询所有分类
func GetCategories(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, total := model.GetCategories(pageSize, pageNum)
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
        "total": total,
	})

}

// 编辑分类
func EditCategory(c *gin.Context) {
	idInt, _ := strconv.Atoi(c.Param("id"))
	id := uint(idInt)
	var data model.Category
	c.ShouldBindJSON(&data)
	code := model.CheckCategoryID(id)
	if code == errmsg.SUCCESS {
		code = model.CheckCategory(data.Name)
	}
	if code == errmsg.SUCCESS {
		code = model.EditCategory(id, data.Name)
	}
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
		})
		return
	}
	data.ID = id
	c.JSON(http.StatusOK, gin.H{
		"code": errmsg.SUCCESS,
		"msg":  errmsg.GetErrMsg(errmsg.SUCCESS),
		"data": data,
	})
}

// 删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.CheckCategoryID(uint(id))
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
		})
		return
	}
	code = model.DeleteCategory(id)
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
	})
}
