/*
 
The errors.As() function in Go is used for type assertion when dealing with errors.
 
It allows you to check if an error is of a specific type and, if so, access that error's details.
 
 
In Go, an error can be wrapped or represented by different types. For example, you might have custom error types, and you want to check whether an error is of a certain type.
 
errors.As() helps to convert or assert an error into a specific type.
 
When you have an error and you want to see if it's a specific kind of error (such as a custom error), you use errors.As().
 
How It Works:
 
errors.As() checks if an error matches a certain type.
If it matches, it lets you work with that error as the specific type you're checking for.
 
Example:
Let's say we have a custom error type called MyError, and we want to check if a given error is of that type.
 
*/
 
package main
 
import (
    "errors"
    "fmt"
)
 
type MyError struct {
    Message string
}
 
// Implement the Error method to satisfy the error interface
func (e MyError) Error() string {
    return e.Message
}
 
func checkError(err error) {
    var myErr MyError  // myErr is an instance of MyError struct
 
    // Use errors.As() to check if the error is of type MyError
    if errors.As(err, &myErr) {
        fmt.Println("This is a MyError:", myErr.Message)
    } else {
        fmt.Println("This is a different type of error")
    }
}
 
func main() {
    err1 := MyError{Message: "Something went wrong"}
    err2 := errors.New("Another generic error")
 
    checkError(err1) // This will match MyError
    checkError(err2) // This will not match MyError
}