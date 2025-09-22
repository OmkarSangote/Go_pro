// Extending Functionality with Reusable Components

// Embedded structs can be used to enhance entities with reusable functionalities, such as logging or error handling.

// Example: Adding Logging Capabilities to Multiple Components

package main

import "fmt"

// Logger struct provides logging functionality
type Logger struct {
	Level string
}

func (l Logger) Log(message string) {
	fmt.Printf("[%s] %s\n", l.Level, message)
}

// Server struct embeds Logger to add logging
type Server struct {
	IP     string
	Logger // Embed Logger for logging
}

// Database struct embeds Logger to add logging
type Database struct {
	Name   string
	Logger // Embed Logger for logging
}

func main() {
	server := Server{IP: "192.168.1.1", Logger: Logger{Level: "INFO"}}
	db := Database{Name: "CustomerDB", Logger: Logger{Level: "DEBUG"}}

	server.Log("Server started")    // Using embedded logger
	db.Log("Connected to database") // Using embedded logger
}

//Use Case: When you need to add logging functionality to multiple components (e.g., Server and Database), embedding a Logger struct allows each to reuse the logging logic without duplicating it.
