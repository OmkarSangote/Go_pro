package main

import "fmt"

func main() {
	const (
		num1 = iota
		num2
		num3
	)
	fmt.Println(num1, num2, num3)

	const (
		sun = iota
		mon
		tue
		wed
		thur
		fri
		sat
	)
	fmt.Println(sun, mon, tue, wed, thur, fri, sat)
}
