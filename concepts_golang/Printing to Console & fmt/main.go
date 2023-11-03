package main

import "fmt" // display messages or text, print to the console

func main() { // entrypoint

	// %T - datatype of variable
	// %v - value of variable
	// %% - literal %
	// %t - boolean(true or false)

	fmt.Printf("Hello %t\n", true)

	// %b - base 2 (binary)
	// %o - base 8 (octal)
	// %d - base 10 (decimal)
	// %x - base 16 (hexadecimal)

	fmt.Printf("Number (binary): %b\n", 1025)
	fmt.Printf("Number (octal): %o\n", 1025)
	fmt.Printf("Number (decimal): %d\n", 1025)
	fmt.Printf("Number (hexadecimal): %x\n", 1025)

	// %e - scientific notation
	// %f/%F - decimal no exponent
	// %g - for large exponents

	fmt.Printf("Number (scientific notation): %e\n", 2.364485569651212)
	fmt.Printf("Number: %f\n", 2.364485569651212)
	fmt.Printf("Number: %g\n", 2.364485569651212)

	// %s - default
	// %q - double-quoted string

	fmt.Printf("Number: %s\n", "Himanshu")
	fmt.Printf("Number: %q\n", "Himanshu")

	// %f - default width, default precision
	// %9f - width 9, default precision
	// %.2f - default width, precision 2
	// %9.2f - width 9, precision 2
	// %9.f - width 9, precision 0

	fmt.Printf("Number: %15q\n", "Himanshu")
	fmt.Printf("Number: %-15q is cool \n", "Himanshu")
	fmt.Printf("Number: %.2f\n", 3.45698452)
	fmt.Printf("Number: %.f\n", 3.45698452)

	// %09d - pads digit to length 9 with preceding 0
	// %-4d - pads with spaces (width 9, left justified)

	fmt.Printf("Number: %07d\n", 58)

	var out string = fmt.Sprintf("Number: %d\n", 58)
	fmt.Println(out)

}
