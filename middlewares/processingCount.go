package middlewares

import (
	"fmt"
	"todo/config"
	"todo/utils"

	"github.com/gin-gonic/gin"
)

var count = 0

func ProcessingCount(c *gin.Context) {
	count++
	if count >= config.ProcessingCountThreshold {
		utils.Logger(
			" SYSTEM ",
			"WARN",
			fmt.Sprintf("ProcessingCount more than threshold: %d|%d", count, config.ProcessingCountThreshold),
		)
	}
	c.Next()
	count--
}

func GetProcessingCount() int {
	return count
}
