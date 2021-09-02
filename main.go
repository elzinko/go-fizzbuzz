package main

import (
	"github.com/elzinko/go-fizzbuzz/internal/api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api/fizzbuzz", controllers.GetFizzBuzz)
	r.Run()
}
