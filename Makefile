build:
	go build -o bin/fizzbuzz cmd/api/main.go

testing: build
	go test -v ./test

run: testing
	go run cmd/api/main.go
