package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AuthCheck(c *gin.Context) {
	c.Next()
}
