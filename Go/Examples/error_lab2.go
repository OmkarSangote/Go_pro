/*
Lab Exercise: User Registration Error Handling

Problem Statement:

You are developing a user registration system.

The system should validate user inputs, specifically ensuring that the username is not empty and that the age is a valid number.

If validation fails, the system should return appropriate errors.

You will implement a custom error type to handle registration errors and check for this specific error type using `errors.As()`.

Requirements:

1. Create a custom error type called `RegistrationError` to represent errors related to user registration.
2. Implement the `Error()` method for the `RegistrationError` type.
3. Write a function `validateUserRegistration()` that validates the username and age. It should return a `RegistrationError` if the username is empty or if the age is invalid (negative).
4. In the `main()` function, call `validateUserRegistration()` and use the `checkRegistrationError()` function to verify the error type using `errors.As()
*/

package main

import (
	"errors"
	"fmt"
)

// Define the RegistrationError struct
type RegistrationError struct {
	Message string
}

// Implement the Error method for RegistrationError
func (e RegistrationError) Error() string {
	return e.Message
}

// Function that validates user registration
func validateUserRegistration(username string, age int) error {
	if username == "" {
		return RegistrationError{Message: "username cannot be empty"}
	}
	if age < 0 {
		return RegistrationError{Message: "age cannot be negative"}
	}
	return nil
}

// Function to check the type of registration error
func checkRegistrationError(err error) {
	var regErr RegistrationError
	// Use errors.As() to check if the error is of type RegistrationError
	if errors.As(err, &regErr) {
		fmt.Println("This is a RegistrationError:", regErr.Message)
	} else {
		fmt.Println("This is a different type of error:", err)
	}
}

func main() {
	// Simulate a user registration scenario
	err := validateUserRegistration("", -1) // Example with invalid inputs
	checkRegistrationError(err)
}
