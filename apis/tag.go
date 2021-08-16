package apis

import (
	"todo/models"

	"github.com/gin-gonic/gin"
)

type TagApis struct {
	Api
}

type TagCondition struct {
	Ids []int `json:"ids"`
}

type TagCreation struct {
	Name        *string `json:"name" binding:"required"`
	Description *string `json:"description" binding:"required"`
}

type TagUpdation struct {
	Id          int     `json:"id" binding:"required"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type GetTagsReq struct {
	Condition TagCondition `json:"condition" binding:"required"`
	Keyword   *string      `json:"keyword" binding:"required"`
}

type CreateTagReq struct {
	Creations []TagCreation `json:"creations" binding:"required"`
}

type UpdateTagReq struct {
	Updations []TagUpdation `json:"updations" binding:"required"`
}

type DeleteTagsReq struct {
	Condition TagCondition `json:"condition" binding:"required"`
	Keyword   *string      `json:"keyword" binding:"required"`
}

func (a *TagApis) GetTags(c *gin.Context) {
	a.MakeContext(c)
	body := GetTagsReq{}
	if a.MakeBody(&body) != nil {
		return
	}

	condition := models.TagCondition{}
	tags, count := models.GetTags(condition, *body.Keyword, nil, nil)

	a.ResJson(gin.H{
		"code":    0,
		"message": "success",
		"data":    tags,
		"count":   count,
	}, nil)
}

func (a *TagApis) CreateTags(c *gin.Context) {
	a.MakeContext(c)
	body := CreateTagReq{}
	if a.MakeBody(&body) != nil {
		return
	}

	tags := []*models.Tag{}
	for _, creation := range body.Creations {
		tag := models.Tag{
			Name:        *creation.Name,
			Description: *creation.Description,
		}
		tags = append(tags, &tag)
	}
	err := models.CreateTags(tags)

	a.ResJson(gin.H{
		"code":    0,
		"message": "success",
		"data":    tags,
	}, err)
}

func (a *TagApis) UpdateTags(c *gin.Context) {
	a.MakeContext(c)
	body := UpdateTagReq{}
	if a.MakeBody(&body) != nil {
		return
	}

	updations := []models.Tag{}

	tagIds := []int{}
	for _, updation := range body.Updations {
		tagIds = append(tagIds, updation.Id)
	}

	tags, _ := models.GetTags(models.TagCondition{Ids: tagIds}, "", nil, nil)
	tagById := make(map[int]models.Tag)
	for _, tag := range tags {
		tagById[tag.Id] = tag
	}
	for _, updation := range body.Updations {
		tag, ok := tagById[updation.Id]
		if ok {
			if updation.Name != nil {
				tag.Name = *updation.Name
			}
			if updation.Description != nil {
				tag.Description = *updation.Description
			}
		}
		updations = append(updations, tag)
	}

	err := models.UpdateTags(updations)

	a.ResJson(gin.H{
		"code":    0,
		"message": "success",
	}, err)
}

func (a *TagApis) DeleteTags(c *gin.Context) {
	a.MakeContext(c)
	body := DeleteTagsReq{}
	if a.MakeBody(&body) != nil {
		return
	}
	condition := models.TagCondition{
		Ids: body.Condition.Ids,
	}

	err := models.DeleteTags(condition, *body.Keyword)

	a.ResJson(gin.H{
		"code":    0,
		"message": "success",
	}, err)
}
