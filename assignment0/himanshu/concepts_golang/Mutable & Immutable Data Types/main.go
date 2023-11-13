package main

import "fmt"

func main() {
	// Mutable Data Types
	// ------------------

	// Slices (Mutable)
	numbers := []int{1, 2, 3, 4, 5}
	numbers[2] = 30
	numbers = append(numbers, 6)

	// Maps (Mutable)
	scores := map[string]int{
		"Alice": 85,
		"Bob":   92,
	}
	scores["Alice"] = 88
	scores["Charlie"] = 78

	// Immutable Data Types
	// --------------------

	// Strings (Immutable)
	greeting := "Hello, "
	name := "Alice"
	message := greeting + name

	// Arrays (Immutable)
	fixedSizeArray := [3]int{1, 2, 3}

	// Attempting to change the size or add elements would result in a compile-time error
	// numbers = append(numbers, 4) // Error: cannot use append on array
	// numbers[3] = 4 // Error: invalid array index 3 (out of bounds for 3-element array)

	// Tuples (Immutable)
	point := []int{2, 3}
	// Attempting to change elements would create a new slice
	// point[0] = 4 // Error: cannot assign to point[0]

	// Output
	fmt.Println("Mutable Slice:", numbers)
	fmt.Println("Mutable Map:", scores)
	fmt.Println("Immutable String:", message)
	fmt.Println("Immutable Array:", fixedSizeArray)
	fmt.Println("Immutable Tuple (Slice):", point)
}
