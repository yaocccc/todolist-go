package models

import (
	"gorm.io/gorm"
)

type ArticleTagRefModel struct{}

type ArticleTagRef struct {
	Id        int `json:"id"`
	ArticleId int `json:"article_id"`
	TagId     int `json:"tag_id"`
}

type ArticleTagRefCondition struct {
	Ids        []int
	ArticleIds []int
	TagIds     []int
}

func makeArticleTagRefQuery(query *gorm.DB, condition ArticleTagRefCondition) *gorm.DB {
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

func (m ArticleTagRefModel) getArticleTagRefs(condition ArticleTagRefCondition) (article_tag_refs []ArticleTagRef) {
	query := getTxOrDb(nil)
	query = makeArticleTagRefQuery(query, condition)
	query.Find(&article_tag_refs)
	return
}

func (m ArticleTagRefModel) createArticleTagRefs(article_tag_refs []*ArticleTagRef, tx *gorm.DB) error {
	if article_tag_refs == nil || len(article_tag_refs) == 0 {
		return nil
	}
	query := getTxOrDb(tx)
	return query.Create(article_tag_refs).Error
}

func (m ArticleTagRefModel) deleteArticleTagRefs(condition ArticleTagRefCondition, tx *gorm.DB) error {
	query := getTxOrDb(tx)
	query = makeArticleTagRefQuery(query, condition)
	return query.Delete(ArticleTagRef{}).Error
}
