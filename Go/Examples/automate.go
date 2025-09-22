// 1. File Operations
// Example: Copying Files
 
// You can automate file operations such as copying files, renaming, or moving them.
 
package main
 
import (
    "io"
    "log"
    "os"
)
 
func main() {
    sourceFile := "sample_file.txt"
    destinationFile := "copied_file.txt"
 
    // Open the source file
    src, err := os.Open(sourceFile)
    if err != nil {
        log.Fatal(err)
    }
    defer src.Close()
 
    // Create the destination file
    dst, err := os.Create(destinationFile)
    if err != nil {
        log.Fatal(err)
    }
    defer dst.Close()
 
    // Copy the file contents
    _, err = io.Copy(dst, src)
    if err != nil {
        log.Fatal(err)
    }
 
    log.Println("File copied successfully!")
}