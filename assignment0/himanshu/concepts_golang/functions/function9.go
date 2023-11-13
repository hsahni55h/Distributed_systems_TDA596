package main

import "fmt"

// applyFunction applies a function to each element of a slice.
func applyFunction(numbers []int, fn func(int) int) []int {
    result := make([]int, len(numbers))
    for i, num := range numbers {
        result[i] = fn(num)
    }
    return result
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    
    double := func(x int) int { return x * 2 }
    squared := func(x int) int { return x * x }
    
    doubledNumbers := applyFunction(numbers, double)
    squaredNumbers := applyFunction(numbers, squared)
    
    fmt.Println("Original Numbers:", numbers)
    fmt.Println("Doubled Numbers:", doubledNumbers)
    fmt.Println("Squared Numbers:", squaredNumbers)
}
