// Basic dead lock
package main

func main() {
	ch := make(chan int)

	// This send operation will block because there's no receiver
	ch <- 1

	// The program will deadlock here, and we will see a runtime error
}
