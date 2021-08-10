package apis

import (
	"github.com/gin-gonic/gin"
)

type UserApis struct {
	Api
}

type GetUsersReq struct {
	Name string `json:"name"`
}

type CreateUsersReq struct {
	Name string `json:"name"`
}

type UpdateUsersReq struct {
	Name string `json:"name"`
}

type DeleteUsersReq struct {
	Name string `json:"name"`
}

func (a *UserApis) GetUsers(c *gin.Context) {
	a.MakeContext(c)
	body := GetUsersReq{}
	if a.MakeBody(&body) != nil {
		return
	}
	a.Info(body)
	c.JSON(200, gin.H{
		"message": "Bye " + "CCC",
	})
}

func (a *UserApis) CreateUsers(c *gin.Context) {
	a.MakeContext(c)
	body := CreateUsersReq{}
	if a.MakeBody(&body) != nil {
		return
	}
	a.Info("CREATE")
	c.JSON(200, gin.H{
		"message": "Bye " + "CCC",
	})
}

func (a *UserApis) UpdateUsers(c *gin.Context) {
	a.MakeContext(c)
	body := UpdateUsersReq{}
	if a.MakeBody(&body) != nil {
		return
	}
	a.Info("UPDATE")
	c.JSON(200, gin.H{
		"message": "Bye " + "CCC",
	})
}

func (a *UserApis) DeleteUsers(c *gin.Context) {
	a.MakeContext(c)
	body := DeleteUsersReq{}
	if a.MakeBody(&body) != nil {
		return
	}
	a.Info("DELETE")
	c.JSON(200, gin.H{
		"message": "Bye " + "CCC",
	})
}
