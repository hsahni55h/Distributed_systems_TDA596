package main

import "fmt"

// add function accepts two integers, adds them, and returns the result.
func add(x, y int) int {
    return x + y
}

func main() {
    result := add(5, 3)
    fmt.Printf("The sum is: %d\n", result)
}
