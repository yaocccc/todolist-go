package apis

import (
	"todo/utils"

	"github.com/gin-gonic/gin"
)

type Api struct {
	RequestUuid string
	Context     *gin.Context
}

func (a *Api) MakeContext(c *gin.Context) {
	uid, _ := c.Get("RequestUuid")
	a.Context = c
	a.RequestUuid = uid.(string)
}

func (a *Api) MakeBody(body interface{}) error {
	c := a.Context
	err := c.BindJSON(&body)
	if err != nil {
		a.Error(err)
		c.JSON(404, gin.H{
			"code":    406,
			"message": err.Error(),
		})
		return err
	}
	return nil
}

func (a *Api) Info(message ...interface{}) {
	utils.Logger(a.RequestUuid, "INFO", message...)
}

func (a *Api) Error(message ...interface{}) {
	utils.Logger(a.RequestUuid, "ERROR", message...)
}
