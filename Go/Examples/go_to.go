package main

import "fmt"

func main() {
	fmt.Println("Start")

	goto skip

	fmt.Println("Not executed")

skip:
	fmt.Println("Jumped here")

}
