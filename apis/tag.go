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

type GetTagsRes struct {
	CommonRes
	Data  []models.Tag `json:"data"`
	Count int          `json:"count"`
}

// @Tags TAG
// @Summary Get tags by condition
// @Description Get tags
// @Body {object} GetTagsReq
// @Success 200 {object} GetTagsRes
// @Router /GetTags [post]
func (a *TagApis) GetTags(c *gin.Context) {
	a.MakeContext(c)
	body := GetTagsReq{}
	if a.MakeBody(&body) != nil {
		return
	}

	condition := models.TagCondition{Ids: body.Condition.Ids}
	tags, count := models.GetTags(condition, *body.Keyword, nil, nil)

	a.ResJson(gin.H{
		"code":    0,
		"message": "success",
		"data":    tags,
		"count":   count,
	}, nil)
}

type CreateTagsReq struct {
	Creations []TagCreation `json:"creations" binding:"required"`
}

type CreateTagsRes struct {
	CommonRes
	Data []int `json:"data"`
}

// @Tags TAG
// @Summary Create tags by creations
// @Description Create tags
// @Body {object} CreateTagsReq
// @Success 200 {object} CreateTagsRes
// @Router /CreateTags [post]
func (a *TagApis) CreateTags(c *gin.Context) {
	a.MakeContext(c)
	body := CreateTagsReq{}
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
	if err := models.CreateTags(tags); err == nil {
		tagIds := []int{}
		for _, tag := range tags {
			tagIds = append(tagIds, tag.Id)
		}
		a.ResJson(gin.H{
			"code":    0,
			"message": "success",
			"data":    tagIds,
		}, nil)
	} else {
		a.ResJson(nil, err)
	}
}

type UpdateTagsReq struct {
	Updations []TagUpdation `json:"updations" binding:"required"`
}

type UpdateTagsRes struct {
	CommonRes
}

// @Tags TAG
// @Summary Update tags by updations
// @Description Update tags
// @Body {object} UpdateTagsReq
// @Success 200 {object} UpdateTagsRes
// @Router /UpdateTags [post]
func (a *TagApis) UpdateTags(c *gin.Context) {
	a.MakeContext(c)
	body := UpdateTagsReq{}
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

type DeleteTagsReq struct {
	Condition TagCondition `json:"condition" binding:"required"`
	Keyword   *string      `json:"keyword" binding:"required"`
}

type DeleteTagsRes struct {
	CommonRes
}

// @Tags TAG
// @Summary Delete tags by condition
// @Description Delete tags
// @Body {object} DeleteTagsReq
// @Success 200 {object} DeleteTagsRes
// @Router /DeleteTags [post]
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
