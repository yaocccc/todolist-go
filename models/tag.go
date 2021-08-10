package models

import (
	"fmt"
	"todo/types/pagination"

	"gorm.io/gorm"
)

type Tag struct {
	Id          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"default=''"`
	Description string `json:"description" gorm:"default=''"`
}

type TagCondition struct {
	Ids []int
}

func makeTagQuery(condition TagCondition, keyword string) *gorm.DB {
	query := db
	if condition.Ids != nil {
		query = query.Where("id IN (?)", condition.Ids)
	}
	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
	}
	return query
}

func GetTags(condition TagCondition, keyword string, p *pagination.Pagination) (tags []Tag, count int64) {
	query := makeTagQuery(condition, keyword)
	if p != nil {
		query = query.Limit(p.Limit).Offset(p.Offset).Order(fmt.Sprintf("%s %s", p.OrderBy, p.OrderDir.String()))
	}
	query.Find(&tags).Count(&count)
	return
}

func CreateTags(tags []*Tag) {
	query := db
	query.Create(tags)
}

func UpdateTags(condition TagCondition, keyword string, updation Tag) {
	query := makeTagQuery(condition, keyword)
	query.Updates(updation)
}

func UpdateTagsByEntities(tags []*Tag) {
	query := db
	query.Save(tags)
}

func DeleteTags(condition TagCondition, keyword string) {
	query := makeTagQuery(condition, keyword)
	query.Delete(Tag{})
}
