/*
Example 3: Appending to a File
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	// Open the file in append mode, creating it if it doesn't exist
	file, err := os.OpenFile("example2.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Append text to the file
	_, err = file.WriteString("This is appended text!\n")
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}

	fmt.Println("Data appended successfully!")

	err = os.Remove("example2.txt")

	if err != nil {
		fmt.Println("Error deleting file :", err)
		return
	}

	fmt.Println("File deleted successfully!")

	if fileExists("example2.txt") {
		fmt.Println("File exists")
	} else {
		fmt.Println("File does not exist")
	}

}

func fileExists(filename string) bool {
	_, err := os.Stat("filename")
	return !os.IsNotExist(err)
}

/*

Explanation:
1. `os.OpenFile()`: Opens a file with specific flags. In this case, `os.O_APPEND` appends data to the file,

`os.O_CREATE` creates the file if it doesn’t exist, and `os.O_WRONLY` allows writing only.

2. Permissions (`0644`): The file permissions for the new file if it's created. Here, it allows the owner to read and write, while others can only read.

3. os.Remove() is used to delete the file

*/
