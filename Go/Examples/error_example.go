/*
Lab Exercise 1: Understanding Basic Error Creation and Handling in Go

Problem Statement:

You are required to create a Go program that simulates basic error creation using the `errors.New()` function and handles error conditions.

In this exercise, you will:

1. Accept two string inputs from the user.

2. Compare the two strings.

3. If the two strings are not the same, you will create and display a custom error using `errors.New()`.

4. If the strings match, print a success message.

Steps to Implement:

1. Create a function to compare the strings.
2. If the strings don't match, generate a custom error using `errors.New()`.
3. Handle the error by printing a message to the user.

Requirements:
- Use the `errors.New()` function to create custom error messages.
- Use an `if` statement to check for the error condition (whether the strings are equal or not).
- Print the error message if an error occurs, otherwise print "Strings match".

---

Skeleton Code for Students:

package main

import (
    "errors"
    "fmt"
)

func main() {
    // Take two string inputs
    str1 := "Your first string here"
    str2 := "Your second string here"

    // Create an error message
    // myError := errors.New("YOUR ERROR MESSAGE HERE")

    // Compare the strings
    // If they don't match, print the error message
    // If they match, print a success message
}

*/

package main

// import the errors package
import (
	"errors"
	"fmt"
)

var ErrNotAllowed = errors.New("action not allowed")

func main() {
	// var ip, op string
	// fmt.Println("Enter the first string")
	// fmt.Scanf("%s", &ip)
	// fmt.Println("Enter the second string")
	// fmt.Scanf("%s", &op)

	// // create error using New() function
	// myError := errors.New("String mismatch")

	// if ip != op {
	// 	fmt.Println(myError)
	// } else {
	// 	fmt.Println("Both strings are equal")
	// }

	age := -5
	myerror2 := fmt.Errorf("Invalid age : %d, since age cannot be in negative", age)
	if age <= 0 {
		fmt.Println(myerror2)
	}

	// Example 1.................................
	// package main

	// import (
	//     "fmt"
	// )

	// func main() {
	//     age := -5
	//     error := fmt.Errorf("invalid age: %d. Age cannot be negative.", age)
	//     fmt.Println(error) // Output: invalid age: -5. Age cannot be negative.
	// }

	//Example 2......................................

	err := performAction("delete")
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: failed to perform delete action: action not allowed
	}
}

func performAction(action string) error {
	if action == "delete" {
		return fmt.Errorf("failed to perform %s action: %w", action, ErrNotAllowed)
	}
	return nil
}
