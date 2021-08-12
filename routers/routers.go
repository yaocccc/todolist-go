package routers

import (
	"todo/apis"

	"github.com/gin-gonic/gin"
)

var apimap = make(map[string]gin.HandlerFunc)

func Router(r *gin.Engine) {
	articleApis := apis.ArticleApis{}
	apimap["GetArticles"] = articleApis.GetArticles
	apimap["CreateArticle"] = articleApis.CreateArticle
	apimap["UpdateArticle"] = articleApis.UpdateArticle
	apimap["DeleteArticles"] = articleApis.DeleteArticles

	tagApis := apis.TagApis{}
	apimap["GetTags"] = tagApis.GetTags
	apimap["CreateTag"] = tagApis.CreateTag
	apimap["UpdateTag"] = tagApis.UpdateTag
	apimap["DeleteTag"] = tagApis.DeleteTag

	r.POST("", func(c *gin.Context) {
		action := c.Query("Action")
		api, ok := apimap[action]
		if !ok {
			c.JSON(404, gin.H{
				"code":    404,
				"message": "Can't find action " + action,
			})
			return
		}
		api(c)
	})
}
