build:
	go build -o bin/fizzbuzz cmd/api/main.go

testing: build
	go test -v ./test

run: testing
	go run cmd/api/main.go

benchmark: build
	go test -v ./test -bench=Benchmark_FizzBuzz_100 -cpuprofile=log/cpu.out

build-docker: build
	docker build . -t fizzbuzz

run-docker: build-docker
	docker run -p 3333:3333 fizzbuzz
