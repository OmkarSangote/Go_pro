// 2. Using errors.Is, errors.As, and fmt.Errorf for Error Chaining

// Go 1.13 introduced errors.Is and errors.As for checking and unwrapping errors, and they are useful when dealing with wrapped errors.

// errors.Is Example
// errors.Is checks whether an error matches a specific error, even if it has been wrapped.

package main

import (
	"errors"
	"fmt"
)

// Predefined error
var ErrNotAllowed = errors.New("action not allowed")

// Function that wraps the error
func performAction(action string) error {
	if action == "delete" {
		return fmt.Errorf("operation failed: %w", ErrNotAllowed)
	}
	return nil
}

func main() {
	err := performAction("delete")
	if errors.Is(err, ErrNotAllowed) {
		fmt.Println("Operation not allowed")
	} else {
		fmt.Println("Operation successful")
	}
}

// In this example, errors.Is checks if the wrapped error matches ErrNotAllowed.
