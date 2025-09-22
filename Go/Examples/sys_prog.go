/*
   Working with OS Signals and Handling Interrupts

   Go provides a way to handle OS signals using the os/signal package.

   Signals can notify your program about various events, such as user interruptions.
*/

// Handling Interrupts
package main

import (
	"fmt"
	"os"        // for interacting with the OS e.g to exit the program
	"os/signal" // To catch and handle OS signals.
	"syscall"   // To specify which system signals we want to capture.
	"time"      // For adding delays using time.Sleep().
)

func main() {
	// Create a channel to receive OS signals
	sigs := make(chan os.Signal, 1)
	// os.Signal is the type that will hold signals, like SIGINT or SIGTERM, which represent interrupt and terminate requests.

	// The signal.Notify function tells Go to relay specific OS signals to the sigs channel. In this case, it will capture:

	//SIGINT: This signal is typically sent when you press Ctrl+C.

	//SIGTERM: This signal is sent by the operating system to request termination (e.g., in Docker or system services).
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Once the program receives these signals, they will be sent to the sigs channel.

	// Goroutine to handle signals
	// A goroutine is launched to handle the signals. Inside this goroutine:
	// The program waits for a signal to arrive via the channel sigs.
	//When a signal is received, it prints the signal (e.g., SIGINT or SIGTERM).
	//The program then gracefully exits using os.Exit(0).
	go func() {
		sig := <-sigs
		fmt.Println("Received signal:", sig)
		os.Exit(0)
	}()

	// Simulate work
	// This part of the program simulates some ongoing work by repeatedly printing "Working..." every 2 seconds.
	for {
		fmt.Println("Working...")
		time.Sleep(2 * time.Second)
	}
}
