package router

import (
	"github.com/elzinko/go-fizzbuzz/internal/api/controllers"
	"github.com/gin-gonic/gin"
)

func Setup(basePath string) *gin.Engine {
	app := gin.New()
	app.Use(gin.Recovery())
	app.GET("/api/fizzbuzz", controllers.GetFizzFuzz)
	return app
}
