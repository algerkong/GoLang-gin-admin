package model

import (
	"gin-gorm/utils/errmsg"

	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title      string   `gorm:"type:varchar(100);not null" json:"title"`
	Content    string   `gorm:"type:longtext;not null" json:"content"`
	Desc       string   `gorm:"type:varchar(200);not null" json:"desc"`
	Img        string   `gorm:"type:varchar(100)" json:"img"`
	Author     string   `gorm:"type:varchar(20);not null" json:"author"`
	CategoryId uint     `gorm:"type:int;not null" json:"category_id"`
	Category   Category `gorm:"foreignkey:CategoryId" json:"category"`
}

// 添加文章
func AddArticle(article *Article) (code int) {
	err := db.Create(&article)
	if err.Error != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询文章是否存在
func CheckArticleID(id uint) (code int) {
	var article Article
	db.Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return errmsg.SUCCESS
	}
	return errmsg.ERROR_ARTICLE_NOT_EXIST
}

// 查询单个文章
func GetArtilce(id uint) (Article, int) {

	var article Article
	code := CheckArticleID(id)
	if code != errmsg.SUCCESS {
		return article, code
	}
	err := db.Preload("Category").Where("id = ?", id).First(&article)
	if err.Error != nil {
		return article, errmsg.ERROR
	}
	return article, errmsg.SUCCESS
}

// 查询单个分类下的文章
func GetCategoryArticles(id uint, pageSize, pageNum int) ([]Article, int, int) {
	var articleList []Article
	var total int
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("category_id=?", id).Find(&articleList).Count(&total)
	if err.Error != nil {
		return nil, 0, errmsg.ERROR
	}
	return articleList, total, errmsg.SUCCESS
}

// 查询文章列表
func GetArtilces(pageSize, pageNum int) ([]Article, int, int) {
	var articleList []Article
	var total int
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Count(&total)
	if err.Error != nil && err.Error != gorm.ErrRecordNotFound {
		return nil, total, errmsg.ERROR
	}
	return articleList, total, errmsg.SUCCESS
}

// 编辑文章
func EditArticle(id uint, article *Article) (code int) {
	var maps = make(map[string]interface{})
	maps["title"] = article.Title
	maps["content"] = article.Content
	maps["desc"] = article.Desc
	maps["img"] = article.Img
	maps["author"] = article.Author
	maps["category_id"] = article.CategoryId
	err := db.Model(&Article{}).Where("id = ?", id).Updates(maps)
	if err.Error != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除文章
func DeleteArticle(id uint) (code int) {
	err := db.Where("id = ?", id).Delete(&Article{})
	if err.Error != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
