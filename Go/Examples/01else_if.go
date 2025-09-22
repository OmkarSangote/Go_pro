package main

import "fmt"

func main() {
	score := 85
	if score > 90 {
		fmt.Println("A+ grade")
	} else if score > 80 {
		fmt.Println("B grade")
	} else if score > 70 {
		fmt.Println("C grade")
	} else {
		fmt.Println("D grade")
	}
}
