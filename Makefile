get-swag: 
	go get -u github.com/swaggo/swag/cmd/swag

swagger: get-swag
	swag init -d cmd/fizzbuzz/ --parseDependency

build: swagger
	go build -o bin/fizzbuzz cmd/fizzbuzz/main.go

testing: build
	go test -v ./test/...

benchmark: build
	go test -v ./test/... -bench=Benchmark_FizzBuzz_100 -cpuprofile=log/cpu.out

run: testing
	FIZZBUZZ_BASE=/workspace/go-fizzbuzz/ CONFIG_PATH=data/config.yml go run cmd/fizzbuzz/main.go

build-docker: build
	docker build . -t fizzbuzz

run-docker: build-docker
	docker run --env FIZZBUZZ_BASE=./ --env CONFIG_PATH=data/config.yml -p 3333:3333 fizzbuzz
