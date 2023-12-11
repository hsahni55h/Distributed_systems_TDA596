package main

import "fmt"

func main() {
    // Create an unbuffered channel of integers
    ch := make(chan int)

    // Goroutine to send a value to the channel
    go func() {
        ch <- 42
    }()

    // Receive the value from the channel
    value := <-ch

    fmt.Println("Received:", value)
}
