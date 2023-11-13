package main

import "fmt"

// operate performs an operation on two integers and returns the result.
type operation func(int, int) int

// add and subtract functions have the same signature as the 'operation' type.
func add(x, y int) int {
    return x + y
}

func subtract(x, y int) int {
    return x - y
}

func main() {
    var op operation
    op = add
    result := op(5, 3)
    fmt.Printf("5 + 3 = %d\n", result)

    op = subtract
    result = op(5, 3)
    fmt.Printf("5 - 3 = %d\n", result)
}
