package main

import "fmt"

// func main() {
// 	ch := make(chan int, 3) // create buff channel with capacity 3

// 	ch <- 1 //send values into chnl (non blocking until buff has space)
// 	ch <- 2
// 	ch <- 3

// 	fmt.Println(<-ch) //Rx and print values from chnl
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

//

func sum(numbers []int, ch chan int) {
	total := 0
	for _, num := range numbers {
		total += num
	}
	ch <- total // Send the result to the channel
}

func main() {
	numbers1 := []int{1, 2, 3, 4, 5}
	numbers2 := []int{6, 7, 8, 9, 10}

	ch := make(chan int) // Create a channel

	// Start two goroutines to calculate the sum of two slices
	go sum(numbers1, ch) // Worker1
	go sum(numbers2, ch) // Worker2

	sum1 := <-ch // Receive the first result
	sum2 := <-ch // Receive the second result

	fmt.Println("Sum of numbers1:", sum1)
	fmt.Println("Sum of numbers2:", sum2)
}
