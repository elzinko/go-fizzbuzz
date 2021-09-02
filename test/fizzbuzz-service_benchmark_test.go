package test

import (
	"testing"

	"github.com/elzinko/go-fizzbuzz/internal/api/services"
)

func Benchmark_FizzBuzz_100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		services.FizzBuzz(3, 5, 15, "Fizz", "Buzz")
	}
}
