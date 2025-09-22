/*
Problem Statement:

--> Create a Go program that simulates a simple ticket booking system.

--> One goroutine should represent the ticket booking counter, which waits for customers (other goroutines) to request tickets.

--> When a request is made, the booking counter should send back a confirmation message via channels.

--> The program should demonstrate how channels synchronize communication between multiple goroutines.

Instructions:
- Define a function `requestTicket` where customers request tickets from the booking counter.
- The `bookingCounter` function should wait for ticket requests from multiple customers and send back confirmation.
- Use channels to communicate between the `bookingCounter` and `requestTicket` functions.
- Ensure proper synchronization between goroutines using channels.
*/

package main

import (
	"fmt"
	"time"
)

func bookingCounter(ch chan string) {
	for i := 0; i < 3; i++ {
		request := <-ch // Receive request from customer
		fmt.Println("Processing request for:", request)
		time.Sleep(2 * time.Second)             // Simulate processing time
		ch <- "Ticket confirmed for " + request // Send confirmation
	}
}

func requestTicket(name string, ch chan string) {
	fmt.Println(name, "is requesting a ticket...")
	ch <- name           // Send ticket request
	confirmation := <-ch // Receive confirmation from booking counter
	fmt.Println(confirmation)
}

func main() {
	ch := make(chan string) // Create a channel

	go bookingCounter(ch) // Start booking counter goroutine

	go requestTicket("Customer 1", ch) // Start customer goroutine
	go requestTicket("Customer 2", ch)
	go requestTicket("Customer 3", ch)

	time.Sleep(10 * time.Second) // Give enough time for goroutines to complete
}
