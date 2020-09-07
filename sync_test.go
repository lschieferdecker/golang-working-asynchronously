package main

import (
	"testing"
)

var syncAmount amountNumber

func benchmarkSyncEvenOrOdd(i int, b *testing.B) {
	numbers := buildRandomNumbers(i)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		resSync := syncProcess(numbers)
		syncAmount = resSync
	}
}

func BenchmarkSyncEvenOrOdd10000(b *testing.B)   { benchmarkSyncEvenOrOdd(10000, b) }
func BenchmarkSyncEvenOrOdd100000(b *testing.B)  { benchmarkSyncEvenOrOdd(100000, b) }
func BenchmarkSyncEvenOrOdd500000(b *testing.B)  { benchmarkSyncEvenOrOdd(500000, b) }
func BenchmarkSyncEvenOrOdd1000000(b *testing.B) { benchmarkSyncEvenOrOdd(1000000, b) }
func BenchmarkSyncEvenOrOdd5000000(b *testing.B) { benchmarkSyncEvenOrOdd(5000000, b) }
