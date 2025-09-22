package main

import (
	"fmt"
	"math"
)

func main() {
	quotient, err := divide(10, 5)
	if err == nil {
		fmt.Println(quotient)
	} else {
		fmt.Println("Error", err)
	}

	add := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(add)

	//Anonyms function assigned to variable
	square := func(x int) int {
		return x * x
	}
	fmt.Println("5 X 5 =", square(5))

	// call the outer function && outer func will always be assigned to a value
	message := greet()
	// call the inner function
	fmt.Println(message())

	sqRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println("Square root of 225 is:", sqRoot(225))

	myArr := [5]int{1, 2, 3, 4, 5}
	modifyArray(myArr)

	fmt.Println("Outside func but inside main func :", myArr)

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Call the function
	modifySlice(mySlice)

	// The original slice is modified
	fmt.Println("Outside function but inside main func:", mySlice)

	myMap := map[string]int{"age": 25, "year": 2024}

	modifyMap(myMap)

	fmt.Println("Outside function but inside main func:", myMap)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by 0")
	}
	return a / b, nil
}

func sum(nums ...int) int {
	total := 0
	for _, i := range nums {
		total += i
	}
	return total
}

// nested function
// outer function
func greet() func() string {
	//variable defined outside the innner function
	name := "Omkar"
	//return a nested anonymous func
	//this is a closure
	return func() string {
		name = "Hi " + name
		return name
	}
}

func modifyArray(arr [5]int) {
	arr[0] = 100
	fmt.Println("Inside Function", arr)
}

func modifySlice(slc []int) {
	slc[0] = 100
	fmt.Println("Inside function:", slc)
}

func modifyMap(m map[string]int) {
	m["age"] = 30
	fmt.Println("Inside function:", m)
}
