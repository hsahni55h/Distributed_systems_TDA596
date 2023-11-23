package main

import (
	"fmt"
	"time"
)

func main() {
	// Create two channels for communication
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1: Simulate a task that takes 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Channel 1: Done"
	}()

	// Goroutine 2: Simulate a task that takes 1 second
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Channel 2: Done"
	}()

	// Use select to wait for either ch1 or ch2
	select {
	case result := <-ch1:
		// Case 1: Goroutine 1 completes first
		fmt.Println(result)
	case result := <-ch2:
		// Case 2: Goroutine 2 completes first
		fmt.Println(result)
	case <-time.After(3 * time.Second):
		// Case 3: Timeout if neither completes within 3 seconds
		fmt.Println("Timeout: Both channels are slow")
	}
}
