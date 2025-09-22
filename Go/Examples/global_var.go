package main

import "fmt"

var globalVar int = 100
var g int
var g_var int = 10

func main() {
	var a, b int
	a, b = 2, 4
	g = a * b
	fmt.Println("Printing g after updating in main func:", g)
	var g_var int
	g_var = 20
	fmt.Println("Printing g_var within main func:", g_var)
	fmt.Println("Before modify global var:")
	displayGlobal()
	fmt.Println("After modify global var:")
	modifyGlobal()
	displayGlobal()

}

func displayGlobal() {
	fmt.Println(globalVar)
}

func modifyGlobal() {
	globalVar = 200
}
