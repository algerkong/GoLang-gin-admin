package model

import (
	"gin-gorm/utils/errmsg"

	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null;unique" json:"name"`
}

// 查询分类是否存在
func CheckCategory(name string) (code int) {
	var category Category
	db.Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return errmsg.ERROR_CATEGORY_USED
	}
	return errmsg.SUCCESS
}

// 查询分类Id是否存在
func CheckCategoryID(id uint) (code int) {
	var category Category
	db.Where("id = ?", id).First(&category)
	if category.ID > 0 {
		return errmsg.SUCCESS
	}
	return errmsg.ERROR_CATEGORY_NOT_EXIST
}

// 添加分类
func AddCategory(data *Category) int {
	err := db.Create(&data)
	if err.Error != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询单个分类
func GetCategory(id int) (category Category) {
	db.Where("id = ?", id).First(&category)
	return category
}

// 查询所有分类
func GetCategories(pageSize, pageNum int) ([]Category, int) {
	var categories []Category
	var total int
	db.Model(&Category{}).Count(&total)
	err := db.Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&categories)
	if err.Error != nil && err.Error != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return categories, total
}

// 编辑分类
func EditCategory(id uint, name string) int {
	category := make(map[string]interface{})
	category["Name"] = name
	err := db.Model(&Category{}).Where("id = ?", id).Updates(category)
	if err.Error != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除分类
func DeleteCategory(id int) int {
	err := db.Where("id = ?", id).Delete(&Category{})
	if err.Error != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
