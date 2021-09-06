package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/elzinko/go-fizzbuzz/internal/api"
)

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
