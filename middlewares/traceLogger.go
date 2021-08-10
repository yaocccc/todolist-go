package middlewares

import (
	"time"
	"todo/utils"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func TraceLogger(c *gin.Context) {
	_uid, _ := uuid.NewV4()
	uid := _uid.String()[:8]
	c.Set("RequestUuid", uid)

	method, path, clientIP := c.Request.Method, c.Request.URL, c.ClientIP()
	utils.Logger(uid, "INFO", method, path, clientIP, "START")

	start := time.Now()
	c.Next()
	end := time.Now()

	statusCode := c.Writer.Status()
	latency := end.Sub(start)
	utils.Logger(uid, "INFO", method, path, clientIP, "END -- ", statusCode, "|", latency)
}
