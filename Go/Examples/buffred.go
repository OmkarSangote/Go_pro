package main

import (
	"fmt"
	"time"
)

type Values struct {
	v1, v2, v3, v4, v5, v6, v7 int
}

func main() {
	ch := make(chan Values, 1)

	// go func() {
	// 	for i := 1; i <= 5; i++ {
	// 		fmt.Printf("Sending %d\n", i)
	// 		ch <- i
	// 		ch <- i * 2
	// 		ch <- i * 4
	// 		ch <- i * 6
	// 		ch <- i * 8
	// 		ch <- i * 10
	// 		ch <- i * 12
	// 	}
	// 	close(ch)
	// }()

	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Sending: %d, %d, %d, %d, %d, %d, %d\n", i, i*2, i*4, i*5, i*6, i*7, i*8)
			//ch <- Values{i, i * 2, i * 4, i * 5, i * 6, i * 7, i * 8}
			//time.Sleep(0 * time.Second)
			select {
			case ch <- Values{i, i * 2, i * 4, i * 5, i * 6, i * 7, i * 8}:
				// Successfully sent value
			default:
				// Channel is full, handle accordingly
				fmt.Println("Channel is full, cannot send value")
			}
			time.Sleep(0 * time.Second) // No delay in sending
		}
		close(ch)
	}()

	for value := range ch {
		//fmt.Printf("Received : %d\n", value)
		fmt.Printf("Received: %d, %d, %d, %d, %d, %d, %d\n", value.v1, value.v2, value.v3, value.v4, value.v5, value.v6, value.v7)
		time.Sleep(2 * time.Second)
	}
}
