package main

import "fmt"

// Define a struct named "Person"
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

func main() {
    // Create a new "Person" instance
    p := Person{
        FirstName: "John",
        LastName:  "Doe",
        Age:       30,
    }

    // Access and print the fields of the "Person" instance
    fmt.Println("First Name:", p.FirstName)
    fmt.Println("Last Name:", p.LastName)
    fmt.Println("Age:", p.Age)
}
