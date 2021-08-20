// @title TODO LIST API
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

package main

import (
	"log"
	"todo/middlewares"
	"todo/models"
	"todo/routers"

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "todo/docs"

	"github.com/gin-gonic/gin"
)

func init() {
	models.Setup()
}

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middlewares.TraceLogger)
	r.Use(middlewares.AuthCheck)
	r.Use(middlewares.ProcessingCount)

	routers.Router(r)
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	log.Fatal(r.Run(":9527"))
}
