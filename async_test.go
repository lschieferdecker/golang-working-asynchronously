package main

import (
	"testing"
)

var asyncAmount amountNumber

func benchmarkAsyncEvenOrOdd(i int, b *testing.B) {
	numbers := buildRandomNumbers(i)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		resAsync := asyncProcess(numbers)
		asyncAmount = resAsync
	}
}

func BenchmarkAsyncEvenOrOdd10000(b *testing.B)   { benchmarkAsyncEvenOrOdd(10000, b) }
func BenchmarkAsyncEvenOrOdd100000(b *testing.B)  { benchmarkAsyncEvenOrOdd(100000, b) }
func BenchmarkAsyncEvenOrOdd500000(b *testing.B)  { benchmarkAsyncEvenOrOdd(500000, b) }
func BenchmarkAsyncEvenOrOdd1000000(b *testing.B) { benchmarkAsyncEvenOrOdd(1000000, b) }
func BenchmarkAsyncEvenOrOdd5000000(b *testing.B) { benchmarkAsyncEvenOrOdd(5000000, b) }
