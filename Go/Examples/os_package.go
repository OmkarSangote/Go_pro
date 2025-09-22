package main

import (
	"fmt"
	"os" // Use the os package instead
)

func main() {
	// Get current working directory
	cwd, err := os.Getwd() // Correct function to get the current working directory
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Current Directory:", cwd)
}
