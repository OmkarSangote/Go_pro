package main

import (
	"fmt"
	"strings"
)

func isalphanumeric(s string) bool {
	if len(s) != 1 {
		return false
	}
	r := s[0]
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func ispallaindrome(s []string) bool {
	l := 0
	r := len(s) - 1

	for l < r {
		for l < r && !isalphanumeric(s[l]) {
			l += 1
		}
		for r > l && !isalphanumeric(s[r]) {
			r -= 1
		}
		if strings.ToLower(s[l]) != strings.ToLower(s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	input := []string{"A", " ", "b", "a"} // represents "A b a"
	fmt.Println(ispallaindrome(input))    // Output: true
}
