/*
Problem 6: Mixed Channel Communication (Buffered and Unbuffered Channels)

Problem Statement:
-->Create a system where two channels are used: one buffered and one unbuffered. Each channel should send different types of messages (e.g., status updates and error reports).

--> Use the select statement to listen to both channels and process whichever message arrives first.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Status updated : Task completed "
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Error: Nteork issue"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
