package main

import (
	"log"
	"sort"
)

func main(){

	var myslice[] string

	myslice = append(myslice, "Dudu")
	myslice = append(myslice, "bubu")

	log.Println(myslice)

	numbers:= []int{3,1,2,4,5,6,7,8,9}

	sort.Ints(numbers)

	log.Println(numbers)

	log.Println(numbers[0:3])

}

