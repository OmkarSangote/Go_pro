package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Print("Enter your name and age")
	fmt.Scanln(&name, &age)
	fmt.Println("Hello", name, "This is your age", age)
}
