package main

import (
	"fmt"
)

func main() {
	var size int

	// Prompt user for the size of the slice
	fmt.Println("Enter the size of the slice: ")
	fmt.Scan(&size)

	// Initialize an empty slice
	slice := make([]int, 0, size)

	// Prompt user for the values of the slice
	fmt.Println("Enter the values of the slice: ")
	for i := 0; i < size; i++ {
		var value int
		fmt.Scan(&value)
		slice = append(slice, value) // Append user input to the slice
	}

	// Display the created slice
	fmt.Println("Created slice:", slice)
}
