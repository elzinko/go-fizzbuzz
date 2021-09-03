package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/elzinko/go-fizzbuzz/internal/api"
)

// @Golang FizzBuzz REST API
// @version 0.5.1
// @description FizzBuzz REST API in Golang with Gin Framework

// @contact.name Thomas Couderc
// @contact.email thomas.couderc@gmail.com

// @license.name MIT
// @license.url https://github.com/elzinko/go-fizzbuzz/blob/master/LICENSE

// @BasePath /

func main() {
	basePath, ok := os.LookupEnv("FIZZBUZZ_BASE")
	if !ok {
		log.Fatal("FIZZBUZZ_BASE should be set in environment variables")
	}

	configPath, ok := os.LookupEnv("CONFIG_PATH")
	if !ok {
		log.Fatal("CONFIG_PATH should be set in environment variables")
	}

	api.Run(basePath, configPath)
}
