package apis

type TagApis struct {
	Api
}

type GetTagsReq struct {
	Keyword string `json:"keyword"`
}

type CreateTagReq struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateTagReq struct {
	Id          int    `json:"id" binding:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DeleteTagsReq struct {
	Ids []int `json:"ids" binding:"required"`
}

// func (a *TagApis) GetTags(c *gin.Context) {
// 	a.MakeContext(c)
// 	body := GetTagsReq{}
// 	if a.MakeBody(&body) != nil {
// 		return
// 	}

// 	condition := models.TagCondition{}
// 	tags, count := models.GetTags(condition, body.Keyword, nil)

// 	c.JSON(200, gin.H{
// 		"code":    0,
// 		"message": "success",
// 		"data":    tags,
// 		"count":   count,
// 	})
// }

// func (a *TagApis) CreateTag(c *gin.Context) {
// 	a.MakeContext(c)
// 	body := CreateTagReq{}
// 	if a.MakeBody(&body) != nil {
// 		return
// 	}

// 	tag := models.Tag{
// 		Name:        body.Name,
// 		Description: body.Description,
// 	}
// 	models.CreateTags([]*models.Tag{&tag})

// 	c.JSON(200, gin.H{
// 		"code":    0,
// 		"message": "success",
// 		"data":    tag,
// 		"count":   1,
// 	})
// }

// func (a *TagApis) UpdateTag(c *gin.Context) {
// 	a.MakeContext(c)
// 	body := UpdateTagReq{}
// 	if a.MakeBody(&body) != nil {
// 		return
// 	}
// 	condition := models.TagCondition{Ids: []int{body.Id}}
// 	updation := models.Tag{
// 		Name:        body.Name,
// 		Description: body.Description,
// 	}
// 	models.UpdateTags(condition, "", updation)
// 	c.JSON(200, gin.H{
// 		"code":    0,
// 		"message": "success",
// 	})
// }

// func (a *TagApis) DeleteTag(c *gin.Context) {
// 	a.MakeContext(c)
// 	body := DeleteTagsReq{}
// 	if a.MakeBody(&body) != nil {
// 		return
// 	}
// 	tagCondition := models.TagCondition{Ids: body.Ids}
// 	refConditon := models.ArticleTagRefCondition{TagIds: body.Ids}

// 	models.DeleteTags(tagCondition, "")
// 	models.DeleteArticleTagRefs(refConditon)

// 	c.JSON(200, gin.H{
// 		"code":    0,
// 		"message": "success",
// 	})
// }
