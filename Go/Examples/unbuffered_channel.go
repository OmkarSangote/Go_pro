package main

import "fmt"

func main() {
	// Create an unbuffered channel of type string
	ch := make(chan string)

	// Start a goroutine
	go func() {
		// Send data to the channel (this will block until the value is received)
		ch <- "Hello, world"
	}()

	// Receive the value from the channel (this will also block until a value is sent)
	msg := <-ch

	// Print the received message
	fmt.Println(msg)
}
