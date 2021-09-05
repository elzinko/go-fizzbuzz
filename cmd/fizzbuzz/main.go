package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/elzinko/go-fizzbuzz/docs"
	"github.com/elzinko/go-fizzbuzz/internal/api"
)

// @title FizzBuzz generator using REST
// @version 0.6.0-SNAPSHOT
// @description FizzFuzz REST API in Golang with Gin Framework

// @contact.name Thomas Couderc
// @contact.email thomas.couderc@gmail.com

// @license.name MIT
// @license.url https://github.com/elzinko/go-fizzbuzz/blob/master/LICENSE

// @termsOfService http://swagger.io/terms/

// @BasePath /api/v0.6.0-SNAPSHOT

func main() {
	basePath, ok := os.LookupEnv("FIZZBUZZ_BASE")
	if !ok {
		log.Fatal("FIZZBUZZ_BASE should be set in environment variables")
	}

	configPath, ok := os.LookupEnv("CONFIG_PATH")
	if !ok {
		log.Fatal("CONFIG_PATH should be set in environment variables")
	}

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Test"
	// docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "petstore.swagger.io"
	// docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	api.Run(basePath, configPath)
}
