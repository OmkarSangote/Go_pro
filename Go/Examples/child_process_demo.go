package main


import (
    "fmt"
    "os/exec"
	"bytes"
)


// func main() {
//     // Define the command to execute (e.g. listing files)


//      cmd := exec.Command("ls", "-l", "/Users/osangote/Desktop")
//     //cmd := exec.Command("cmd", "/C", "dir")


//     output, err := cmd.Output()


//     if err != nil {
//         fmt.Printf("Error: %v\n", err)
//         return
//     }


//     // Print the output
//     fmt.Println("Command Output")
//     fmt.Println(string(output))


// }

// long- running process (like sleep) using non-blocking execution


// func main() {
//  // cmd := exec.Command("sleep", "5") // simulate a long-running task


//  cmd := exec.Command("powershell", "-Command", " Start-Sleep -Seconds 5")


//  // start the command (non-blocking)
//  err := cmd.Start()
//  if err != nil {
//      fmt.Printf("Error starting command: %v\n", err)
//      return
//  }


//  fmt.Println("Command is running")
//  time.Sleep(2 * time.Second) // simulate other work


//  // wait for the command to complete
//  err = cmd.Wait()
//  if err != nil {
//      fmt.Printf("Command failed: %v\n", err)
//      return
//  }
//  fmt.Println("Command completed")


// }


// ------------context.timeout() to limit how long a process can run-----------


// func main() {


//  // set a 3-second timeout
//  ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)


//  defer cancel()


//  cmd := exec.CommandContext(ctx, "powershell", "-Command", " Start-Sleep -Seconds 5")


//  err := cmd.Run()


//  if ctx.Err() == context.DeadlineExceeded {
//      fmt.Println("Connection timeout")
//  } else if err != nil {
//      fmt.Printf("Command faild: %v\n", err)
//  } else {
//      fmt.Println("Command completed successfully")
//  }


// }


//  Communication Input and Output Streams-----------------
// Stdin, Stdout and Stderr


func main() {


    cmd := exec.Command("findstr", "hello")
    //cmd := exec.Command("powershell", "-Command", "echo 'hello world' | Select-String 'hello'")


    // Provide input to child process
    cmd.Stdin = bytes.NewBufferString("hello world\n foor bar")


    // capture output
    output, err := cmd.Output()


    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }


    fmt.Println("Command Output")
    fmt.Println(string(output))


}

