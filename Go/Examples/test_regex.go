package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Define the updated regex for SSN
	//ssnRegex := `(?:^|[^%\d])(?!000)(?:[0-6]\d{2}|7(?:[0-6]\d|7[0-2]))(?:[ -]*)(?!00)\d\d(?:[ -]*)(?!0000)\d{4}(?:$|[^\-\d])`
	//ssnRegex := `(?:^|[^%\d])([0-6]\d{2}|7(?:[0-6]\d|7[0-2]))(?:[ -]*)(?!00)\d\d(?:[ -]*)(?!0000)\d{4}(?:$|[^\-\d])`
	ssnRegex := `(?:^|[^%\d])([0-6]\d{2}|7(?:[0-6]\d|7[0-2]))(?:[ -]*)(\d{2})(?:[ -]*)(\d{4})(?:$|[^\-\d])`
	// Compile the regex
	re, err := regexp.Compile(ssnRegex)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	// Test cases
	testSSNs := []string{
		"001-03-1240",     // Standard format
		"001 03 1240",     // Spaces instead of hyphens
		"001- 03- 1240",   // Extra space after hyphen
		"001-03 1240",     // Mixed hyphens and spaces
		"001   03   1240", // Multiple spaces
		"000   03   1240",
		"001   00   1240",
		"001   03   0000",
	}

	// Check each SSN against the regex
	for _, ssn := range testSSNs {
		if re.MatchString(ssn) {
			fmt.Printf("%s: SSN match found!\n", ssn)
		} else {
			fmt.Printf("%s: SSN not matched.\n", ssn)
		}
	}
}
