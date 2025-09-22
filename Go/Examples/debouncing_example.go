package main

import (
	"fmt"
	"time"
)

// debounce function waits for a delay before executing the function
// The `debounce` function takes a function `fn` and a `delay`
func debounce(fn func(), delay time.Duration) {
	var timer *time.Timer

	// If there's an existing timer, stop it to prevent the function from running
	if timer != nil {
		timer.Stop()
	}

	// Create a new timer that waits for the delay, then executes the function
	timer = time.AfterFunc(delay, fn)
}

func main() {
	// Call the debounce function with a 2-second delay
	// When the `debounce` function is called, it sets a timer (`time.AfterFunc`) to execute the function after the given `delay`.
	debounce(func() {
		fmt.Println("Debounced function executed after 2 seconds!")
	}, 2*time.Second)

	// Simulate another event happening 1 second later, so the previous execution is canceled
	time.Sleep(1 * time.Second)
	debounce(func() {
		fmt.Println("Debounced function executed again after a new 2-second delay!")
	}, 2*time.Second)

	// Sleep long enough to allow the final debounce to execute
	time.Sleep(3 * time.Second)
}
