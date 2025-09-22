package main

import (
	"fmt"
	"sync"
)

var (
	count int
	mu    sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	mu.Lock() // Acquire the lock
	count++
	mu.Unlock() // Release the lock
	wg.Done()   // wg.Done() tells the WaitGroup that this goroutine has finished its work.
}

func main() {
	var wg sync.WaitGroup // It will keep track of all the goroutines so the program can wait for them to finish.
	for i := 0; i < 1000; i++ {
		wg.Add(1)         // Tells the compiler that there is another go routine in line , add the next go worker
		go increment(&wg) // Launching 1000 goroutines
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Final count:", count)
}
