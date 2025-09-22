/*
Labs to practice file handling using the `os` package in Go. Each exercise comes with a problem statement and a solution example.


Exercise 1: Create and Write to a File

Problem Statement:
Create a Go program that:
- Prompts the user for a file name.
- Writes the string "Hello, GoLang!" to the file.
- Closes the file after writing.
*/

package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	var filename string
	fmt.Println("Enter the filename")
	fmt.Scanf("%s", &filename)
	// Open the file in append mode, creating it if it doesn't exist
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Append text to the file
	_, err = file.WriteString("Hello, GoLang!!\n")
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}

	fmt.Println("Data appended successfully!")

	logfile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Write the current timestamp to the file
	timestamp := time.Now().Format(time.RFC3339)
	_, err = logfile.WriteString("Log entry at: " + timestamp + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Log entry added successfully!")

	// Create a buffer to hold the data
	buf := make([]byte, 1024)

	// Read from the file
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break // End of file reached
			}
			fmt.Println("Error reading file:", err)
			return
		}
		// Print the content read
		fmt.Print(string(buf[:n]))
	}
}
