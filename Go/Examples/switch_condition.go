package main

import "fmt"

func main() {
	var num int64
	fmt.Println("Enter a number")
	fmt.Scanln(&num)
	switch {
	case num%2 == 0:
		fmt.Println("Even number")

	case num%2 != 0:
		fmt.Println("Odd number")
		fallthrough

	case num%1 == num:
		fmt.Println("Its a number")

	default:
		fmt.Println("Fall through test")
	}

}
