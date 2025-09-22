package main

import "fmt"

func main() {
	const (
		num1 = 1
		num2 = 2
		num3 = 3
	)
	var var1, var2, var3 = num1, num2, num3

	fmt.Printf("%d is of type %T and has value %p\n", var1, var1, &var1)
	fmt.Printf("%d is of type %T and has value %p\n", var2, var2, &var2)
	fmt.Printf("%d is of type %T and has value %p\n", var3, var3, &var3)

}
