package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the number for which table is needed")
	fmt.Scan(&num)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", num, i, num*i)
	}
}
