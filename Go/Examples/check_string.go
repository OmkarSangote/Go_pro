package main

import "fmt"

func main() {
	var usr_string string
	fmt.Println("Enter the string")
	fmt.Scanln(&usr_string)
	length := len(usr_string)
	switch {
	case length%2 == 0:
		fmt.Println("Entered string is of even length")

	case length%2 != 0:
		fmt.Println("Entered string is of odd length")
	}
}
