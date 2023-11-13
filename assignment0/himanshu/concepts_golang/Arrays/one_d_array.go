package main

import "fmt"

func main() {
	// Array declaration and initialization
	fmt.Println("Array declaration and initialization:")
	var a [5]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5
	fmt.Println(a)

	// Short array declaration
	fmt.Println("\nShort array declaration:")
	b := [3]string{"apple", "banana", "cherry"}
	fmt.Println(b)

	// Accessing array elements
	fmt.Println("\nAccessing array elements:")
	fmt.Printf("a[2]: %d\n", a[2])
	fmt.Printf("b[1]: %s\n", b[1])

	// Looping through an array
	fmt.Println("\nLooping through an array:")
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d]: %d\n", i, a[i])
	}

	// Using range in a for loop
	fmt.Println("\nUsing range in a for loop:")
	for index, value := range b {
		fmt.Printf("b[%d]: %s\n", index, value)
	}

	// Array size and range
	fmt.Println("\nArray size and range:")
	c := [4]int{10, 20, 30, 40}
	fmt.Printf("Length of c: %d\n", len(c))

	// Modifying array elements
	fmt.Println("\nModifying array elements:")
	c[1] = 25
	fmt.Println(c)

	// Array with mixed data types (using interface{})
	fmt.Println("\nArray with mixed data types:")
	mixed := [3]interface{}{5, "apple", true}
	fmt.Println(mixed)
}
