/*
Lab Exercises: Rate Limiting, Debouncing, and Throttling in Go

 1. Rate Limiting Lab Exercise

Problem Statement:

You are tasked with creating an API that receives requests and processes them at a controlled rate.

In this exercise, you will build a system that limits the rate at which requests are handled.

The program should use Go's `time.Ticker` or `time.After` for rate limiting and ensure that only 5 requests can be processed per second.

Requirements:

- Implement rate limiting such that no more than 5 requests are processed in one second.

- Simulate an incoming stream of requests using a `for` loop.

- Display a message when a request is processed and when a request is dropped due to rate limiting.

Expected Output Example:

Processing request 1
Processing request 2
Processing request 3
Processing request 4
Processing request 5
Rate limit exceeded, dropping request 6
has context menu
*/
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func sendData(ch chan<- string) {
// 	ticker := time.NewTicker(1 * time.Second) // Create a ticker that ticks every second
// 	defer ticker.Stop()                       // Stop the ticker once we're done

// 	for i := 1; i <= 10; i++ { // Simulate 10 requests
// 		select {
// 		case <-ticker.C: // Wait for the next tick
// 			if i <= 5 { // Only process the first 5 requests
// 				ch <- fmt.Sprintf("request %d", i) // Send data through the channel
// 			} else {
// 				fmt.Println("Rate limit exceeded, dropping request", i)
// 			}
// 		}
// 	}
// 	close(ch) // Close the channel when done
// }

// func main() {
// 	ch := make(chan string) // Create a channel

// 	// and after each tick, we send a message through the channel (ch <- "Message X")
// 	go sendData(ch) // Start the goroutine that sends data to the channel every second

// 	for msg := range ch { // Receive messages from the channel
// 		fmt.Println("Processing", msg) // Print the received messages
// 	}

// 	fmt.Println("All messages received.")
// }

//OR

package main

import (
	"fmt"
	"time"
)

func main() {
	// Simulate incoming requests
	requests := make(chan int, 20)

	// Generate 20 requests
	for i := 1; i <= 20; i++ {
		requests <- i
	}
	close(requests)

	// Rate limit: process 5 requests per second
	limiter := time.Tick(200 * time.Millisecond) // 5 requests per second (200ms interval)

	for req := range requests {
		<-limiter // Wait for the limiter tick
		if req <= 5 {
			fmt.Println("Processing request", req)
		} else {
			fmt.Println("Rate limit exceeded, dropping request", req)
		}

	}
}
