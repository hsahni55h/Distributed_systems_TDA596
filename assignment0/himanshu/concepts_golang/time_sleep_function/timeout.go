package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start of the program")

	// Set a timeout of 5 seconds
	timeout := 5 * time.Second

	// Simulate some work that may take longer than the timeout
	fmt.Println("Doing some work that may take longer...")

	// Check if the work exceeds the timeout
	select {
	case <-time.After(timeout):
		fmt.Println("Timeout exceeded! The work took too long.")
	case <-time.After(2 * time.Second):
		// case <-time.After(7 * time.Second):
		// Simulate some additional work that takes 2 seconds
		fmt.Println("Additional work done in time.")
	}

	// Continue with the program
	fmt.Println("After the timeout check, continuing with the program")
}
