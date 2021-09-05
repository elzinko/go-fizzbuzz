build:
	go build -o bin/fizzbuzz cmd/fizzbuzz/main.go

testing: build
	go test -v ./test

swagger: build
	swag init -d cmd/fizzbuzz/ --parseDependency

run: testing
	FIZZBUZZ_BASE=/workspace/go-fizzbuzz/ CONFIG_PATH=data/config.yml go run cmd/fizzbuzz/main.go

benchmark: build
	go test -v ./test -bench=Benchmark_FizzBuzz_100 -cpuprofile=log/cpu.out

build-docker: build
	docker build . -t fizzbuzz

run-docker: build-docker
	docker run --env FIZZBUZZ_BASE=./ --env CONFIG_PATH=data/config.yml -p 3333:3333 fizzbuzz
