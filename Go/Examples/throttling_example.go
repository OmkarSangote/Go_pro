package main

import (
	"fmt"
	"time"
)

// throttle limits the execution of a function to once every interval
func throttle(fn func(), interval time.Duration) func() {
	var lastExecution time.Time

	return func() {
		now := time.Now()                      // current time
		if now.Sub(lastExecution) > interval { // Execute if enough time has passed
			lastExecution = now
			fn()
		}
	}
}

func main() {
	throttledFunc := throttle(func() {
		fmt.Println("Throttled function executed!")
	}, 1*time.Second)

	// Simulate events happening quickly, but function only executes once every 1 second
	for i := 0; i < 5; i++ {
		throttledFunc()
		time.Sleep(500 * time.Millisecond) // Simulate rapid calls
	}

	// Sleep long enough to allow throttling to happen again
	time.Sleep(2 * time.Second)
	throttledFunc() // This will execute because 2 seconds have passed
}
