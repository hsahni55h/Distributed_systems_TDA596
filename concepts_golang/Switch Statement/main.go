package main

import "fmt"

func main() {
	// Simple switch statement
	fmt.Println("Simple switch statement:")
	dayOfWeek := 2
	switch dayOfWeek {
	case 1:
		fmt.Println("It's Monday")
	case 2:
		fmt.Println("It's Tuesday")
	case 3:
		fmt.Println("It's Wednesday")
	default:
		fmt.Println("It's another day of the week")
	}

	// Multiple cases with fallthrough
	fmt.Println("\nMultiple cases with fallthrough:")
	number := 3
	switch number {
	case 1:
		fmt.Println("The number is 1")
		fallthrough
	case 2:
		fmt.Println("The number is 2")
	case 3:
		fmt.Println("The number is 3")
	default:
		fmt.Println("The number is not 1, 2, or 3")
	}

	// Using a switch without an expression
	fmt.Println("\nUsing a switch without an expression:")
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	case score >= 60:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}

	// Type switch with interface{}
	fmt.Println("\nType switch with interface{}:")
	var x interface{}
	x = 3.14
	switch x.(type) {
	case int:
		fmt.Println("x is an integer")
	case float64:
		fmt.Println("x is a float")
	default:
		fmt.Println("x is of another type")
	}

	// Switch statement in a loop
	fmt.Println("\nSwitch statement in a loop:")
	for i := 1; i <= 5; i++ {
		switch i {
		case 1, 2:
			fmt.Println("One or two")
		case 3, 4:
			fmt.Println("Three or four")
		default:
			fmt.Println("Other number")
		}
	}
}
