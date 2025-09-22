package main

import "fmt"

func swapper(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	a, b := 10, 20
	fmt.Printf("Before swapping value of a = %d && b = %d\n", a, b)
	swapper(&a, &b)
	fmt.Printf("After swapping value of a = %d && b = %d\n", a, b)

}
