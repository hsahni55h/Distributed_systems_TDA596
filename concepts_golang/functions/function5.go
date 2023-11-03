package main

import "fmt"

// divide returns both the quotient and remainder of dividing two integers.
func divide(x, y int) (int, int) {
    quotient := x / y
    remainder := x % y
    return quotient, remainder
}

func main() {
    quotient, remainder := divide(10, 3)
    fmt.Printf("Quotient: %d, Remainder: %d\n", quotient, remainder)
}
