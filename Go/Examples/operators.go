package main

import "fmt"

func main() {
	a := true
	b := false
	x := 1
	y := 0
	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(x ^ y)
	fmt.Println(!a, !b)

}
