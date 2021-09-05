package router

import (
	"github.com/elzinko/go-fizzbuzz/internal/api/controllers"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func Setup(basePath string) *gin.Engine {
	app := gin.New()
	app.Use(gin.Recovery())
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.GET("/api/fizzbuzz", controllers.GetFizzFuzz)
	return app
}
