package routers

import (
	v1 "gin-gorm/api/v1"
	"gin-gorm/middleware"
	"gin-gorm/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
    
    r.Use(middleware.Cors())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
        auth.POST("/user/add",v1.AddUser)
		auth.PUT("/user/:id", v1.EditUser)
		auth.DELETE("/user/:id", v1.DeleteUser)

		// 分类模块的路由接口
		auth.POST("/category/add", v1.AddCategory)
		auth.PUT("/category/:id", v1.EditCategory)
		auth.DELETE("/category/:id", v1.DeleteCategory)

		// 文章模块的路由接口
		auth.POST("/article/add", v1.AddArticle)
		auth.PUT("/article/:id", v1.EditArticle)
		auth.DELETE("/article/:id", v1.DeleteArticle)

        // 上传文件
        auth.POST("/upload", v1.Upload)
	}

	router := r.Group("api/v1")
	{
        router.POST("/login", v1.Login)
		router.GET("/users", v1.GetUsers)
		router.GET("/user/:id", v1.GetUser)
		router.GET("/categories", v1.GetCategories)
		router.GET("/articles", v1.GetArtilces)
		router.GET("/articles/category/:id", v1.GetCategoryArticles)
		router.GET("/article/:id", v1.GetArtilce)
	}

	r.Run(":" + utils.HttpPort)
}
