package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
    Username string `gorm:"type:varchar(20);not null;unique" json:"username"`
    Password string `gorm:"type:varchar(20);not null" json:"password"`
    Age      int    `gorm:"type:int" json:"age"`
    Role     int    `gorm:"type:int" json:"role"`
    Avatar   string `gorm:"type:varchar(100)" json:"avatar"`
}
