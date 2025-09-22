package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("         Embedding Interfaces ")
	fmt.Println(strings.Repeat("*", 60))

	// Creating Instances of the struct Employee
	emp1 := Employee{
		name:         "Alice",
		age:          25,
		basic_pay:    20000,
		pf:           2000,
		total_leaves: 15,
		leaves_taken: 5,
	}

	// Create a variable of type EmployeeOperations --> Embedded Interface
	var emp_op EmployeeOperations
	emp_op = emp1

	emp_op.display_salary()
	emp_op.leaves_left()
}

// Define interfaces
type SalaryCalculator interface {
	display_salary()
}

type LeaveCalculator interface {
	leaves_left()
}

// Embedding the interfaces
type EmployeeOperations interface {
	SalaryCalculator
	LeaveCalculator
}

// Define a struct Employee
type Employee struct {
	name         string
	age          int
	basic_pay    int
	pf           int
	total_leaves int
	leaves_taken int
}

// Define methods of Employee struct
func (e Employee) display_salary() {
	fmt.Printf("Total salary of %s is %d\n", e.name, (e.basic_pay + e.pf))
}

func (e Employee) leaves_left() {
	fmt.Println("Remaining Leaves is:", e.total_leaves-e.leaves_taken)
}
