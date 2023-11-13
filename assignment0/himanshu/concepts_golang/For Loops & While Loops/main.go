package main

import "fmt"

func main() {
	// Basic for loop
	fmt.Println("Basic for loop:")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// While loop using for
	fmt.Println("\nWhile loop using for:")
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// Infinite loop with break
	fmt.Println("\nInfinite loop with break:")
	k := 0
	for {
		fmt.Println(k)
		k++
		if k >= 5 {
			break
		}
	}

	// Continue and conditional statements in loops
	fmt.Println("\nContinue and conditional statements in loops:")
	for m := 0; m < 10; m++ {
		if m%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Println(m)
	}

	// Nested loops with break
	fmt.Println("\nNested loops with break:")
	for n := 1; n <= 3; n++ {
		fmt.Printf("Outer loop iteration %d\n", n)
		for p := 1; p <= 4; p++ {
			fmt.Printf("Inner loop iteration %d\n", p)
			if p == 3 {
				break // Break out of the inner loop
			}
		}
	}

	// Loop with a conditional statement
	fmt.Println("\nLoop with a conditional statement:")
	for q := 1; q <= 5; q++ {
		if q%2 == 0 {
			fmt.Printf("%d is even\n", q)
		} else {
			fmt.Printf("%d is odd\n", q)
		}
	}
}
