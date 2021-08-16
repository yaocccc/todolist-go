package models

import (
	"fmt"
	"todo/types/pagination"

	"gorm.io/gorm"
)

type TagModel struct{}

type Tag struct {
	Id          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"default=''"`
	Description string `json:"description" gorm:"default=''"`
}

type TagCondition struct {
	Ids []int
}

func makeTagQuery(query *gorm.DB, condition TagCondition, keyword string) *gorm.DB {
	if condition.Ids != nil {
		query = query.Where("id IN (?)", condition.Ids)
	}
	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
	}
	return query
}

func (m TagModel) getTags(condition TagCondition, keyword string, p *pagination.Pagination) (tags []Tag, count int64) {
	query := getTxOrDb(nil)
	query = makeTagQuery(query, condition, keyword)
	if p != nil {
		query = query.Limit(p.Limit).Offset(p.Offset).Order(fmt.Sprintf("%s %s", p.OrderBy, p.OrderDir.String()))
	}
	query.Find(&tags).Count(&count)
	return
}

func (m TagModel) createTags(tags []*Tag, tx *gorm.DB) error {
	if tags == nil || len(tags) == 0 {
		return nil
	}
	query := getTxOrDb(tx)
	return query.Create(tags).Error
}

func (m TagModel) updateTags(tags []Tag, tx *gorm.DB) error {
	if tags == nil || len(tags) == 0 {
		return nil
	}
	query := getTxOrDb(tx)
	return query.Save(tags).Error
}

func (m TagModel) deleteTags(condition TagCondition, keyword string, tx *gorm.DB) error {
	query := getTxOrDb(tx)
	query = makeTagQuery(query, condition, keyword)
	return query.Delete(Tag{}).Error
}
