/*
Problem Statement:
--> Write a Go program that simulates a grading system for two different classes.

--> Each class has a list of student grades.

--> You need to calculate the average grade for each class concurrently using goroutines. The results should then be passed through a channel and printed in the main function.

Instructions:

- Create a function `calculateAverage` that accepts a list of grades and calculates the average.

- Use two goroutines to concurrently calculate the average for two different classes.

- The main function should receive the two averages from the channel and display them.
*/
package main

import (
	"fmt"
)

// Function to calculate the average grade and send it to the channel
func calculateAverage(grades []int, ch chan float64) {
	total := 0
	for _, grade := range grades {
		total += grade
	}
	avg := float64(total) / float64(len(grades))
	ch <- avg // Send the average to the channel
}

func main() {
	gradesClass1 := []int{85, 90, 78, 92, 88}
	gradesClass2 := []int{76, 81, 85, 89, 90}

	// Create a channel to receive average grades
	ch := make(chan float64)

	// Start goroutines to calculate averages for both classes
	go calculateAverage(gradesClass1, ch)
	go calculateAverage(gradesClass2, ch)

	// Receive the averages from the channel
	avgClass1 := <-ch
	avgClass2 := <-ch

	// Print the results
	fmt.Printf("Average grade for Class 1: %.2f\n", avgClass1)
	fmt.Printf("Average grade for Class 2: %.2f\n", avgClass2)
}
