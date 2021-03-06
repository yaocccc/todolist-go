package apis

import (
	"time"
	"todo/models"

	"github.com/gin-gonic/gin"
)

type ArticleApis struct {
	Api
}

type ArticleCondition struct {
	Ids         []int `json:"ids"`
	Types       []int `json:"types"`
	Statuses    []int `json:"statuses"`
	TagIds      []int `json:"tag_ids"`
	CreatedTime []int `json:"created_time"`
	IsDeleteds  []int `json:"is_deleteds"`
}

type ArticleCreation struct {
	Type    *int    `json:"type" binding:"required"`
	Status  *int    `json:"status" binding:"required"`
	Title   *string `json:"title" binding:"required"`
	Content *string `json:"content" binding:"required"`
	TagIds  []int   `json:"tag_ids" binding:"required"`
}

type ArticleUpdation struct {
	Id        *int    `json:"id" binding:"required"`
	Status    *int    `json:"status"`
	Title     *string `json:"title"`
	Content   *string `json:"content"`
	TagIds    []int   `json:"tag_ids"`
	IsDeleted *int    `json:"is_deleted"`
}

type ResArticle struct {
	models.Article
	Tags []models.Tag `json:"tags"`
}

type GetArticlesReq struct {
	Condition ArticleCondition `json:"condition" binding:"required"`
	Keyword   *string          `json:"keyword" binding:"required"`
}

type GetArticlesRes struct {
	CommonRes
	Data  []ResArticle `json:"data"`
	Count int          `json:"count"`
}

// @Tags ARTICLE
// @Summary Get articles by condition
// @Description Get articles
// @Body {object} GetArticlesReq
// @Success 200 {object} GetArticlesRes
// @Router /GetArticles [post]
func (a *ArticleApis) GetArticles(c *gin.Context) {
	a.MakeContext(c)
	body := GetArticlesReq{}
	if a.MakeBody(&body) != nil {
		return
	}

	/** GET ARTICLES */
	articleCondition := models.ArticleCondition{
		Ids:         body.Condition.Ids,
		Types:       body.Condition.Types,
		Statuses:    body.Condition.Statuses,
		CreatedTime: body.Condition.CreatedTime,
		IsDeleteds:  body.Condition.IsDeleteds,
	}
	articles, count := models.GetArticles(articleCondition, *body.Keyword, nil, body.Condition.TagIds)

	/** GET TAGS */
	articleIds := []int{}
	for _, article := range articles {
		articleIds = append(articleIds, article.Id)
	}
	tags, _ := models.GetTags(models.TagCondition{}, "", nil, articleIds)
	tagById := make(map[int]models.Tag)
	for _, tag := range tags {
		tagById[tag.Id] = tag
	}

	/** GROUP TAG_IDS BY ARTICLE_ID */
	refs := models.GetArticleTagRefs(models.ArticleTagRefCondition{ArticleIds: articleIds})
	tagIdsGroupByArticleId := make(map[int]([]int))
	for _, ref := range refs {
		if tagIds, ok := tagIdsGroupByArticleId[ref.ArticleId]; ok {
			tagIds = append(tagIds, ref.TagId)
			tagIdsGroupByArticleId[ref.ArticleId] = tagIds
		} else {
			tagIdsGroupByArticleId[ref.ArticleId] = []int{ref.TagId}
		}
	}

	/** MAP TO RES */
	data := []ResArticle{}
	for _, article := range articles {
		tempRes := ResArticle{
			Article: article,
			Tags:    []models.Tag{},
		}
		if tagIds, ok := tagIdsGroupByArticleId[article.Id]; ok {
			for _, tagId := range tagIds {
				if tag, ok := tagById[tagId]; ok {
					tempRes.Tags = append(tempRes.Tags, tag)
				}
			}
		}
		data = append(data, tempRes)
	}

	a.ResJson(gin.H{
		"code":    0,
		"message": "success",
		"data":    data,
		"count":   count,
	}, nil)
}

type CreateArticlesReq struct {
	Creations []ArticleCreation `json:"creations" binding:"required"`
}

type CreateArticlesRes struct {
	CommonRes
	Data []int `json:"data"`
}

// @Tags ARTICLE
// @Summary Create articles by creations
// @Description Create articles
// @Body {object} CreateArticlesReq
// @Success 200 {object} CreateArticlesRes
// @Router /CreateArticles [post]
func (a *ArticleApis) CreateArticles(c *gin.Context) {
	a.MakeContext(c)
	body := CreateArticlesReq{}
	if a.MakeBody(&body) != nil {
		return
	}

	/** CREATE ARTICLES */
	creations := []*models.ArticleCreation{}
	for _, creation := range body.Creations {
		article := models.Article{
			Type:    *creation.Type,
			Status:  *creation.Status,
			Title:   *creation.Title,
			Content: *creation.Content,
		}
		creations = append(creations, &models.ArticleCreation{
			Article: article,
			TagIds:  creation.TagIds,
		})
	}

	if err := models.CreateArticles(creations); err == nil {
		articleIds := []int{}
		for _, article := range creations {
			articleIds = append(articleIds, article.Id)
		}

		a.ResJson(gin.H{
			"code":    0,
			"message": "success",
			"data":    articleIds,
		}, nil)
	} else {
		a.ResJson(nil, err)
	}

}

type UpdateArticlesReq struct {
	Updations []ArticleUpdation `json:"updations" binding:"required"`
}

type UpdateArticlesRes struct {
	CommonRes
}

// @Tags ARTICLE
// @Summary Update articles by updations
// @Description Update articles
// @Body {object} UpdateArticlesReq
// @Success 200 {object} UpdateArticlesRes
// @Router /UpdateArticles [post]
func (a *ArticleApis) UpdateArticles(c *gin.Context) {
	a.MakeContext(c)
	body := UpdateArticlesReq{}
	if a.MakeBody(&body) != nil {
		return
	}

	updations := []models.ArticleUpdation{}

	articleIds := []int{}
	updationById := make(map[int]ArticleUpdation)
	for _, updation := range body.Updations {
		articleIds = append(articleIds, *updation.Id)
		updationById[*updation.Id] = updation
	}

	/** TODO: gorm??????save?????????????????????????????? ???????????????????????????????????? ???????????????????????????????????? */
	articles, _ := models.GetArticles(models.ArticleCondition{Ids: articleIds, IsDeleteds: []int{0, 1}}, "", nil, nil)
	now := time.Now().Unix()
	for _, article := range articles {
		updation, ok := updationById[article.Id]
		if ok {
			if updation.Title != nil {
				article.Title = *updation.Title
			}
			if updation.Content != nil {
				article.Content = *updation.Content
			}
			if updation.Status != nil {
				if *updation.IsDeleted == 1 && article.Status == 0 {
					article.CompletedTime = now
				}
				article.Status = *updation.Status
			}
			if updation.IsDeleted != nil {
				if *updation.IsDeleted == 1 && article.IsDeleted == 0 {
					article.DeletedTime = now
				}
				article.IsDeleted = *updation.IsDeleted
			}
			updations = append(updations, models.ArticleUpdation{
				Article: article,
				TagIds:  updation.TagIds,
			})
		}
	}

	err := models.UpdateArticles(updations)
	a.ResJson(gin.H{
		"code":    0,
		"message": "success",
	}, err)
}

type DeleteArticlesReq struct {
	Condition ArticleCondition `json:"condition" binding:"required"`
	Keyword   *string          `json:"keyword" binding:"required"`
}

type DeleteArticlesRes struct {
	CommonRes
}

// @Tags ARTICLE
// @Summary Delete articles by condition
// @Description Delete articles
// @Body {object} DeleteArticlesReq
// @Success 200 {object} DeleteArticlesRes
// @Router /DeleteArticles [post]
func (a *ArticleApis) DeleteArticles(c *gin.Context) {
	a.MakeContext(c)
	body := DeleteArticlesReq{}
	if a.MakeBody(&body) != nil {
		return
	}

	articleCondition := models.ArticleCondition{
		Ids:         body.Condition.Ids,
		Types:       body.Condition.Types,
		Statuses:    body.Condition.Statuses,
		CreatedTime: body.Condition.CreatedTime,
		IsDeleteds:  body.Condition.IsDeleteds,
	}

	err := models.DeleteArticles(articleCondition, *body.Keyword)

	a.ResJson(gin.H{
		"code":    0,
		"message": "success",
	}, err)
}
