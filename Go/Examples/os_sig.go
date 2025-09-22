package main


import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"
)


func main() {
    // Create a channel to receive OS signals
    signalChan := make(chan os.Signal, 1)


    // Notify on SINGINT (ctrl+c)
    signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)


    // create a channel to signal programm termination
    done := make(chan bool, 1)


    go func() {
        sig := <-signalChan
        fmt.Printf("\nReceived Signal: %s\n", sig)
        fmt.Println("Performing cleanup....")
        // simulate cleanup work
        time.Sleep(2 * time.Second)
        fmt.Println("Cleanup complete. shutting down gracefully.")
        done <- true


    }()


    fmt.Println("Server is running . Press Ctrl+C to exit")
    <-done // read
    fmt.Println("Server stopped")


}
