package main

import "fmt"

type Printer interface {
	Print() string
}

func main() {
	var p Printer
	if p == nil {
		fmt.Println("Printer interface is nil")
	} else {
		fmt.Println(p.Print())
	}

}
