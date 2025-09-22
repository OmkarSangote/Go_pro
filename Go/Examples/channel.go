package main

import (
	"log"
	"math/rand"
)

const number = 20
func main(){
	intChan := make (chan int)
	defer close(intChan)

	go calculate(intChan)

	x := <-intChan
	log.Println(x)
}


func calculate(intChan chan int){
	num := randomnum(number)
	intChan <- num
}
	


func randomnum(n int) int{
	value := rand.Intn(n)
	return value

}
