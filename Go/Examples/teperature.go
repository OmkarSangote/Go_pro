package main

import "fmt"

func main() {
	var cel_temp, fahren float64
	fmt.Println("Enter the temeprature in Celcius")
	fmt.Scanln(&cel_temp)
	fahren = ((cel_temp * 9) / 5) + 32
	fmt.Println("temp in fahreheit:", fahren)
}
