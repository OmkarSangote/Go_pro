package main

import "fmt"

// Queue is nothing but a slice
type Queue struct {
	items []int
}

//Enque operation will add element into slice at end, same as psuh in stack
func (q *Queue) enqueue(i int) {
	q.items = append(q.items, i)
}

//Deque will remove elements from the end as well
//returns removed value
func (q *Queue) dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.enqueue(100)
	myQueue.enqueue(200)
	myQueue.enqueue(300)
	fmt.Println(myQueue)
	myQueue.dequeue()
	fmt.Println(myQueue)
}
