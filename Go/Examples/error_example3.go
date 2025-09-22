// Zero division error
package main

import (
	"errors"
	"fmt"
)

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return x / y, nil
}

func main() {
	var numerator, denominator int
	fmt.Println("Enter the numerator")
	fmt.Scanf("%d", &numerator)
	fmt.Println("Enter the denominator")
	fmt.Scanf("%d", &denominator)
	result, err := divide(numerator, denominator)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
