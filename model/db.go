package model

import (
	"fmt"
	"gin-gorm/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func InitDb() {
    // 连接数据库
	db, err = gorm.Open(utils.Db,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			utils.DbUser,
			utils.DbPassword,
			utils.DbHost,
			utils.DbPort,
			utils.DbName),
	)
	if err != nil {
		fmt.Println("连接数据库失败, 请检查",err)
	}
    // 设置数据库表名为单数形式
    db.SingularTable(false)

    // 创建表
    db.AutoMigrate(&User{}, &Category{}, &Article{})

    // defer db.Close()

    // 设置连接数据库时间
    db.DB().SetMaxIdleConns(10)
    db.DB().SetMaxOpenConns(100)
    db.DB().SetConnMaxLifetime(10*time.Second)
}
