package main

import "fmt"

func main() {
	var a, b int
	var operator string
	fmt.Println("Enter the the 2 numbers")
	fmt.Scanln(&a, &b)
	fmt.Println("Enter the airthmeatic operation")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Println("Addition", a+b)
	case "-":
		fmt.Println("Sub", a-b)
	case "*":
		fmt.Println("Product", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Division not possible")
		} else {
			fmt.Println("Division", a/b)
		}
	default:
		fmt.Println("Operation not defined")
	}
}
