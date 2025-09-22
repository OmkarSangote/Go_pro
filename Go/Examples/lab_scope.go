package main

import "fmt"

// Global variable
var y int = 30

func main() {
	y = 50
	fmt.Println(y)
	localVariable()
	fmt.Println(y)
}

func localVariable() {
	y := 40
	fmt.Println(y)
}
