get-swag: 
	go install github.com/swaggo/swag/cmd/swag@v1.7.1

swagger: get-swag
	swag init -d cmd/fizzbuzz/ --parseDependency

build: swagger
	go build -o bin/fizzbuzz cmd/fizzbuzz/main.go

testing:
	go test -v ./test/...

benchmark: build
	go test -v ./test/... -bench=Benchmark_FizzBuzz_100 -cpuprofile=log/cpu.out

run:

	FIZZBUZZ_BASE=$(CURDIR)/ CONFIG_PATH=data/config.yml go run cmd/fizzbuzz/main.go

build-docker: build
	docker build . -t fizzbuzz

run-docker: build-docker
	docker run --env FIZZBUZZ_BASE=./ --env CONFIG_PATH=data/config.yml -p 3333:3333 fizzbuzz
