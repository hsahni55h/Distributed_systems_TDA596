package main

import "fmt"

func main() {
	// Create a buffered channel with a capacity of 3
	bufferedCh := make(chan int, 3)

	// Send values to the channel
	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3

	// Receive values from the channel
	fmt.Println("Received:", <-bufferedCh)
	fmt.Println("Received:", <-bufferedCh)
	fmt.Println("Received:", <-bufferedCh)

	// The channel is now empty, sending more values would block unless there's a corresponding receiver
}
