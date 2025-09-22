package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this function as done when it returns
	fmt.Printf("Worker %d started\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)         // Increment the counter
		go worker(i, &wg) // Start a new goroutine
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All workers completed")
}
