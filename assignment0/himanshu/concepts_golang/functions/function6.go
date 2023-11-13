package main

import "fmt"

// factorial calculates the factorial of a non-negative integer.
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    result := factorial(5)
    fmt.Printf("Factorial of 5 is: %d\n", result)
}
