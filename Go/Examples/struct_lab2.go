/*
Create a Go program to manage information about students in a company training program. Each student has the following details:
 
Name (string)
Age (int)
Email (string)
BatchID (int)
Your task is to:
 
Define a Student struct with the fields mentioned above.
 
Write a function to create a new student (returning a pointer to the Student struct).
 
Write a function to display the student's details.
 
Create a slice to store a list of students.
 
Implement a function to add a student to the list.
 
Implement a function to display all students in the list.
*/

package main

import (
	"fmt"
)

// Student struct to hold the details of each student
type Student struct {
	Name   string
	Age    int
	Email  string
	BatchID int
}

// Function to create a new student and return a pointer to the Student struct
func NewStudent(name string, age int, email string, batchID int) *Student {
	return &Student{
		Name:   name,
		Age:    age,
		Email:  email,
		BatchID: batchID,
	}
}

// Method to display the student's details
func (s *Student) Display() {
	fmt.Printf("Name: %s, Age: %d, Email: %s, BatchID: %d\n", s.Name, s.Age, s.Email, s.BatchID)
}

// Slice to store a list of students
var studentList []*Student

// Function to add a student to the list
func AddStudent(s *Student) {
	studentList = append(studentList, s)
}

// Function to display all students in the list
func DisplayAllStudents() {
	for _, student := range studentList {
		student.Display()
	}
}

func main() {
	// Create new students
	student1 := NewStudent("Alice Smith", 22, "alice@example.com", 101)
	student2 := NewStudent("Bob Johnson", 25, "bob@example.com", 102)

	// Add students to the list
	AddStudent(student1)
	AddStudent(student2)

	// Display all students
	fmt.Println("List of Students:")
	DisplayAllStudents()
}

