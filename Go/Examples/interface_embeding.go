package main

import "fmt"

// Inheritance throgh interface embedding
// interface
type Reader interface {
	Read() string
}

// Embed the reader interface in another interface
type Writer interface {
	Write(data string)
	Reader // Embedding the reader interface
}

// Implement the Writer interface
type mystruct struct{}

func (m mystruct) Read() string {
	return "Reading data"
}
func (m mystruct) Write(data string) {
	fmt.Println("Writing ****", data)
}
func main() {
	var w Writer = mystruct{}
	fmt.Println(w.Read()) // Accessing the embedded method
	w.Write("Hello, world !")
}
