/*
Length (len): The length of a slice is the number of elements it contains. It represents the number of elements accessible within the slice.

Capacity (cap): The capacity of a slice is the maximum number of elements it can hold without allocating additional memory.
It represents the number of elements that can be stored in the underlying array without triggering a reallocation.

*/

package main

import "fmt"

func main() {
	// Creating a slice
	originalSlice := []int{1, 2, 3, 4, 5, 6}

	// Slicing the original slice
	slice1 := originalSlice[1:4] // Length: 3, Capacity: 5
	slice2 := originalSlice[:3]  // Length: 3, Capacity: 5
	slice3 := originalSlice[3:]  // Length: 3, Capacity: 3

	// Displaying the original slice and slices
	fmt.Println("Original Slice:", originalSlice)
	fmt.Printf("Slice 1: Length = %d, Capacity = %d, Values: %v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("Slice 2: Length = %d, Capacity = %d, Values: %v\n", len(slice2), cap(slice2), slice2)
	fmt.Printf("Slice 3: Length = %d, Capacity = %d, Values: %v\n", len(slice3), cap(slice3), slice3)
}
