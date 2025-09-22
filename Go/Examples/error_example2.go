/*
Lab Exercise 2: Handling Errors with `fmt.Errorf()` in Go

Problem Statement:

You are tasked to write a Go program that simulates error handling using the `fmt.Errorf()` function.

Unlike the `errors.New()` function, `fmt.Errorf()` allows you to format error messages dynamically by including variables in the error message.

In this lab, you will:
1. Create a number int variable.
2. If the number is negative, generate a custom error message using `fmt.Errorf()`.
3. If the number is positive or zero, print the number.

4. Create a custom message using errors.New()

5. Try to wrap it in Errorf()

Steps to Implement:
1. Create a function that accepts an integer input.
2. If the number is negative, use `fmt.Errorf()` to create an error message that includes the input value.
3. Handle the error by printing the formatted error message.
4. If no error occurs, print the valid number.
*/

package main

import (
	"errors"
	"fmt"
)

var ErrNotAllowed = errors.New("negative number not allowed")

func performAction(num int) error {
	if num < 0 {
		return fmt.Errorf("failed to perform %d action: %w", num, ErrNotAllowed)
	} else {
		fmt.Println(num)
	}
	return nil
}

func main() {

	var num int
	fmt.Println("Enter the number:")
	fmt.Scanln(&num)

	err := performAction(num)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
