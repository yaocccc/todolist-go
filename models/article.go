package models

import (
	"fmt"
	"todo/types/pagination"

	"gorm.io/gorm"
)

type Article struct {
	Id            int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Type          int    `json:"type" gorm:"default=1"`
	Status        int    `json:"status" gorm:"default=0"`
	Title         string `json:"title" gorm:"default=''"`
	Content       string `json:"content" gorm:"default=''"`
	CreatedTime   int64  `json:"created_time" gorm:"default=0"`
	UpdatedTime   int64  `json:"updated_time" gorm:"default=0"`
	CompletedTime int64  `json:"completed_time" gorm:"default=0"`
	DeletedTime   int64  `json:"deleted_time" gorm:"default=0"`
	IsDeleted     int    `json:"is_deleted" gorm:"default=0"`
}

type ArticleCondition struct {
	Ids         []int
	Types       []int
	Statuses    []int
	CreatedTime []int
	IsDeleteds  []int
}

func makeArticleQuery(condition ArticleCondition, keyword string) *gorm.DB {
	query := db
	if condition.Ids != nil {
		query = query.Where("id IN (?)", condition.Ids)
	}
	if condition.Types != nil {
		query = query.Where("type IN (?)", condition.Types)
	}
	if condition.Statuses != nil {
		query = query.Where("status IN (?)", condition.Statuses)
	}
	if condition.CreatedTime != nil {
		query = query.Where("created_time > (?) and created_time <= (?)", condition.CreatedTime[0], condition.CreatedTime[1])
	}
	if condition.IsDeleteds != nil {
		query = query.Where("is_deleted IN (?)", condition.IsDeleteds)
	} else {
		query = query.Where("is_deleted = 0")
	}
	if keyword != "" {
		query = query.Where("title LIKE ?", "%"+keyword+"%")
	}
	return query
}

func GetArticles(condition ArticleCondition, keyword string, p *pagination.Pagination) (articles []Article, count int64) {
	query := makeArticleQuery(condition, keyword)
	if p != nil {
		query = query.Limit(p.Limit).Offset(p.Offset).Order(fmt.Sprintf("%s %s", p.OrderBy, p.OrderDir.String()))
	}
	query.Find(&articles).Count(&count)
	return
}

func CreateArticles(articles []*Article) {
	query := db
	query.Create(articles)
}

func UpdateArticles(condition ArticleCondition, keyword string, updation Article) {
	query := makeArticleQuery(condition, keyword)
	query.Updates(updation)
}

func UpdateArticlesByEntities(articles []*Article) {
	query := db
	query.Save(articles)
}

func DeleteArticles(condition ArticleCondition, keyword string, hardDel bool) {
	query := makeArticleQuery(condition, keyword)
	if hardDel {
		query.Updates(Article{IsDeleted: 1})
	} else {
		query.Delete(Article{})
	}
}
