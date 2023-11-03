package main

import "fmt"

func main() {
	// Addition (+)
	var num1 int = 5
	var num2 int = 3
	sum := num1 + num2
	fmt.Printf("Addition: %d + %d = %d\n", num1, num2, sum)

	// Subtraction (-)
	var num3 int = 8
	var num4 int = 2
	difference := num3 - num4
	fmt.Printf("Subtraction: %d - %d = %d\n", num3, num4, difference)

	// Multiplication (*)
	var num5 float64 = 6.5
	var num6 float64 = 4.0
	product := num5 * num6
	fmt.Printf("Multiplication: %g * %g = %g\n", num5, num6, product)

	// Division (/)
	var num7 float32 = 9.0
	var num8 float32 = 2.0
	quotient := num7 / num8
	fmt.Printf("Division: %g / %g = %g\n", num7, num8, quotient)

	// Integer Division (//) - Discards the fractional part
	var num9 int = 9
	var num10 int = 4
	intDivision := num9 / num10
	fmt.Printf("Integer Division: %d / %d = %d\n", num9, num10, intDivision)

	// Modulus (%) - Calculates the remainder
	var num11 int = 11
	var num12 int = 3
	remainder := num11 % num12
	fmt.Printf("Modulus: %d %% %d = %d\n", num11, num12, remainder)

	// Increment (++)
	var num13 int = 7
	num13++
	fmt.Printf("Increment: After incrementing, num13 = %d\n", num13)

	// Decrement (--)
	var num14 int = 10
	num14--
	fmt.Printf("Decrement: After decrementing, num14 = %d\n", num14)
}
