package main

import "fmt"

func main() {
    // Create an anonymous struct
	// We create an anonymous struct directly in the main function without defining a separate type.
    car := struct {
        Make  string
        Model string
        Year  int
    }{
        Make:  "Toyota",
        Model: "Camry",
        Year:  2022,
    }

    // Access and print the fields of the anonymous struct
    fmt.Println("Make:", car.Make)
    fmt.Println("Model:", car.Model)
    fmt.Println("Year:", car.Year)
}
