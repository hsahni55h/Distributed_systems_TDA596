package main

import "fmt"

func main() {
    // Creating a slice
    mySlice := []string{"apple", "banana", "cherry", "date"}

    // Looping through a slice using range
    fmt.Println("Looping through a slice:")
    for index, value := range mySlice {
        fmt.Printf("Index: %d, Value: %s\n", index, value)
    }

    // Using range with arrays
    myArray := [3]int{10, 20, 30}
    fmt.Println("\nUsing range with an array:")
    for index, value := range myArray {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}
