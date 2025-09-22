package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Input string
	text := "Medical Term: alpha-fodrin - Exp: November 9, 2003 - Phone: 408-342-5400 - SSN: 001-031240"

	// Define the SSN regex pattern
	ssnPattern := `(?:^|[^%\d])(?!000)(?:[0-6]\d{2}|7(?:[0-6]\d|7[0-2]))(?:[ -])(?!00)\d\d(?:[ -])(?!0000)\d{4}(?:$|[^\-\d])`

	// Compile the regex pattern
	re := regexp.MustCompile(ssnPattern)

	// Find the SSN in the text
	match := re.FindString(text)

	// Print the matched SSN if found
	if match != "" {
		fmt.Println("Found SSN:", match)
	} else {
		fmt.Println("No SSN found in the text.")
	}
}

