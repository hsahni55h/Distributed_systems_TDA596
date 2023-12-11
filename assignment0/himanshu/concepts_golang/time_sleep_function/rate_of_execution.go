package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start of the program")

	// Simulate processing items in a loop with a rate limit
	items := []string{"item1", "item2", "item3"}

	// Process each item with a delay of 2 seconds
	for _, item := range items {
		processItem(item)
		time.Sleep(2 * time.Second)
	}
}

func processItem(item string) {
	fmt.Printf("Processing %s...\n", item)
}
