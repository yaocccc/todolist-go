package main

import (
	"log"
	"todo/middlewares"
	"todo/models"
	"todo/routers"

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
	log.Fatal(r.Run(":9527"))
}
