package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Printf - Print Formatted
	name := "Alice"
	age := 30
	fmt.Printf("fmt.Printf - Print Formatted:\n")
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// fmt.Println - Print with Newline
	fmt.Printf("\nfmt.Println - Print with Newline:\n")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	// fmt.Sprintf - Format to String
	fmt.Printf("\nfmt.Sprintf - Format to String:\n")
	formattedString := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(formattedString)

	// fmt.Errorf - Formatted Error
	fmt.Printf("\nfmt.Errorf - Formatted Error:\n")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// fmt.Fprint, fmt.Fprintln, fmt.Fprintf - Print to Writer
	fmt.Printf("\nfmt.Fprint, fmt.Fprintln, fmt.Fprintf - Print to Writer:\n")
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Print to the file using Fprintf
	fmt.Fprintf(file, "Name: %s, Age: %d\n", name, age)
}

// divide performs integer division and returns the result.
// It also checks for division by zero and returns an error in such cases.
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
