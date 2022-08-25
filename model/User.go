package model

import (
	"encoding/base64"
	"gin-gorm/utils/errmsg"
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null;unique" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Age      int    `gorm:"type:int" json:"age" validate:"min=1,max=120" label:"年龄"`
	Role     int    `gorm:"type:int; default:2" json:"role" validate:"required,gte=2" label:"角色码"`
	Avatar   string `gorm:"type:varchar(100)" json:"avatar" validate:"max=100" label:"头像"`
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

// 查询用户是否存在
func CheckUserID(id int) (code int) {
    var user User
    db.Where("id = ?", id).First(&user)
    if user.ID > 0 {
        return errmsg.SUCCESS
    }
    return errmsg.ERROR_USER_NOT_EXIST
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
func GetUsers(pageSize, pageNum int) ([]User,int) {
	var users []User
    var total int
    db.Model(&User{}).Count(&total)
    err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
    if err.Error != nil && err.Error != gorm.ErrRecordNotFound {
        return nil,0
    }
	return users,total
}

// 编辑用户
func EditUser (id int,data *User) int {
    userMaps := make(map[string]interface{})
    if data.Username != ""{
        userMaps["username"] = data.Username
    }
    if data.Age != 0{
        userMaps["age"] = data.Age
    }
    if data.Role != 0{
        userMaps["role"] = data.Role
    }
    if data.Avatar != ""{
        userMaps["avatar"] = data.Avatar
    }
    err := db.Model(&User{}).Where("id = ?", id).Updates(userMaps)
    if err.Error != nil {
        return errmsg.ERROR
    }
    return errmsg.SUCCESS
}

// 删除用户
func DeleteUser(id int) int {
    err := db.Where("id = ?", id).Delete(&User{})
    if err.Error != nil {
        return errmsg.ERROR
    }
    return errmsg.SUCCESS
}

// 框架调用函数
func (u *User)BeforeSave(){
    u.Password = ScryptPwd(u.Password)
}

func ScryptPwd(password string) string{
    HashPwd,err := scrypt.Key([]byte(password), []byte("scrypt"), 16384, 8, 1, 10)
    if err != nil {
        log.Fatal(err)
    }
    return base64.StdEncoding.EncodeToString(HashPwd)
}

// 登录验证
func CheckLogin (username,password string) (User,int) {
    var user User
    db.Where("username = ?", username).First(&user)
    if user.ID > 0 {
        if ScryptPwd(password) != user.Password {
            return user, errmsg.ERROR_PASSWORD_WRONG
        }
        return user,errmsg.SUCCESS
    }
    if user.Role != 0{
        return user,errmsg.ERROR_USER_NO_RIGHT
    }
    return user,errmsg.ERROR_USER_NOT_EXIST
}