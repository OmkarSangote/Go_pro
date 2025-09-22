package main

import "fmt"

func sensors(temp chan int) {
	temp <- 75
}

func main() {
	temp := make(chan int)

	go sensors(temp)

	temperature := <-temp

	fmt.Println(temperature)

}
