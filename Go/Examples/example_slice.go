package main

import "fmt"

func main() {
	num := []int{10, 20, 30, 40, 50}
	fmt.Println(num)

	fmt.Println()
	var fruit_list = []string{"Apple", "banana"}
	fmt.Println(fruit_list)

	scores := make([]int, 4, 8)
	scores[0] = 10
	scores[1] = 20
	scores[2] = 30
	scores[3] = 40

	scores = append(scores, 50)
	scores = append(scores, 60)
	scores = append(scores, 70)
	scores = append(scores, 80)
	scores = append(scores, 90)
	fmt.Println(scores)
	fmt.Printf("Capacity and length of scores is: Cap = %d && Length = %d\n", cap(scores), len(scores))
}
