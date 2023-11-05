package main

import "fmt"

func sayHello() {
    fmt.Println("Hello, World!")
}

func invoke(fn func()) {
    fn()
}

func main() {
    var pf func() // Declare a pointer to a function
    pf = sayHello
    invoke(pf) // Call the function through the pointer
}
