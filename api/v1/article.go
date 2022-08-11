package v1

import (
	"gin-gorm/model"
	"gin-gorm/utils/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code := model.CheckCategoryID(data.CategoryId)
	if code == errmsg.SUCCESS {
		code = model.AddArticle(&data)
	}
	c.JSON(errmsg.GetHttpReq(code), gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
		"data": data,
	})
}

// 查询文章
func GetArtilce(c *gin.Context) {
	idInt, _ := strconv.Atoi(c.Param("id"))
	id := uint(idInt)
	data, code := model.GetArtilce(id)
	c.JSON(errmsg.GetHttpReq(code), gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
		"data": data,
	})
}

// 查询单个分类下的文章
func GetCategoryArticles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	idInt, _ := strconv.Atoi(c.Param("id"))
	id := uint(idInt)
	var data []model.Article
	var code int
	var total int
	code = model.CheckCategoryID(id)
	if code == errmsg.SUCCESS {
		data, total, code = model.GetCategoryArticles(id, pageSize, pageNum)
	}
	c.JSON(errmsg.GetHttpReq(code), gin.H{
		"code":  code,
		"msg":   errmsg.GetErrMsg(code),
		"data":  data,
		"total": total,
	})

}

// 查询所有文章
func GetArtilces(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, total, code := model.GetArtilces(pageSize, pageNum)

	c.JSON(errmsg.GetHttpReq(code), gin.H{
		"code":  code,
		"msg":   errmsg.GetErrMsg(code),
		"data":  data,
		"total": total,
	})
}

// 编辑文章
func EditArticle(c *gin.Context) {
	idInt, _ := strconv.Atoi(c.Param("id"))
	id := uint(idInt)
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code := model.CheckArticleID(id)
	if code == errmsg.SUCCESS {
		code = model.EditArticle(id, &data)
	}
	if code != errmsg.SUCCESS {
		c.JSON(errmsg.GetHttpReq(code), gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
		})
		return
	}
	if data.CategoryId != 0 {
		code = model.CheckCategoryID(data.CategoryId)
	}
	c.JSON(errmsg.GetHttpReq(code), gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})
}

// 删除文章
func DeleteArticle(c *gin.Context) {
	idInt, _ := strconv.Atoi(c.Param("id"))
	id := uint(idInt)
	code := model.CheckArticleID(id)
	if code == errmsg.SUCCESS {
		code = model.DeleteArticle(id)
	}
	c.JSON(errmsg.GetHttpReq(code), gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})

}
