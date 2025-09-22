package main

import "fmt"

func multiply(a int, b int) (result int) {
	result = a * b
}

func main() {
	fmt.Println(multiply(4, 5))
}
