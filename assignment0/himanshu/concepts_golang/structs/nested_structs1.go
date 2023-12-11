package main

import "fmt"

// Define an "Address" struct
type Address struct {
    Street  string
    City    string
    Country string
}

// Define an "Employee" struct with a nested "Address" struct
type Employee struct {
    FirstName    string
    LastName     string
    Age          int
    HomeAddress  Address
}

func main() {
    // Create an instance of the "Employee" struct with nested "Address" struct
    emp := Employee{
        FirstName: "John",
        LastName:  "Doe",
        Age:       30,
        HomeAddress: Address{
            Street:  "123 Main St",
            City:    "Anytown",
            Country: "USA",
        },
    }

    // Access and print the fields of the "Employee" and "Address" structs
    fmt.Println("First Name:", emp.FirstName)
    fmt.Println("Last Name:", emp.LastName)
    fmt.Println("Age:", emp.Age)
    fmt.Println("Street:", emp.HomeAddress.Street)
    fmt.Println("City:", emp.HomeAddress.City)
    fmt.Println("Country:", emp.HomeAddress.Country)
}
