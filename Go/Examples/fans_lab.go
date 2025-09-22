/*
Problem Statement: Fan-In and Fan-Out Concurrency Pattern in Go

You are required to create a Go program that demonstrates the Fan-In and Fan-Out concurrency pattern. The system should perform the following tasks:

1. Fan-Out: Multiple goroutines will process different sets of data concurrently. The data consists of integers, and each goroutine should process a subset of integers by squaring them.

2. Fan-In: After each goroutine finishes processing its set of integers, the results should be sent to a central goroutine, which collects and sums the squared values from all the goroutines.

                     Key Requirements:

1. Use Go channels to communicate between goroutines.
2. Use `sync.WaitGroup` to wait for all processing goroutines to finish.
3. The system should process the data concurrently using Fan-Out
4. A single goroutine should gather results from all workers (Fan-In) and calculate the total sum.
5. The program should handle a list of integers (input by the user) and output the sum of their squares.

 Expected Output Example:

- Input: [2, 4, 6, 8]
- Expected Output: "The sum of the squares is 120" (i.e., 2^2 + 4^2 + 6^2 + 8^2)
*/
// Worker function: Squares numbers and sends results to the results channel
package main

import (
	"fmt"
	"sync"
)

// Worker function: Squares numbers and sends results to the results channel
func worker(nums []int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that the goroutine is done
	sum := 0
	for _, num := range nums {
		sum += num * num
	}
	results <- sum // Send the sum of squares to the results channel
}

// Fan-In: Collect results from multiple workers and calculate the final total
func fanIn(results <-chan int, done chan<- int, numWorkers int) {
	total := 0
	for i := 0; i < numWorkers; i++ {
		total += <-results // Receive from each worker
	}
	done <- total // Send the final result to the done channel
}

func main() {
	// Input: A list of integers
	numbers := []int{2, 4, 6, 8}
	numWorkers := 2 // Number of workers (goroutines) to process data

	// Split data for workers (Fan-Out pattern)
	chunks := [][]int{
		numbers[:len(numbers)/2], // First half
		numbers[len(numbers)/2:], // Second half
	}

	// Channels for communication
	results := make(chan int, numWorkers) // To collect results from workers
	done := make(chan int)                // To signal when all results are collected

	// WaitGroup to synchronize goroutines
	var wg sync.WaitGroup

	// Launch worker goroutines (Fan-Out)
	for _, chunk := range chunks {
		wg.Add(1)
		go worker(chunk, results, &wg)
	}

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results) // Close results channel after all workers are done
	}()

	// Launch a goroutine to fan-in results
	go fanIn(results, done, numWorkers)

	// Receive and print the final result
	total := <-done
	fmt.Printf("The sum of the squares is %d\n", total)
}
