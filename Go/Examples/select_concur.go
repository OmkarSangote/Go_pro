/*
Select with Buffered and Unbuffered Channels

The select statement allows you to work with multiple channels and handle whichever one is ready first.

*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string, 1) // Buffered channel
	ch2 := make(chan string)    // Unbuffered channel

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from buffered channel"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from unbuffered channel"
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
