/*
--> Using a ticker with channels in Go is a common pattern to trigger actions at regular intervals.

--> You can use channels to combine the behavior of a ticker with other concurrent operations or message passing.

--> The ticker’s channel can be integrated with select statements to coordinate the timing of actions between goroutines.

	Example: Ticker with Channels

In this example, we use a ticker to trigger the sending of messages through a channel every second, simulating periodic events like sensor readings, logging, or timed API calls.
*/
package main

import (
	"fmt"
	"time"
)

func sendData(ch chan<- string) {
	ticker := time.NewTicker(1 * time.Second) // Create a ticker that ticks every second
	defer ticker.Stop()                       // Stop the ticker once we're done

	for i := 1; i <= 5; i++ {
		// The ticker sends the current time on its channel ticker.C every second.
		<-ticker.C                         // Wait for the next tick
		ch <- fmt.Sprintf("Message %d", i) // Send data through the channel after each tick
	}
	close(ch) // Close the channel when done
}

func main() {
	ch := make(chan string) // Create a channel

	// and after each tick, we send a message through the channel (ch <- "Message X")
	go sendData(ch) // Start the goroutine that sends data to the channel every second

	for msg := range ch { // Receive messages from the channel
		fmt.Println("Received:", msg) // Print the received messages
	}

	fmt.Println("All messages received.")
}

/*

                            Explanation of the Code

1. Creating a Ticker

--> We use time.NewTicker(1 * time.Second) to create a ticker that ticks every second.

--> The ticker sends the current time on its channel ticker.C every second.


2. Sending Data on Ticker Ticks
--> The function sendData runs as a separate goroutine.
Inside the goroutine, we use a loop to wait for each tick (<-ticker.C), and after each tick, we send a message through the channel (ch <- "Message X").

--> This simulates periodic sending of messages at regular intervals.


3. Receiving Data

--> In the main function, we use range to read from the channel.
--> The range loop reads messages from the channel until the channel is closed.

--> In this case, the sendData function closes the channel once it has sent 5 messages.

4. Ticker’s Role in Rate-Limiting

--> The ticker is limiting how fast the messages are sent into the channel.

--> If we didn't have the ticker, all messages would be sent instantly.

--> The ticker ensures that one message is sent every second by causing the goroutine to wait for a tick before sending the next message.

Output:

Received: Message 1
Received: Message 2
Received: Message 3
Received: Message 4
Received: Message 5
All messages received.
Each message is received 1 second apart.

*/
