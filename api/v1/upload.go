package v1

import (
	"gin-gorm/model"
	"gin-gorm/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	url, code := model.UploadFile(file, fileSize)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
		"data": url,
	})
}
