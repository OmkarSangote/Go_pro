package main

import "log"

func main(){
	
	mymap := make(map[string] string)

	mymap["Husband"] = "Dudu"

	mymap["wife"] = "Bubu"

	log.Println(mymap["Husband"], " is Husband of ", mymap["wife"])



}