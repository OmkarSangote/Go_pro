// Example 2: Payment System

// Problem:
// Create an interface Payment that has a method Pay(). Implement two payment methods: CreditCard and PayPal. Each payment method should have its own Pay() method, describing how the payment is processed.

// package main

// import "fmt"

// func main() {
// Create instances of the structs

// Declare a variable of type Payment

// Call the Pay method for each payment method
//}

// Define a struct for CreditCard

// Define a struct for PayPal

// Define methods for CreditCard
// Method to Pay using CreditCard

// Define methods for PayPal
// Method to Pay using PayPal

// Define an interface Payment with method Pay

package main

import "fmt"

func main() {
	// Create instances of the structs
	credit := CreditCard{holderName: "John Doe"}
	paypal := PayPal{email: "john@example.com"}

	// Declare a variable of type Payment
	var p Payment

	// Assign CreditCard to Payment
	p = credit
	p.Pay()

	fmt.Println()

	// Assign PayPal to Payment
	p = paypal
	p.Pay()
}

// Define a struct for CreditCard
type CreditCard struct {
	holderName string
}

// Define a struct for PayPal
type PayPal struct {
	email string
}

// Define methods for CreditCard
func (c CreditCard) Pay() {
	fmt.Printf("Processing payment through Credit Card of %s\n", c.holderName)
}

// Define methods for PayPal
func (p PayPal) Pay() {
	fmt.Printf("Processing payment through PayPal with email %s\n", p.email)
}

// Define an interface Payment
type Payment interface {
	Pay()
}
