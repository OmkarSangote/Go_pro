/*
Project Overview
Create a real-time stock market tracker that simulates stock price changes and allows users to buy and sell stocks.

You’ll use Go's concurrency features, including channels and goroutines, to handle real-time updates and manage user transactions while avoiding race conditions and deadlocks.

	Features

Simulated Stock Prices: Generate random stock prices for a list of predefined stocks.

Real-Time Updates: Continuously update stock prices at random intervals, simulating a real stock market environment.

User Transactions: Allow users to buy and sell stocks, updating their portfolios accordingly.

Concurrency Handling: Use goroutines and channels to manage multiple users making transactions at the same time.

Error Handling: Ensure users cannot buy more stocks than they have funds for or sell stocks they do not own.

Portfolio View: Provide a way for users to view their current stock holdings and available balance.

	Detailed Steps

1. Define Stock Structure:

2. Create a struct for Stock with fields like Name, Price, and Quantity.

3. Define User Structure:

4. Create a struct for User that includes fields like Balance and a map of owned stocks.

	Simulated Stock Price Updates:

1. Use a goroutine to simulate price changes for each stock at random intervals.

	Transaction Functions:

Implement functions for buying and selling stocks:
Buy Stock:
Check if the user has enough balance, update the user’s portfolio and balance.

Sell Stock: Check if the user owns the stock, update the portfolio and balance accordingly.

Concurrency with Channels:
--> Use channels to handle transactions safely, ensuring that no race conditions occur when multiple users are trying to buy or sell stocks simultaneously.

--> Use a channel for sending price update notifications to users.

	User Interface:

Provide a simple command-line interface for users to interact with the application, allowing them to view stock prices, buy, and sell stocks.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Stock structure to represent stock information
type Stock struct {
	Name  string
	Price float64
}

// User structure to represent user information
type User struct {
	Balance float64
	Stocks  map[string]int // Stock name and quantity
}

// Global variables
var (
	stocks       = []Stock{{"AAPL", 150.0}, {"GOOGL", 2800.0}, {"AMZN", 3400.0}} // List of stocks
	users        = make(map[string]*User)                                        // Store users
	mu           sync.Mutex                                                      // Mutex for synchronizing access
	priceUpdates = make(chan Stock)                                              // Channel for stock price updates
)

// Function to simulate stock price updates
func updateStockPrices() {
	for {
		time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second) // Random interval for updates
		for i := range stocks {
			// Simulate price change
			stocks[i].Price += rand.Float64()*10 - 5
			priceUpdates <- stocks[i] // Send updated stock to the channel
		}
	}
}

// Function to handle user transactions (buy/sell)
func transaction(userName string, action string, stockName string, quantity int) {
	mu.Lock()         // Lock to ensure safe access to shared data
	defer mu.Unlock() // Unlock when the function returns

	user, exists := users[userName]
	if !exists {
		fmt.Printf("User %s does not exist.\n", userName)
		return
	}

	var stockPrice float64
	// Find the stock price
	for _, stock := range stocks {
		if stock.Name == stockName {
			stockPrice = stock.Price
			break
		}
	}

	if action == "buy" {
		totalCost := stockPrice * float64(quantity)
		if user.Balance >= totalCost {
			user.Balance -= totalCost
			user.Stocks[stockName] += quantity
			fmt.Printf("%s bought %d shares of %s.\n", userName, quantity, stockName)
		} else {
			fmt.Printf("%s does not have enough balance to buy %d shares of %s.\n", userName, quantity, stockName)
		}
	} else if action == "sell" {
		if user.Stocks[stockName] >= quantity {
			totalEarning := stockPrice * float64(quantity)
			user.Balance += totalEarning
			user.Stocks[stockName] -= quantity
			fmt.Printf("%s sold %d shares of %s.\n", userName, quantity, stockName)
		} else {
			fmt.Printf("%s does not own enough shares of %s to sell.\n", userName, stockName)
		}
	}
}

// Main function
func main() {
	// Initialize users
	users["Alice"] = &User{Balance: 10000, Stocks: make(map[string]int)}
	users["Bob"] = &User{Balance: 5000, Stocks: make(map[string]int)}

	// Start stock price updates in a separate goroutine
	go updateStockPrices()

	// Simulate some transactions
	time.Sleep(1 * time.Second) // Allow some time for price updates
	transaction("Alice", "buy", "AAPL", 10)
	transaction("Bob", "sell", "GOOGL", 5) // Should fail, as Bob doesn't own GOOGL
	transaction("Alice", "sell", "AAPL", 5)

	// Keep the main function running
	select {}
}
