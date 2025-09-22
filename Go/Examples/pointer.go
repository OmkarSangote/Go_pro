package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Println("value of a", a)
	fmt.Println("value of b", b)
	*b = 20
	fmt.Println("new value of a", a)
	a = 30
	fmt.Println("New value of b", b)

	var color string

	color = "Red"

	fmt.Println("The colour is", color)

	pointer(&color)

	fmt.Println("After funcion call", color)

}

func pointer(s *string) {

	fmt.Println("The value of string", s)

	new := "Green"

	*s = new

}
