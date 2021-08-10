package middlewares

import (
	"bytes"
	"io/ioutil"
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
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	utils.Logger(uid, "INFO", method, path, clientIP, "START")
	utils.Logger(uid, "INFO", "REQ:", string(bodyBytes))

	start := time.Now()
	c.Next()
	end := time.Now()

	statusCode := c.Writer.Status()
	latency := end.Sub(start)
	utils.Logger(uid, "INFO", method, path, clientIP, "END -- ", statusCode, "|", latency)
}
