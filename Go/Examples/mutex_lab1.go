/*
Problem Statement: Bank Account with Multiple Transactions

You are tasked with simulating a basic Bank Account system using Go’s goroutines and sync.Mutex for synchronization.

The goal is to safely handle concurrent deposits and withdrawals on a shared bank account balance.

Multiple clients (goroutines) will perform deposits and withdrawals simultaneously, and you must ensure that no data race occurs.

Requirements:
1. Implement a `BankAccount` struct that holds the balance.
2. Use a `sync.Mutex` to protect access to the balance to prevent race conditions.
3. Implement methods `Deposit(amount int)` and `Withdraw(amount int)` that update the balance.
4. The program should simulate 5 clients concurrently performing deposits and withdrawals.
5. Ensure that after all transactions are complete, the final balance is printed correctly.

Input/Output:
- The balance starts at 0.
- Each client can either deposit or withdraw a random amount between $10 and $100.
- After all transactions, print the final balance.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// BankAccount struct to hold the balance and a mutex for synchronization
type BankAccount struct {
	balance int
	mu      sync.Mutex
}

// Deposit method to add money to the account
func (acc *BankAccount) Deposit(amount int) {
	acc.mu.Lock()         // Acquire lock before modifying balance
	defer acc.mu.Unlock() // Ensure lock is released after deposit
	acc.balance += amount // Add the deposit amount to the balance
	fmt.Printf("Deposited %d, New Balance: %d\n", amount, acc.balance)

	// The defer keyword in Go is used to schedule a function call to be executed after the surrounding function completes, regardless of how the function returns.
	// defer acc.mu.Unlock(), it ensures that the Unlock() method on the mutex acc.mu is called after the function finishes, even if the function returns early due to an error or another condition.
}

// Withdraw method to subtract money from the account
func (acc *BankAccount) Withdraw(amount int) {
	acc.mu.Lock()         // Acquire lock before modifying balance
	defer acc.mu.Unlock() // Ensure lock is released after withdrawal

	if acc.balance >= amount { // Check if sufficient balance exists
		acc.balance -= amount
		fmt.Printf("Withdrew %d, New Balance: %d\n", amount, acc.balance)
	} else {
		fmt.Printf("Insufficient funds to withdraw %d, Current Balance: %d\n", amount, acc.balance)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed random number generator with current time

	// Initialize the bank account with a zero balance
	account := &BankAccount{balance: 0}

	var wg sync.WaitGroup // WaitGroup to ensure all goroutines complete

	// Simulate 5 clients performing transactions concurrently
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment WaitGroup counter

		go func(clientID int) {
			defer wg.Done() // Decrement WaitGroup counter when goroutine completes

			// Randomly decide whether to deposit or withdraw
			for j := 0; j < 3; j++ { // Each client performs 3 transactions
				transactionType := rand.Intn(2) // 0 for deposit, 1 for withdraw
				amount := rand.Intn(91) + 10    // Random amount between $10 and $100

				if transactionType == 0 {
					account.Deposit(amount) // Perform a deposit
				} else {
					account.Withdraw(amount) // Perform a withdrawal
				}

				time.Sleep(time.Millisecond * 500) // Simulate delay between transactions
			}
		}(i)
	}

	wg.Wait() // Wait for all goroutines to finish

	// Print the final balance after all transactions
	fmt.Printf("Final Balance: %d\n", account.balance)
}
