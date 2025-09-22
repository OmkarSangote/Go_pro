package main

import "fmt"

func main() {
	var num int
	var sum int
	fmt.Println("Enter the number till which sum is needed")
	fmt.Scan(&num)
	for i := 0; i <= num; i++ {
		if i%2 != 0 {
			continue
		}
		sum += i
	}
	fmt.Printf("Sum of numbers is: %d\n", sum)
}
