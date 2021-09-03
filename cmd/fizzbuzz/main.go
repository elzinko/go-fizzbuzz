package main

import (
	"os"

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
	api.Run(os.Getenv("FIZZBUZZ_BASE") + "/data/config.yml")
}
