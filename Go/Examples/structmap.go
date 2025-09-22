package main

import (
	"log"
)

type info struct {

	Name string
	Age int

}

func main(){

	user:= info {
		Name:"Omkar",
		Age:23,
	}

	mymap:=make(map[string]info)

	mymap["me"] = user
	mymap["age"] = user

	log.Println(mymap["me"].Name, mymap["age"].Age)


}
