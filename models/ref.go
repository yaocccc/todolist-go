package models

import (
	"gorm.io/gorm"
)

type ArticleTagRef struct {
	Id        int `json:"id" gorm:"primaryKey;autoIncrement"`
	ArticleId int `json:"article_id" gorm:"default=0"`
	TagId     int `json:"tag_id" gorm:"default=0"`
}

type ArticleTagRefCondition struct {
	Ids        []int
	ArticleIds []int
	TagIds     []int
}

func makeArticleTagRefQuery(condition ArticleTagRefCondition) *gorm.DB {
	query := db
	if condition.Ids != nil {
		query = query.Where("id IN (?)", condition.Ids)
	}
	if condition.ArticleIds != nil {
		query = query.Where("article_id IN (?)", condition.ArticleIds)
	}
	if condition.TagIds != nil {
		query = query.Where("tag_id IN (?)", condition.TagIds)
	}
	return query
}

func GetArticleTagRefs(condition ArticleTagRefCondition) (article_tag_refs []ArticleTagRef, count int64) {
	query := makeArticleTagRefQuery(condition)
	query.Find(&article_tag_refs).Count(&count)
	return
}

func CreateArticleTagRefs(article_tag_refs []*ArticleTagRef) {
	query := db
	query.Create(article_tag_refs)
}

func DeleteArticleTagRefs(condition ArticleTagRefCondition) {
	query := makeArticleTagRefQuery(condition)
	query.Delete(ArticleTagRef{})
}
