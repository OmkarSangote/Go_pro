package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// Execute the wmic command to get disk space information
	out, err := exec.Command("df", "-h").Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the output
	fmt.Println(strings.TrimSpace(string(out)))
}
