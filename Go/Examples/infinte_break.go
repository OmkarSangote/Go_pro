package main

import "fmt"

func main() {
	var num int

	for {
		fmt.Println("Enter the number")
		fmt.Scan(&num)
		if num < 0 {
			break
		}
	}
}
