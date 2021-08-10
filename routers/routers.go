package routers

import (
	"todo/apis"

	"github.com/gin-gonic/gin"
)

var apimap = make(map[string]gin.HandlerFunc)

func Router(r *gin.Engine) {
	userapis := apis.UserApis{}
	apimap["GetUsers"] = userapis.GetUsers
	apimap["CreateUsers"] = userapis.CreateUsers
	apimap["UpdateUsers"] = userapis.UpdateUsers
	apimap["DeleteUsers"] = userapis.DeleteUsers

	r.POST("", func(c *gin.Context) {
		action := c.Query("Action")
		api, ok := apimap[action]
		if !ok {
			c.JSON(404, gin.H{
				"Code":    404,
				"Message": "Can't find action " + action,
			})
			return
		}
		api(c)
	})
}
