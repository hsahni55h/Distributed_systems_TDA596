package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start of the program")

	// Simulate some work
	fmt.Println("Doing some initial work...")

	// Introduce a delay of 3 seconds
	time.Sleep(3 * time.Second)

	// Continue with the program
	fmt.Println("After the delay, continuing with the program")
}
