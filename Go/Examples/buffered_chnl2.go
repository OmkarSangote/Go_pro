/*
Advanced Example: Worker Pool Using Buffered Channels

In this example, we create a worker pool where a limited number of goroutines process tasks concurrently.

The tasks are distributed through a buffered channel.
*/
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Simulate time taken to process job
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2 // Send result to the results channel
	}
}

func main() {
	numJobs := 5 // creating a buffer
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start 3 worker goroutines  (3 workers running concurrently)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs to the jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Close the jobs channel after sending all jobs

	// Collect all results
	for a := 1; a <= numJobs; a++ {
		fmt.Printf("Result: %d\n", <-results)
	}
}

/*

Explanation:

Worker pool: We have 3 workers running concurrently, and each worker processes tasks (jobs) from the jobs channel.

Buffered channel: The jobs and results channels are buffered to allow non-blocking sends and receives.

The workers block on the jobs channel until a job is available, and the main goroutine blocks on the results channel until results are available.

*/
