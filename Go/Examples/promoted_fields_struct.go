package main

import "fmt"

type details struct {
	Name  string
	Email string
	Address
}

type Address struct {
	Street string
	City   string
	State  string
	Zip    int
}

func main() {
	details1 := details{
		Name:  "Omkar",
		Email: "osangote@cuda.com",
		Address: Address{
			Street: "Bakers street",
			City:   "Snk",
			State:  "KA",
			Zip:    591313,
		},
	}

	// Accessing fields of outer struct
	fmt.Println("Name :", details1.Name)

	// Accessing fileds of inner or nested struct directly
	// Accessing promoted fields directly
	fmt.Println(details1.Zip)

}
