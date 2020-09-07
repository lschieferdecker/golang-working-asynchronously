package main

import (
	"sync"
)

func asyncProcess(numbers []int) amountNumber {
	amount := amountNumber{}
	wg := &sync.WaitGroup{}

	// Chunk the slice into chunks with 10000 elements each
	chunkedNumbers := chunkNumbers(numbers, 10000)

	// An buffered channel to increase performance
	ammountChannel := make(chan amountNumber, len(chunkedNumbers))

	// Iterate each chunk and proccess them async with
	// goroutines
	for _, chunk := range chunkedNumbers {
		wg.Add(1)
		go processChunk(chunk, ammountChannel, wg)
	}

	// Closes the channel when all data wait groups
	// are done
	go processMonitor(ammountChannel, wg)

	// Receives data from channel
	for receivedAmount := range ammountChannel {
		amount.even += receivedAmount.even
		amount.odd += receivedAmount.odd
	}

	return amount
}

func processMonitor(channel chan amountNumber, wg *sync.WaitGroup) {
	// Wait all groups to be done and then closes the channel
	// in order to avoid 'all goroutines are asleep - deadlock!'
	// error
	wg.Wait()
	close(channel)
}

func processChunk(chunk []int, channel chan amountNumber, wg *sync.WaitGroup) {
	// Closes this wait group when finish this function
	defer wg.Done()

	amount := amountNumber{}

	for _, number := range chunk {
		if number%2 == 0 {
			amount.even++
			continue
		}
		amount.odd++
	}

	// Send data do channel
	channel <- amount

}
