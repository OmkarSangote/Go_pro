package main

import (
	"fmt"
)

type marriage interface {
	bride() string
	groom() string

}
type couples struct {

	dudu string
	bubu string

}


func main() {

	marry := couples{
		dudu: "Omkar",
		bubu: "Mahima",
	}

	printinfo(marry)


}

func printinfo(a couples){
	fmt.Println ("The bride is",a.bride(),"and the groom is",a.groom())
}

func(c couples) bride() string{
	return "Me"
} 

func (c couples) groom() string{
	return "Wife"
}