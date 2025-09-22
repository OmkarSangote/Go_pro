package main

import (
	"fmt"
	"os"
	"plugin"
)

type Calculator interface {
	Calculate(num ...int)
}

func main() {
	calculatorPlugin, err := plugin.Open("./calculator/calculator.so") // Ensure this path is correct
	if err != nil {
		fmt.Println("error while opening shared object file:", err)
		os.Exit(1)
	}

	// Load the NewCalculator function
	symNewCalculator, err := calculatorPlugin.Lookup("NewCalculator")
	if err != nil {
		fmt.Println("error while lookup:", err)
		os.Exit(1)
	}

	// Assert that it's a function with the correct signature
	newCalculatorFunc, ok := symNewCalculator.(func() Calculator)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	// Create a calculator instance
	calculator := newCalculatorFunc()

	// Use the calculator
	calculator.Calculate(3, 4)
}
