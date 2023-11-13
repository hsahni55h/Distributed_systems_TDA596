package main

import "fmt"

func main() {
	x := 5
	y := 6
	str1 := "Hello"
	str2 := "World"

	// Simple if statement
	if x < 10 {
		fmt.Printf("x is less than 10\n")
	}

	// if-else statement
	if y > 7 {
		fmt.Printf("y is greater than 7\n")
	} else {
		fmt.Printf("y is not greater than 7\n")
	}

	// if-else if-else statement
	if x > 10 {
		fmt.Printf("x is greater than 10\n")
	} else if x < 10 {
		fmt.Printf("x is less than 10\n")
	} else {
		fmt.Printf("x is equal to 10\n")
	}

	// Complex conditions
	if x > 3 && y < 8 {
		fmt.Printf("x is greater than 3 and y is less than 8\n")
	}

	if str1 == "Hello" && str2 == "World" {
		fmt.Printf("Both strings are as expected\n")
	} else {
		fmt.Printf("At least one string is not as expected\n")
	}

	// Nested if statements
	if x > 0 {
		if y > 0 {
			fmt.Printf("Both x and y are positive\n")
		} else {
			fmt.Printf("x is positive, but y is not\n")
		}
	} else {
		fmt.Printf("x is not positive\n")
	}

	// Short if statement with variable assignment
	result := "Even"
	if x%2 != 0 {
		result = "Odd"
	}
	fmt.Printf("x is %s\n", result)
}
