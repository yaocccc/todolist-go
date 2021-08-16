package routers

import (
	"todo/apis"

	"github.com/gin-gonic/gin"
)

var apimap = make(map[string]gin.HandlerFunc)

func Router(r *gin.Engine) {
	articleApis := apis.ArticleApis{}
	apimap["GetArticles"] = articleApis.GetArticles
	apimap["CreateArticles"] = articleApis.CreateArticles
	apimap["UpdateArticles"] = articleApis.UpdateArticles
	apimap["DeleteArticles"] = articleApis.DeleteArticles

	tagApis := apis.TagApis{}
	apimap["GetTags"] = tagApis.GetTags
	apimap["CreateTags"] = tagApis.CreateTags
	apimap["UpdateTags"] = tagApis.UpdateTags
	apimap["DeleteTags"] = tagApis.DeleteTags

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
