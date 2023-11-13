package main

import (
	"bufio"    // Import the bufio package for reading input
	"fmt"      // Import the fmt package for printing
	"os"       // Import the os package for accessing standard input
	"strconv"  // Import the strconv package for string to integer conversion
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)  // Create a new scanner to read from standard input
	fmt.Printf("Type your birth year: ")   // Print a prompt to the console
	scanner.Scan()                         // Read a line of input from the user

	// Get the user's input as a string
	input := scanner.Text()

	// Convert the input to an integer
	birthYear, err := strconv.ParseInt(input, 10, 64)

	// Check for errors when parsing the input
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid year.") // Print an error message
		return // Exit the program
	}

	// Calculate the age by subtracting the birth year from the current year (2023)
	age := 2023 - birthYear
	fmt.Printf("You are %d years old\n", age) // Print the calculated age
}
