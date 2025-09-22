/*
Problem Statement: Building a CLI Tool for File Operations in Go

You are tasked with creating a command-line interface (CLI) tool in Go that allows users to perform basic file operations.

The CLI should accept various command-line arguments to specify the operation the user wants to perform on a file, such as reading a file, writing to a file, and appending content to an existing file.

Key Requirements:
1. File Reading: Users should be able to specify a file to read from and print its contents.

2. File Writing: Users should be able to create a new file or overwrite an existing file with content provided through the command line.

3. File Appending: Users should be able to append new content to an existing file.

4. Command-Line Flags:
   - `-op` (string): Specifies the operation to perform. Can be "read", "write", or "append".

   - `-file` (string): Specifies the file to perform the operation on.

   - `-content` (string, optional): Specifies the content to write or append to the file (only required for "write" and "append" operations).

Example Scenarios:

- Read a file:
  `./file-cli -op read -file example.txt`
  This should print the contents of `example.txt` to the console.

- Write to a file:
  `./file-cli -op write -file newfile.txt -content "Hello, world!"`
  This should create a new file `newfile.txt` and write `"Hello, world!"` into it.

- Append to a file:
  `./file-cli -op append -file example.txt -content "Appending this text"`
  This should append `"Appending this text"` to `example.txt`.

*/

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define command-line flags
	op := flag.String("op", "", "Operation to perform: read, write, or append")
	file := flag.String("file", "", "File to perform the operation on")
	content := flag.String("content", "", "Content to write or append (required for write/append)")

	// Parse flags
	flag.Parse()

	// Validate required flags
	if *op == "" || *file == "" {
		fmt.Println("Both -op and -file flags are required")
		return
	}

	switch *op {
	case "read":
		readFile(*file)
	case "write":
		if *content == "" {
			fmt.Println("Content required for write operation")
			return
		}
		writeFile(*file, *content)
	case "append":
		if *content == "" {
			fmt.Println("Content required for append operation")
			return
		}
		appendToFile(*file, *content)
	default:
		fmt.Println("Invalid operation. Use 'read', 'write', or 'append'.")
	}
}

func readFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Printf("Contents of %s:\n%s\n", filename, string(data))
}

func writeFile(filename, content string) {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
	fmt.Printf("Successfully wrote to %s\n", filename)
}

func appendToFile(filename, content string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file for append: %v\n", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		fmt.Printf("Error appending to file: %v\n", err)
		return
	}
	fmt.Printf("Successfully appended to %s\n", filename)
}
