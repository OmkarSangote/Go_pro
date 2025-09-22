//Exercise 6: Copy a File

// Problem Statement:
// Create a Go program that:
// - Prompts the user for two file names: the source file and the destination file.
// - Copies the content of the source file to the destination file.
// - Handles any errors that may occur during the process.

// Solution Example:

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var srcFile, dstFile string

	// Get source and destination file names from the user
	fmt.Print("Enter source file name: ")
	fmt.Scanln(&srcFile)
	fmt.Print("Enter destination file name: ")
	fmt.Scanln(&dstFile)

	// Open source file
	src, err := os.Open(srcFile)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer src.Close()

	// Create destination file
	dst, err := os.Create(dstFile)
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer dst.Close()

	// Copy contents from source to destination
	_, err = io.Copy(dst, src)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}

	fmt.Println("File copied successfully!")
}
