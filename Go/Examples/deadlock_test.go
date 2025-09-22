package main

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1 // Goroutine 1 sends data to ch1
		<-ch2    // Goroutine 1 waits for data from ch2
	}()

	go func() {
		ch2 <- 2 // Goroutine 2 sends data to ch2
		<-ch1    // Goroutine 2 waits for data from ch1
	}()

	// Both goroutines will be blocked and the program deadlocks
}
