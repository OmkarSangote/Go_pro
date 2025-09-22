package main


import (
    "fmt"
    "os"
)


// func main() {
//  // open a file
//  fd, err := syscall.Open("example.txt", syscall.O_RDWR, 0666)


//  if err != nil {
//      fmt.Printf("Error opening file: %v\n", err)
//      return
//  }


//  fmt.Printf("File Descriptor: %d\n", fd)


//  // write to the file
//  data := []byte("Hello , syscall!\n")
//  _, err = syscall.Write(fd, data)


//  if err != nil {
//      fmt.Printf("Error writing to file %v\n", err)


//  }


//  // close the file
//  err = syscall.Close(fd)
//  if err != nil {
//      fmt.Printf("Error closing to file %v\n", err)


//  }


// }


// -------Handle OS signal-----------------


// func main() {
//  sigChan := make(chan os.Signal, 1)


//  signal.Notify(sigChan, syscall.SIGINT)


//  fmt.Println("waiting for signal......")


//  // block until a signal is received
//  sig := <-sigChan
//  fmt.Printf("Recevied signal: %v\n", sig)


//  // Perform cleanup or shutdown
//  fmt.Println("Exiting gracefully")


// }


// ------------Custom System Calls


func main() {
    // invoke the getpid system call


    //   pid, _, _ := syscall.Syscall(syscall.SYS_GETPID,0,0,0)


    fmt.Printf("Current process Id: %d\n", os.Getpid())
}
