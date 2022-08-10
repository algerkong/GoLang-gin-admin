package model

import (
	"gin-gorm/utils/errmsg"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null;unique" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Age      int    `gorm:"type:int" json:"age"`
	Role     int    `gorm:"type:int" json:"role"`
	Avatar   string `gorm:"type:varchar(100)" json:"avatar"`
}

// 查询用户是否存在
func CheckUser(username string) (code int) {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USER_USED
	}
	return errmsg.SUCCESS
}

// 添加用户
func CreateUser(data *User) int {
	err := db.Create(&data)
	if err.Error != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询单个用户
func GetUser(id int) User {
    var user User
    db.Where("id = ?", id).First(&user)
    return user
}
    

// 查询所有用户
func GetUsers(pageSize, pageNum int) []User {
	var users []User
	err := db.Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&users)
    if err.Error != nil && err.Error != gorm.ErrRecordNotFound {
        return nil
    }
	return users
}

// 编辑用户

// 删除用户
