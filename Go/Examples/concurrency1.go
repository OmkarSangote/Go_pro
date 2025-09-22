package main

import (
	"fmt"
	"time"
)

func printnumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep((100 * time.Millisecond)) // Not right approach for con-currency instead use channeling
	}
}

func printletters() {
	for ch := 'a'; ch <= 'e'; ch++ {
		fmt.Printf("%c\n", ch)
		time.Sleep((150 * time.Millisecond))
	}
}

func main() {
	go printnumbers()
	go printletters()
	time.Sleep((1 * time.Second))
}
