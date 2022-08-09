package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
    gorm.Model
    Title string `gorm:"type:varchar(100);not null" json:"title"`
    Content string `gorm:"type:longtext;not null" json:"content"`
    Desc string `gorm:"type:varchar(200);not null" json:"desc"`
    Img string `gorm:"type:varchar(100)" json:"img"`
    Author string `gorm:"type:varchar(20);not null" json:"author"`
    CategoryId uint `gorm:"type:int;not null" json:"category_id"`
    Category Category `gorm:"foreignkey:CategoryId" json:"category"`
}