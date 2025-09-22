package main

import "fmt"

// Define the Creature interface
type Creature interface {
	Walk()
	Sleep()
}

// Define the Bird interface that embeds Creature and adds Fly
type Bird interface {
	Creature
	Fly()
}

// Define the Fish interface that embeds Creature and adds Swim
type Fish interface {
	Creature
	Swim()
}

// Define the Sparrow struct
type Sparrow struct {
	name string
}

// Define the Shark struct
type Shark struct {
	species string
}

// Implement methods for Sparrow
func (s Sparrow) Walk() {
	fmt.Printf("%s is walking\n", s.name)
}

func (s Sparrow) Sleep() {
	fmt.Printf("%s is sleeping\n", s.name)
}

func (s Sparrow) Fly() {
	fmt.Printf("%s is flying\n", s.name)
}

// Implement methods for Shark
func (s Shark) Walk() {
	fmt.Printf("%s cannot walk, but swims well\n", s.species)
}

func (s Shark) Sleep() {
	fmt.Printf("%s is sleeping\n", s.species)
}

func (s Shark) Swim() {
	fmt.Printf("%s is swimming\n", s.species)
}

func main() {
	// Create instances of Sparrow and Shark
	sparrow := Sparrow{name: "Jack"}
	shark := Shark{species: "Great White"}

	// Use the Bird interface for controlling Sparrow
	var b Bird = sparrow
	b.Walk()
	b.Fly()
	b.Sleep()

	fmt.Println()

	// Use the Fish interface for controlling Shark
	var f Fish = shark
	f.Walk() // Special case for shark: "cannot walk"
	f.Swim()
	f.Sleep()
}
