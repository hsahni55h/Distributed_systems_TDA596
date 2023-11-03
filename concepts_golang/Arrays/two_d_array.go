package main

import "fmt"

func main() {
	// Declaration and initialization of a 2D array
	fmt.Println("2D Array declaration and initialization:")
	var matrix [3][3]int
	matrix[0] = [3]int{1, 2, 3}
	matrix[1] = [3]int{4, 5, 6}
	matrix[2] = [3]int{7, 8, 9}

	// Accessing elements in a 2D array
	fmt.Println("\nAccessing elements in a 2D array:")
	fmt.Printf("matrix[0][1]: %d\n", matrix[0][1])
	fmt.Printf("matrix[2][2]: %d\n", matrix[2][2])

	// Looping through a 2D array
	fmt.Println("\nLooping through a 2D array:")
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			fmt.Printf("matrix[%d][%d]: %d\n", row, col, matrix[row][col])
		}
	}

	// Using range with a 2D array
	fmt.Println("\nUsing range with a 2D array:")
	for row, rowData := range matrix {
		for col, value := range rowData {
			fmt.Printf("matrix[%d][%d]: %d\n", row, col, value)
		}
	}

	// Modifying elements in a 2D array
	fmt.Println("\nModifying elements in a 2D array:")
	matrix[1][1] = 0
	fmt.Println(matrix)

	// Declaration and initialization of a jagged 2D array
	fmt.Println("\nJagged 2D Array declaration and initialization:")
	jaggedMatrix := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(jaggedMatrix)
}
