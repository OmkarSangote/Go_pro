package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker((1 * time.Second)) // Rate limit to 1 msg/sec
	defer ticker.Stop()

	for i := 0; i < 5; i++ {
		fmt.Println("Sending msg", i+1)
		<-ticker.C //wait for next tick
	}
}
