/* Banking Transaction System with Error Checking

Problem Statement:

You are developing a simple banking transaction system that allows users to perform deposit and withdrawal operations.

Certain actions, such as withdrawing more money than is available in the account, should not be allowed.

If a user attempts to make such a transaction, the program should return a wrapped error indicating that the transaction is not allowed.

Requirements:

1. Create a predefined error variable for insufficient funds.

2. Write a function that checks the transaction type and amount and wraps the predefined error if a withdrawal exceeds the available balance.

3. In the `main()` function, simulate a few transactions and handle the errors appropriately using `errors.Is`.

Tasks:

1. Define a variable `ErrInsufficientFunds` for the error indicating insufficient funds for a withdrawal.

2. Implement the `processTransaction()` function that takes an action ("deposit" or "withdraw") and an amount, checks against the current balance, and returns the wrapped error if necessary.

3. In the `main()` function, simulate a few transactions, and handle the error messages by checking if the error corresponds to `ErrInsufficientFunds` using `errors.Is`.



Skeleton Code for Students:

package main

import (
    "errors"
    "fmt"
)

// Define a variable for the error indicating insufficient funds
// var ErrInsufficientFunds = ...

// Function that processes transactions
func processTransaction(action string, amount float64, balance float64) (float64, error) {
    // Implement the transaction logic and return the wrapped error if needed
}

// The main function
func main() {
    // Initialize the account balance
    balance := 100.0  // Example starting balance

    // Simulate some transactions and handle the errors using errors.Is
}

*/

package main

import (
	"errors"
	"fmt"
)

// Define a variable for the error indicating insufficient funds
var ErrInsufficientFunds = errors.New("insufficient funds for withdrawal")

// Function that processes transactions
func processTransaction(action string, amount float64, balance float64) (float64, error) {
	if action == "withdraw" {
		if amount > balance {
			return balance, fmt.Errorf("transaction failed: %w", ErrInsufficientFunds)
		}
		return balance - amount, nil
	} else if action == "deposit" {
		return balance + amount, nil
	}
	return balance, errors.New("invalid transaction action")
}

func main() {
	// Initialize the account balance
	balance := 100.0 // Example starting balance

	// Simulate withdrawal transaction
	amountToWithdraw := 150.0
	balance, err := processTransaction("withdraw", amountToWithdraw, balance)
	if err != nil {
		// Check if the error corresponds to ErrInsufficientFunds using errors.Is
		if errors.Is(err, ErrInsufficientFunds) {
			fmt.Println("Error:", err)
			fmt.Println("Withdrawal unsuccessful due to insufficient funds.")
		} else {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Printf("Withdrawal successful. New balance: %.2f\n", balance)
	}

	// Simulate deposit transaction
	amountToDeposit := 50.0
	balance, err = processTransaction("deposit", amountToDeposit, balance)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Deposit successful. New balance: %.2f\n", balance)
	}

	// Simulate invalid transaction
	balance, err = processTransaction("transfer", 30.0, balance)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
