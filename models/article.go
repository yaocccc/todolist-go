package models

import (
	"fmt"
	"time"
	"todo/types/pagination"

	"gorm.io/gorm"
)

type ArticleModel struct{}

type Article struct {
	Id            int    `json:"id"`
	Type          int    `json:"type"`
	Status        int    `json:"status"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	CreatedTime   int64  `json:"created_time"`
	UpdatedTime   int64  `json:"updated_time"`
	CompletedTime int64  `json:"completed_time"`
	DeletedTime   int64  `json:"deleted_time"`
	IsDeleted     int    `json:"is_deleted"`
}

type ArticleCondition struct {
	Ids         []int
	Types       []int
	Statuses    []int
	CreatedTime []int
	IsDeleteds  []int
}

func makeArticleQuery(query *gorm.DB, condition ArticleCondition, keyword string) *gorm.DB {
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

func (m ArticleModel) getArticles(condition ArticleCondition, keyword string, p *pagination.Pagination) (articles []Article, count int64) {
	query := getTxOrDb(nil)
	query = makeArticleQuery(query, condition, keyword)
	if p != nil {
		query = query.Limit(p.Limit).Offset(p.Offset).Order(fmt.Sprintf("%s %s", p.OrderBy, p.OrderDir.String()))
	}
	query.Find(&articles).Count(&count)
	return
}

func (m ArticleModel) createArticles(articles []*Article, tx *gorm.DB) error {
	if articles == nil || len(articles) == 0 {
		return nil
	}
	now := time.Now().Unix()
	for _, article := range articles {
		article.CreatedTime = now
	}
	query := getTxOrDb(tx)
	return query.Create(articles).Error
}

func (m ArticleModel) updateArticles(articles []*Article, tx *gorm.DB) error {
	if articles == nil || len(articles) == 0 {
		return nil
	}
	now := time.Now().Unix()
	for _, article := range articles {
		article.UpdatedTime = now
	}
	query := getTxOrDb(tx)
	return query.Save(articles).Error
}

func (m ArticleModel) deleteArticles(condition ArticleCondition, keyword string, tx *gorm.DB) error {
	query := getTxOrDb(tx)
	query = makeArticleQuery(query, condition, keyword)
	return query.Delete(Article{}).Error
}
