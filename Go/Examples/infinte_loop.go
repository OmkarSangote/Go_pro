package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Server is running .....")
	for {
		fmt.Println("Waiting for requests ...")
		time.Sleep(5 * time.Second)
	}
}
