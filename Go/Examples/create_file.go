package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Create a new file or overwrite if it exists
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed after use

	// Write some text to the file
	_, err = file.WriteString("Hello, Go file handling!\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File created and data written successfully!")

	file, err = os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

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
