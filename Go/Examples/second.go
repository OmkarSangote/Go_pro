package main

import "fmt"

type user struct {
	name string
	age  int
}

func (u user) isAdult() bool {
	return u.age >= 18
}

func main() {
	omkar := user{"Omkar", 14}
	fmt.Println(omkar.isAdult())
}
