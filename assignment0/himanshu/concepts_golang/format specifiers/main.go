package main

import "fmt"

func main() {
	// %v (default format)
	num := 42
	str := "Hello, Go!"
	fmt.Printf("%%v (default format):\n")
	fmt.Printf("%v\n", num) // Prints: 42
	fmt.Printf("%v\n", str) // Prints: Hello, Go!

	// %d (decimal)
	fmt.Printf("\n%%d (decimal):\n")
	fmt.Printf("%d\n", num) // Prints: 42

	// %o (octal)
	fmt.Printf("\n%%o (octal):\n")
	fmt.Printf("%o\n", num) // Prints: 52 (42 in octal)

	// %x, %X (hexadecimal)
	fmt.Printf("\n%%x, %%X (hexadecimal):\n")
	fmt.Printf("%x\n", num) // Prints: 2a (42 in hexadecimal)
	fmt.Printf("%X\n", num) // Prints: 2A (42 in hexadecimal)

	// %f (float)
	pi := 3.14159
	fmt.Printf("\n%%f (float):\n")
	fmt.Printf("%f\n", pi) // Prints: 3.141590

	// %s (string)
	fmt.Printf("\n%%s (string):\n")
	fmt.Printf("%s\n", str) // Prints: Hello, Go!

	// %t (boolean)
	flag := true
	fmt.Printf("\n%%t (boolean):\n")
	fmt.Printf("%t\n", flag) // Prints: true

	// %c (character)
	char := 'A'
	fmt.Printf("\n%%c (character):\n")
	fmt.Printf("%c\n", char) // Prints: A

	// %p (pointer)
	fmt.Printf("\n%%p (pointer):\n")
	fmt.Printf("%p\n", &num) // Prints: hexadecimal representation of the pointer

	// Additional Examples
	fmt.Printf("\nAdditional Examples:\n")
	fmt.Printf("%9f\n", pi)  // Prints: " 3.141590"
	fmt.Printf("%.2f\n", pi) // Prints: "3.14"
}
