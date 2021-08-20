package apis

import (
	"net/http"
	"todo/utils"

	"github.com/gin-gonic/gin"
)

type Api struct {
	RequestUuid string
	Context     *gin.Context
}

type CommonRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
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
		c.JSON(400, gin.H{
			"code":    400,
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

func (a *Api) ResJson(obj interface{}, err error) {
	if err == nil {
		a.Context.JSON(http.StatusOK, obj)
	} else {
		a.Context.JSON(500, gin.H{
			"code":    500,
			"message": err,
		})
	}
}
