package main

import (
	"fmt"
)

func main() {
	numbers := buildRandomNumbers(12310000)

	amount := syncProcess(numbers)
	fmt.Printf("Even: %v Odd: %v\n", amount.even, amount.odd)

	amount = asyncProcess(numbers)
	fmt.Printf("Even: %v Odd: %v\n", amount.even, amount.odd)
}
