package main

import "fmt"

func main() {
	var x int = 42
	var p *int
	p = &x

	fmt.Printf("x = %d\n", x)
	fmt.Printf("*p = %d\n", *p) // Use the dereference operator to access the value
	*p = 10                     // Modify the value through the pointer
	fmt.Printf("x = %d\n", x)   // x is now 10
}
