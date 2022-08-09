package routers

import (
	"gin-gorm/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
    r := gin.Default()
    
    routerV1 := r.Group("api/v1")
    {
        // 用户模块的路由接口

        // 分类模块的路由接口

        // 文章模块的路由接口
    }

    r.Run(":" + utils.HttpPort)
}