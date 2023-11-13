package main

import "fmt"

func main() {
	var x int = 42
	var p *int // Declare a pointer to an integer
	p = &x     // Assign the memory address of x to p

	fmt.Printf("x = %d\n", x)
	fmt.Printf("p = %p\n", p)
	// fmt.Printf("p = %d\n", *p)
}
