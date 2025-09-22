package main

import (
	"log"
	"time"
)

type User struct {
	name string
	age int
	phone int
	date time.Time
}

func main(){
	details:= User {
		name:"Omkar",
		age:23,
		phone:9513083254,
		}
		log.Println(details.name,details.age,details.phone)
}

func 