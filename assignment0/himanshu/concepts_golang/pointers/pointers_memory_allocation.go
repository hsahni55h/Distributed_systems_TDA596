package main

import "fmt"

func main() {
    var p *int
    p = new(int) // Allocate memory for an integer
    *p = 42
    fmt.Printf("Value at memory address %p: %d\n", p, *p)
    freeMemory(p)
}

func freeMemory(p *int) {
    fmt.Printf("Freeing memory at address %p\n", p)
    // Code to release memory goes here
}
