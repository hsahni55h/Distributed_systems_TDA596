package main

import "fmt"

func main() {
	// Creating an array
	originalArray := [6]int{1, 2, 3, 4, 5, 6}

	// Creating slices from the array
	slice1 := originalArray[1:4] // Elements at index 1, 2, 3
	slice2 := originalArray[:3]  // Elements at index 0, 1, 2
	slice3 := originalArray[3:]  // Elements at index 3, 4, 5

	// Displaying the original array and slices
	fmt.Println("Original Array:", originalArray)
	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)
	fmt.Println("Slice 3:", slice3)

	// Modifying a slice also modifies the original array
	slice2[0] = 99
	fmt.Println("\nAfter modifying Slice 2:")
	fmt.Println("Original Array:", originalArray)
	fmt.Println("Slice 2:", slice2)

	// Creating a new slice using make
	slice4 := make([]int, 3, 5) // Make a slice with a length of 3 and a capacity of 5
	slice4[0] = 10
	slice4[1] = 20
	slice4[2] = 30
	fmt.Println("\nSlice 4:", slice4)

	// Appending elements to a slice
	slice4 = append(slice4, 40, 50) // Appending two elements
	fmt.Println("After appending to Slice 4:", slice4)

	// Copying a slice
	copySlice := make([]int, len(slice4))
	copy(copySlice, slice4)
	fmt.Println("\nCopy of Slice 4:", copySlice)
}
