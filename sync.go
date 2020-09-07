package main

func syncProcess(numbers []int) amountNumber {
	amount := amountNumber{}
	for _, number := range numbers {
		if number%2 == 0 {
			amount.even++
			continue
		}
		amount.odd++
	}
	return amount
}
