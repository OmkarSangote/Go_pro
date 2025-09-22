/*

    Fan-In and Fan-Out in Go (Concurrency Patterns)

--> Fan-In and Fan-Out are concurrency patterns in Go often used in pipeline designs. They help manage how goroutines communicate through channels.

Fan-Out: A single input is distributed to multiple workers (goroutines), each processing the data concurrently.

Fan-In: Merges multiple channels into one, allowing you to gather data from several workers into a single channel.
*/

package main

import (
	"fmt"
)

// worker(id int, ch chan int): This function simulates a worker goroutine that sends 5 values to the channel ch. The values sent are the product of id and i.
func worker(id int, ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- id * i // Sending values to channel
	}
	close(ch) // Close the channel when done
}

func fanIn(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		for {
			// A select statement is used to listen to both channels (ch1 and ch2).
			select {

			case v, ok := <-ch1:
				if ok {
					// // Whenever a value is available on one of the channels, it forwards the value to the newly created ch.

					// This allows us to gather values from multiple channels and pass them into a single channel, achieving Fan-In.
					ch <- v
				} else {
					ch1 = nil // Avoid blocking when ch1 is closed
				}
			case v, ok := <-ch2:
				if ok {
					ch <- v
				} else {
					ch2 = nil // Avoid blocking when ch2 is closed
				}
			}
			if ch1 == nil && ch2 == nil {
				close(ch) // Close the output channel when both inputs are closed
				return
			}
		}
	}()
	return ch
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Start worker goroutines
	// The worker goroutines simulate Fan-Out by distributing work (sending values) concurrently across multiple channels (ch1 and ch2).
	go worker(1, ch1)
	go worker(2, ch2)

	// Use a single fanIn function to merge ch1 and ch2
	merged := fanIn(ch1, ch2)

	// Consume values from the merged channel
	for v := range merged {
		fmt.Println(v) // Receiving values from fan-in channel
	}
}

/*
Worker Function:

func worker(id int, ch chan int) {
    for i := 0; i < 5; i++ {
        ch <- id * i // Sending values to the channel
    }
}

--> worker(id int, ch chan int): This function simulates a worker goroutine that sends 5 values to the channel ch. The values sent are the product of id and i.

Example: For id = 1, it sends 0, 1, 2, 3, 4 to the channel.

For id = 2, it sends 0, 2, 4, 6, 8 to the channel.

Fan-In Function:
func fanIn(ch1, ch2 <-chan int) <-chan int {
    ch := make(chan int) // Create a new channel
    go func() {
        for {
            select {
            case v := <-ch1:
                ch <- v // Forward value from ch1
            case v := <-ch2:
                ch <- v // Forward value from ch2
            }
        }
    }()
    return ch // Return the merged channel
}
fanIn(ch1, ch2): This function takes two input channels (ch1 and ch2) and merges them into a single output channel.

Inside a goroutine:
--> A select statement is used to listen to both channels (ch1 and ch2).

--> Whenever a value is available on one of the channels, it forwards the value to the newly created ch.

--> This allows us to gather values from multiple channels and pass them into a single channel, achieving Fan-In.

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    // Start two worker goroutines
    go worker(1, ch1)
    go worker(2, ch2)

    // Receive values from fan-in channel and print them
    for i := 0; i < 10; i++ {
        fmt.Println(<-fanIn(ch1, ch2)) // Receiving values from fan-in channel
    }
}
Create Channels: ch1 and ch2 are created for the two workers.

Start Workers:

1. Two workers are started: one with id = 1, the other with id = 2.

2. Fan-In:
The fanIn(ch1, ch2) function is called to merge the two channels into one.

Receiving Values:
The program receives and prints 10 values (5 from each worker) using the merged fanIn channel.

                        How This Works:
Fan-Out: The worker goroutines simulate Fan-Out by distributing work (sending values) concurrently across multiple channels (ch1 and ch2).

Fan-In: The fanIn function merges the output of these channels into a single channel, allowing you to collect results from multiple goroutines in one place.

Example Output:
The output is printed from the merged fanIn channel. The order of output may vary depending on how the goroutines are scheduled by the Go runtime:

Copy code
0
0
1
2
2
4
3
6
4
8
Here, you can see that values are received from both ch1 and ch2 in an interleaved manner. The worker with id = 1 sends 0, 1, 2, 3, 4, while the worker with id = 2 sends 0, 2, 4, 6, 8.

Visualization:
Fan-Out:

Worker 1 -> ch1
Worker 2 -> ch2
Two workers running concurrently, each sending values to separate channels (ch1 and ch2).

Fan-In:

fanIn(ch1, ch2) merges the output from both channels into a single channel, so the main function can receive values from either worker without caring about which channel the data comes from.

Why Use Fan-In and Fan-Out?
Fan-Out allows you to run multiple tasks concurrently, making your program more efficient by distributing work across goroutines.

Fan-In is useful when you need to collect results from multiple goroutines and handle them in one place.


*/
