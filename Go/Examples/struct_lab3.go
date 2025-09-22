// Lab Exercise: Embedded Structs in Go
// Problem Statement:
// You are tasked with creating a Go program that manages information about employees working in different departments of a company. Each employee has common details like:

// Name (string)
// Age (int)
// EmployeeID (int)
// Additionally, employees belong to a department, and each department has:

// Department Name (string)
// Manager Name (string)
// Your task is to:

// Define an Employee struct that contains the common fields.
// Define a Department struct that contains department-specific details and embeds the Employee struct.
// Write a function to create a new department with an employee.
// Write a method to display both the employee and department details.
// Create a list of departments, and write a function to display all the department and employee information.

package main

import (
	"fmt"
)

// Step 1: Define the Employee struct
type Employee struct {
	Name       string
	Age        int
	EmployeeID int
}

// Step 2: Define the Department struct that embeds Employee
type Department struct {
	Employee
	DepartmentName string
	ManagerName    string
}

// Step 3: Function to create a new department with an employee
func NewDepartment(name string, age int, empID int, deptName string, managerName string) Department { // Return value instead of pointer
	return Department{
		Employee:       Employee{Name: name, Age: age, EmployeeID: empID},
		DepartmentName: deptName,
		ManagerName:    managerName,
	}
}

// Step 4: Method to display both employee and department details
func (d Department) Display() { // Use value receiver
	fmt.Printf("Employee Name: %s\nAge: %d\nEmployee ID: %d\nDepartment: %s\nManager: %s\n", d.Name, d.Age, d.EmployeeID, d.DepartmentName, d.ManagerName)
	fmt.Println("-------------")
}

// Step 5: List of departments and function to display all
var departments []Department

func AddDepartment(dept Department) {
	departments = append(departments, dept)
}

func DisplayAllDepartments() {
	for _, dept := range departments {
		dept.Display()
	}
}

func main() {
	// Create and add departments
	dept1 := NewDepartment("Alice", 30, 1001, "IT", "John Doe")
	dept2 := NewDepartment("Bob", 25, 1002, "HR", "Jane Smith")

	AddDepartment(dept1) // Pass value instead of pointer
	AddDepartment(dept2) // Pass value instead of pointer

	// Display all departments with employee details
	fmt.Println("List of Departments and Employees:")
	DisplayAllDepartments()
}
