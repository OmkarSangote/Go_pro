package main

import "fmt"

type Stack struct {
	items []int
}

// Push operation will add value at end
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

//POP operation will remove value from begining
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}
