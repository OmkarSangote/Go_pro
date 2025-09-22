//Golang Error Handling

// When an error occurs, the execution of a program stops completely with a built-in error message.

//We can handle errors using:
/*
❑ New() Function
❑ Errof() Function

*/

// Go Error using New() Function

/*
We can use the New() function to handle an error.

This function is defined inside the errors package and
allows us to create our own error message.

*/

// the New() function belongs to the errors package in Go, and it is used to create a new error with a custom message.

// Package: errors
// Function: New()
// Purpose: To create a new error with a custom error message.
// package main
// import (
//     "errors"
//     "fmt"
// )

// func main() {
//     myError := errors.New("This is a custom error message")
//     fmt.Println(myError)
// }

package main

// import the errors package
import (
	"errors"
	"fmt"
)

func main() {
	message := "Hello"

	// create error using New() function
	myError := errors.New("WRONG MESSAGE")

	if message != "Welcome" {
		fmt.Println(myError)
	}
}
