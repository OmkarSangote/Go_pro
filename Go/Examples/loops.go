package main

import "log"

func main(){
	for i := 0; i <= 5; i++{
		log.Println(i)
	}

	animals := []string{"Cat","dog","fish","horse"}

	for _, animal := range(animals){
		log.Println(animal)
	}


	mypets := make(map[string]string)

	mypets["Dog"] = "Simba"
	mypets["cat"] = "persian"
	mypets["horse"] = "Chetak"

	for i, pet := range mypets{
		log.Println(i,pet)
	}

	type info struct {
		Name string
		Age int
		Email string
	}

	var users[] info

	users = []info{{"bublicious", 23, "bu@bu.com"},{"Omkar", 23, "om@kar.com"}}
	users = append(users, info{"bubu", 21, "bubu@dud.com"})
	users = append(users, info{"Arnu", 2, "arnu@chotu.com"})

	for _, l := range(users){
		log.Println(l.Name,l.Age,l.Email)
	}
}