package main

import "log"
import "fmt"

type details struct {
	name string
	age int
}

func (x *details) ptstruct() string {
	fmt.Println("name:", x.name, "age:", x.age)
	return x.name
}


func main() {

	data := details{
		name: "Omkar",
		age: 23,
	}

	log.Println(data.name)
}


