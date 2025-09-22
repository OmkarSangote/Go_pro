// Example 1: Appliance Interface
// Problem:
// Your task is to create an interface Appliance with two methods: Start() and Stop(). You should then create two structs: Washer and Refrigerator, each representing an appliance. Implement the methods for each appliance so that they can start and stop.

// Create instances of the structs

// Declare a variable of type Appliance

// Call the methods for each appliance

// Define a struct for Washer

// Define a struct for Refrigerator

// Define methods for Washer
// Method to Start the washer

// Method to Stop the washer

// Define methods for Refrigerator
// Method to Start the refrigerator

// Method to Stop the refrigerator

// Define an interface Appliance with methods Start and Stop

package main

import "fmt"

func main() {
	// Create instances of the structs
	washer := Washer{brand: "LG"}
	fridge := Refrigerator{brand: "Samsung"}

	// Declare a variable of type Appliance
	var a Appliance

	// Assign Washer to Appliance
	a = washer
	a.Start()
	a.Stop()

	fmt.Println()

	// Assign Refrigerator to Appliance
	a = fridge
	a.Start()
	a.Stop()
}

// Define a struct for Washer
type Washer struct {
	brand string
}

// Define a struct for Refrigerator
type Refrigerator struct {
	brand string
}

// Define methods for Washer
func (w Washer) Start() {
	fmt.Printf("Washer %s is starting\n", w.brand)
}

func (w Washer) Stop() {
	fmt.Printf("Washer %s is stopping\n", w.brand)
}

// Define methods for Refrigerator
func (r Refrigerator) Start() {
	fmt.Printf("Refrigerator %s is starting\n", r.brand)
}

func (r Refrigerator) Stop() {
	fmt.Printf("Refrigerator %s is stopping\n", r.brand)
}

// Define an interface Appliance
type Appliance interface {
	Start()
	Stop()
}
