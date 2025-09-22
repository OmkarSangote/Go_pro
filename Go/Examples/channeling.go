package main

import (
	"fmt"
	"time"
)

// Generates squares of numbers and sends them to the first channel
func generateNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		square := i * i
		fmt.Println("Generated:", square)
		ch <- square
		time.Sleep(1 * time.Second) // Simulate time for generating numbers
	}
	close(ch) // Close the channel when done
}

// Receives numbers from the first channel, multiplies them, and sends to the second channel
func multiplyNumbers(ch1, ch2 chan int) {
	for num := range ch1 {
		result := num * 3
		fmt.Println("Processed:", result)
		ch2 <- result
		time.Sleep(1 * time.Second) // Simulate time for processing numbers
	}
	close(ch2) // Close the channel when done
}

func main() {
	ch1 := make(chan int) // First channel for generated numbers
	ch2 := make(chan int) // Second channel for processed numbers

	go generateNumbers(ch1)      // Start goroutine for generating numbers
	go multiplyNumbers(ch1, ch2) // Start goroutine for processing numbers

	// Receive and print final results
	for result := range ch2 {
		fmt.Println("Final Result:", result)
	}
}
