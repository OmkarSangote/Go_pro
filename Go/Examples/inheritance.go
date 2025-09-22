package main

import "fmt"

type Vehicle struct {
	Make string
	Model string
}

func (v Vehicle) start() {
	fmt.Println(v.Make, v.Model, "is starting....")
}

type car struct {
	Vehicle
	NumDoors int
}

