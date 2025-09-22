package main

import "fmt"

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (c Circle) Area() {
	fmt.Printf("Area of Circle : %.2f\n", 3.142*c.Radius)
}

func (c Circle) Dimension() {
	fmt.Printf("Radius of Circle : %.2f\n", c.Radius)
}

func (r Rectangle) Area() {
	fmt.Printf("Area of Rectangle : %.2f\n", r.Length*r.Width)
}

func (r Rectangle) Dimension() {
	fmt.Printf("Length of Rectangle : %f and heigth : %f\n", r.Length, r.Width)
}

type shape interface {
	Area()
	Dimension()
}

func main() {
	// Create a Circle instance
	circle := Circle{Radius: 5}
	// Create a Rectangle instance
	rectangle := Rectangle{Length: 4, Width: 3}

	// Declare a slice of shapes
	shapes := []shape{circle, rectangle}

	// Iterate through each shape and call Area and Dimension methods
	for _, s := range shapes {
		s.Area()
		s.Dimension()
	}
}
