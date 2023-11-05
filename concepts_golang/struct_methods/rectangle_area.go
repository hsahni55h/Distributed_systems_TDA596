package main

import "fmt"

// Define a "Rectangle" struct
type Rectangle struct {
    Width  float64
    Height float64
}

// Method to calculate the area of the rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    // Create an instance of the "Rectangle" struct
    rect := Rectangle{Width: 5, Height: 3}

    // Call the "Area" method on the "Rectangle" instance
    area := rect.Area()

    // Print the calculated area
    fmt.Printf("Area of the rectangle: %.2f square units\n", area)
}

