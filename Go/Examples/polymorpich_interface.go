/*
Problem Statement:

You are tasked with developing a program to calculate the total monthly expense of a company for different types of workers. The company has three types of employees:

1. Full-time employees: They have a fixed monthly base salary.

2. Part-time employees: They are paid based on an hourly rate and the number of hours worked in a month.

3. Contractors: They are paid a fixed contract amount for their work.

Your task is to:
1. Define an interface `Employee` that requires a `GetSalary()` method for all employee types.

2. Create separate structs for:
   - FullTimeEmployee, which has a `baseSalary`.
   - PartTimeEmployee, which has an `hourlyRate` and `hoursWorked`.
   - Contractor, which has a `contractAmount`.

3. Implement the `GetSalary()` method for each employee type based on how they are paid:
   - Full-time employees return their base salary.
   - Part-time employees return their total earnings by multiplying the hourly rate by the hours worked.
   - Contractors return their contract amount.

4. Write a function `CalculateTotalExpense()` that takes a slice of `Employee` types and calculates the total salary expenses for all employees.

5. In the `main` function, create instances of all employee types and calculate the company's total expense.



Expected Output:
The program should output the total monthly expense incurred by the company for all types of employees. For example:

Total Expense of the Company: $8200.00


This exercise demonstrates the use of interfaces and polymorphism in Go, allowing different employee types to be handled uniformly based on their common behavior (`GetSalary()`), while hiding the implementation details specific to each employee type.
*/

package main

import "fmt"

// Define an Employee interface
type Employee interface {
	GetSalary() float64 // Every employee type must implement GetSalary
}

// FullTimeEmployee struct
type FullTimeEmployee struct {
	baseSalary float64
}

// PartTimeEmployee struct
type PartTimeEmployee struct {
	hourlyRate  float64
	hoursWorked float64
}

// Contractor struct
type Contractor struct {
	contractAmount float64
}

// Implement GetSalary for FullTimeEmployee
func (fte FullTimeEmployee) GetSalary() float64 {
	return fte.baseSalary
}

// Implement GetSalary for PartTimeEmployee
func (pte PartTimeEmployee) GetSalary() float64 {
	return pte.hourlyRate * pte.hoursWorked
}

// Implement GetSalary for Contractor
func (c Contractor) GetSalary() float64 {
	return c.contractAmount
}

// Function to calculate total expenses
func CalculateTotalExpense(employees []Employee) float64 {
	total := 0.0
	for _, employee := range employees {
		total += employee.GetSalary() // Polymorphism in action
	}
	return total
}

func main() {
	// Create different types of employees
	fte := FullTimeEmployee{baseSalary: 5000}
	pte := PartTimeEmployee{hourlyRate: 20, hoursWorked: 160}
	contractor := Contractor{contractAmount: 3000}

	// Create a slice of Employee interface
	employees := []Employee{fte, pte, contractor}

	// Calculate total expense
	totalExpense := CalculateTotalExpense(employees)

	fmt.Printf("Total Expense of the Company: $%.2f\n", totalExpense)
}
