package main

import (
	"gin-gorm/model"
	"gin-gorm/routers"
)

func main() {
    // 初始化数据库
    model.InitDb()
    // 初始化路由
	routers.InitRouter()
}