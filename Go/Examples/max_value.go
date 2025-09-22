package main

import (
	"fmt"
)

func main() {
	// Define a slice to hold integers
	var numbers []int

	// Number of entries to be added
	var n int
	fmt.Print("Enter the number of integers you want to add to the slice: ")
	fmt.Scan(&n)

	// Taking input from the user
	for i := 0; i < n; i++ {
		var value int
		fmt.Printf("Enter integer %d: ", i+1)
		fmt.Scan(&value)
		numbers = append(numbers, value)
	}

	// Initialize variables to hold the maximum value and its index
	if len(numbers) == 0 {
		fmt.Println("No integers were entered.")
		return
	}

	maxValue := numbers[0]
	maxIndex := 0

	// Use range to find the maximum value and its index
	for i, value := range numbers {
		if value > maxValue {
			maxValue = value
			maxIndex = i
		}
	}

	// Print the maximum value and its index
	fmt.Printf("Maximum value: %d, Index: %d\n", maxValue, maxIndex)
}