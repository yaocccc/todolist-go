package apis

import (
	"fmt"
	"todo/models"

	"github.com/gin-gonic/gin"
)

type ArticleApis struct {
	Api
}

type GetArticlesReq struct {
	Ids         []int  `json:"ids"`
	Types       []int  `json:"types"`
	Statuses    []int  `json:"statuses"`
	TagIds      []int  `json:"tag_ids"`
	CreatedTime []int  `json:"created_time"`
	IsDeleteds  []int  `json:"is_deleteds"`
	Keyword     string `json:"keyword"`
}

type ResArticle struct {
	models.Article
	Tags []models.Tag `json:"tags"`
}

type CreateArticleReq struct {
	Type    int    `json:"type" binding:"required"`
	Status  int    `json:"status"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	TagIds  []int  `json:"tag_ids" binding:"required"`
}

type UpdateArticleReq struct {
	Id        int    `json:"id" binding:"required"`
	Status    int    `json:"status"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	TagIds    []int  `json:"tag_ids"`
	IsDeleted int    `json:"is_deleted"`
}

type DeleteArticlesReq struct {
	Ids []int `json:"ids" binding:"required"`
}

func (a *ArticleApis) GetArticles(c *gin.Context) {
	a.MakeContext(c)
	body := GetArticlesReq{}
	if a.MakeBody(&body) != nil {
		return
	}

	/** GET ARTICLES */
	articleCondition := models.ArticleCondition{
		Ids:         body.Ids,
		Types:       body.Types,
		Statuses:    body.Statuses,
		CreatedTime: body.CreatedTime,
		IsDeleteds:  body.IsDeleteds,
	}
	articles, count := models.GetArticles(articleCondition, body.Keyword, nil, body.TagIds)

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

	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
		"data":    data,
		"count":   count,
	})
}

func (a *ArticleApis) CreateArticle(c *gin.Context) {
	a.MakeContext(c)
	body := CreateArticleReq{}
	if a.MakeBody(&body) != nil {
		return
	}

	/** CREATE ARTICLE */
	article := models.Article{
		Type:    body.Type,
		Status:  body.Status,
		Title:   body.Title,
		Content: body.Content,
	}
	models.CreateArticle(&article, body.TagIds)
	tags, _ := models.GetTags(models.TagCondition{}, "", nil, []int{article.Id})

	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
		"data": ResArticle{
			Article: article,
			Tags:    tags,
		},
		"count": 1,
	})
}

func (a *ArticleApis) UpdateArticle(c *gin.Context) {
	a.MakeContext(c)
	body := UpdateArticleReq{}
	if a.MakeBody(&body) != nil {
		return
	}
	article := models.Article{}
	article.Id = body.Id
	article.Title = body.Title
	article.Content = body.Content
	article.Status = body.Status
	article.IsDeleted = body.IsDeleted
	fmt.Printf("[logger-body.IsDeleted ]: %+v %+v\n", body.IsDeleted, body.IsDeleted == 0)

	if err := models.UpdateArticle(article, body.TagIds); err == nil {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "success",
		})
	} else {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err,
		})
	}
}

func (a *ArticleApis) DeleteArticles(c *gin.Context) {
	a.MakeContext(c)
	body := DeleteArticlesReq{}
	if a.MakeBody(&body) != nil {
		return
	}
	articleCondition := models.ArticleCondition{}
	articleCondition.Ids = body.Ids

	models.DeleteArticles(articleCondition, "")

	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
	})
}
