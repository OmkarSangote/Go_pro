package main

import "fmt"

// Global variable
var z int = 7

func main() {
	z := 9
	fmt.Println(z)
	updateGlobal()
	fmt.Println(z)
}

func updateGlobal() {
	z = 11
}
