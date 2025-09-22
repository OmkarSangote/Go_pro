package main

import "log"

func main(){

	num := 80

	condition := false

	if num> 99 && condition{
		log.Println("True case")
	}else if num<100 && !condition{
		log.Println("False case")
	}else{
		log.Println("Not in criteria")
	}

	animal:="fish"

	switch animal{
	case "cat":
			log.Println("Cat")
	case "dog":
			log.Println("Dog")
	default:
			log.Println("Other animal")
	}
}