package main

import "fmt"

func main() {
	var a, b = 5, 6
	fmt.Println("before swap var a:", a)
	fmt.Println("before swap var b:", b)
	b, a = 5, 6
	fmt.Println("after swap var a:", a)
	fmt.Println("after swap var b:", b)

}
