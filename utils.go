package main

import (
	"math/rand"
	"time"
)

type amountNumber struct {
	even int
	odd  int
}

func buildRandomNumbers(quantity int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, quantity, quantity)
	for i := 0; i < quantity; i++ {
		numbers[i] = rand.Intn(quantity)
	}
	return numbers
}

func chunkNumbers(numbers []int, chunkSize int) [][]int {
	chunkedNumbers := [][]int{}
	for i := 0; i < len(numbers); i += chunkSize {
		chunkedNumbers = append(chunkedNumbers, numbers[i:min(i+chunkSize, len(numbers))])
	}
	return chunkedNumbers
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
