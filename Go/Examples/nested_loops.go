package main

import "fmt"

func main() {
	//outer loop
	for i := 1; i <= 3; i++ {
		fmt.Printf("Outer loop iteration : %d\n", i)

		//Inner loop
		for j := 1; j <= 2; j++ {
			fmt.Printf("Inner loop iteration : %d\n", j)
		}
	}
}
