package main

import (
    "fmt"
    "math"
)

// Define a "Shape" interface
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Define a "Rectangle" struct
type Rectangle struct {
    Width  float64
    Height float64
}

// Method to calculate the area of the rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method to calculate the perimeter of the rectangle
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Define a "Circle" struct
type Circle struct {
    Radius float64
}

// Method to calculate the area of the circle
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

// Method to calculate the circumference of the circle
func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

// Function to print the area of a shape
func PrintArea(s Shape) {
    fmt.Printf("Area: %.2f square units\n", s.Area())
}

func main() {
    r := Rectangle{Width: 5, Height: 3}
    c := Circle{Radius: 4.5}

    // Use the PrintArea function with different shapes
    PrintArea(r)
    PrintArea(c)
}
