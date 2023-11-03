package main

import "fmt"

// greet function accepts a name as a parameter and prints a greeting message.
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	// Calling the greet function
	greet("Alice")
	greet("Bob")
}
