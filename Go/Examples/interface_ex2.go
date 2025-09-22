// Lab Interfaces
// Lab Exercise: Employee Management System Using Interfaces

// Problem Statement:
// You are tasked with building a simplified employee management system where different types of employees—such as full-time and part-time employees—can be managed using an interface in Go. The system will calculate the salary of each employee type.

// 1. Create an `Employee` interface that declares a method called `CalculateSalary()` which returns an `int`.

// 2. Implement two structs: `FullTimeEmployee` and `PartTimeEmployee`, each representing a type of employee with fields:
//    - `FullTimeEmployee`: `basicPay` and `bonus`
//    - `PartTimeEmployee`: `hourlyRate` and `hoursWorked`

// 3. Each struct should implement the `CalculateSalary()` method:

// - The salary for a full-time employee is calculated as `basicPay + bonus`.
// - The salary for a part-time employee is calculated as `hourlyRate * hoursWorked`.

// 4. Create a function `printSalary()` that takes an `Employee` interface as an argument and prints the employee's salary.

// 5. Write a main function that creates instances of both `FullTimeEmployee` and `PartTimeEmployee` and uses `printSalary()` to display their salaries.

// Solution:
package main

import "fmt"

// Step 1: Define the Employee interface
type Employee interface {
	CalculateSalary() int
}

// Step 2: Define FullTimeEmployee struct
type FullTimeEmployee struct {
	BasicPay int
	Bonus    int
}

// Implement the CalculateSalary method for FullTimeEmployee
func (fte FullTimeEmployee) CalculateSalary() int {
	return fte.BasicPay + fte.Bonus
}

// Step 2: Define PartTimeEmployee struct
type PartTimeEmployee struct {
	HourlyRate  int
	HoursWorked int
}

// Implement the CalculateSalary method for PartTimeEmployee
func (pte PartTimeEmployee) CalculateSalary() int {
	return pte.HourlyRate * pte.HoursWorked
}

// Step 4: Create a function to print salary of any Employee
func printSalary(e Employee) {
	fmt.Printf("Salary: %d\n", e.CalculateSalary())
}

// Step 5: Main function to create employees and print their salaries
func main() {
	fte := FullTimeEmployee{BasicPay: 50000, Bonus: 10000}
	pte := PartTimeEmployee{HourlyRate: 1000, HoursWorked: 20}

	fmt.Println("Full-Time Employee Salary:")
	printSalary(fte)

	fmt.Println("Part-Time Employee Salary:")
	printSalary(pte)
}
