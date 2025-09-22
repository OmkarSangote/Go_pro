package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define command-line flags
	op := flag.String("op", "add", "Operation to perform (add, sub, mul, div)")
	num1 := flag.Float64("num1", 0, "First number")
	num2 := flag.Float64("num2", 0, "Second number")

	// Parse the command-line flags
	flag.Parse()

	// Perform the operation based on the provided flag
	var result float64
	switch *op {
	case "add":
		result = *num1 + *num2
	case "sub":
		result = *num1 - *num2
	case "mul":
		result = *num1 * *num2
	case "div":
		if *num2 == 0 {
			fmt.Println("Error: Division by zero")
			os.Exit(1)
		}
		result = *num1 / *num2
	default:
		fmt.Println("Unknown operation:", *op)
		os.Exit(1)
	}

	// Print the result
	fmt.Printf("Result: %f\n", result)
}
