/*
Project Activity: "Real-Time Order Processing System" with Mutex in Golang

Objective:
Develop a concurrent order processing system that simulates multiple order managers concurrently updating a shared inventory for a retail business.

The system will showcase how sync.Mutex can be used to prevent data inconsistencies when managing inventory during real-time order placement and cancellations.

Scenario:
1. Imagine you are building a real-time inventory management system for an e-commerce platform.

2. Multiple order managers (representing different microservices or users) will handle incoming customer orders and cancellations concurrently.

3. The inventory needs to be updated properly without causing race conditions, ensuring the stock levels are accurate at all times.

Requirements:
1. Shared Resource: A centralized inventory where multiple products have stock levels.

2. Concurrency: Multiple order managers (goroutines) place and cancel orders simultaneously.

3. Thread Safety: Use a sync.Mutex to synchronize access to the shared inventory, preventing inconsistent stock levels.
Operations:

Place Order: Decrease the stock level when an order is placed.
Cancel Order: Increase the stock level when an order is canceled.
Check Stock: Retrieve the current stock level for a given product.
Output: Log each order placement, cancellation, and the updated stock levels.

Constraints:
1. Ensure no orders can be placed if the product's stock is zero.

2. Orders must reflect correct stock updates without race conditions, even under heavy concurrency.

3. Use sync.Mutex to ensure only one goroutine is updating a product's stock at a time.

                        Instructions:

Inventory Setup:

Create an Inventory struct to manage stock for multiple products.

The inventory should use a map[string]int where the key is the product name and the value is the stock level.

Use a sync.Mutex to control access to the stock levels.


Order Manager:

Implement the PlaceOrder(product string, quantity int) and CancelOrder(product string, quantity int) methods that update the stock levels.

Implement the CheckStock(product string) method to safely retrieve the stock level.
Goroutines:

Simulate multiple concurrent order managers placing and canceling orders on the same products.

Use a sync.WaitGroup to wait for all goroutines to complete.


            Step-by-Step Instructions:
Step 1: Lock the Inventory before modifying the stock in PlaceOrder and CancelOrder and unlock it once the operation is done.

Step 2: Safely access the Inventory in CheckStock by locking it to retrieve the current stock for a product.

Step 3: Test the program by running multiple goroutines to place and cancel orders concurrently and observe the final stock levels.
*/

package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// LogProcessor manages concurrent log writing and reading
type LogProcessor struct {
	logFile string
	mu      sync.Mutex
}

// NewLogProcessor creates a new LogProcessor with a specified log file
func NewLogProcessor(filename string) *LogProcessor {
	return &LogProcessor{
		logFile: filename,
	}
}

// WriteLog writes a log entry to the log file
func (lp *LogProcessor) WriteLog(workerID int, logMsg string) {
	// TODO: Lock the mutex, open the file in append mode, write the log entry, then unlock
}

// ReadLogs reads the latest n log entries from the log file
func (lp *LogProcessor) ReadLogs(n int) {
	// TODO: Lock the mutex, open the file in read mode, read the last n lines, then unlock
}

// worker simulates a worker that writes logs concurrently
func worker(id int, lp *LogProcessor, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		logMsg := fmt.Sprintf("Worker %d: Log entry %d at %s", id, i+1, time.Now().Format(time.RFC3339))
		lp.WriteLog(id, logMsg)
		time.Sleep(time.Millisecond * 100) // Simulate some delay between logs
	}
}

func main() {
	logFile := "logs.txt"

	// Create a new log processor
	lp := NewLogProcessor(logFile)

	// Ensure the log file is empty at the start
	os.Remove(logFile)

	var wg sync.WaitGroup
	numWorkers := 5
	wg.Add(numWorkers)

	// Start multiple worker goroutines
	for i := 0; i < numWorkers; i++ {
		go worker(i, lp, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()

	// TODO: After all workers are done, read the latest log entries from the file
	fmt.Println("Final Logs:")
}
