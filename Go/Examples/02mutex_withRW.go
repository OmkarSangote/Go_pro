package main

import (
	"fmt"
	"sync"
)

var (
	data = make(map[int]int)
	mu   sync.RWMutex
)

func readData(key int, wg *sync.WaitGroup) {
	mu.RLock() // Acquire read lock
	// mu.RLock() is called to acquire the read lock, allowing the goroutine to read from the map while other reads are happening (but blocking any writes).
	value := data[key]
	mu.RUnlock() // Release read lock
	fmt.Println("Read value:", value)
	wg.Done()
}

func writeData(key, value int, wg *sync.WaitGroup) {
	mu.Lock() // Acquire write lock
	data[key] = value
	mu.Unlock() // Release write lock
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeData(i, i*i, &wg) // Writing data
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go readData(i, &wg) // Reading data
	}
	wg.Wait()
}
