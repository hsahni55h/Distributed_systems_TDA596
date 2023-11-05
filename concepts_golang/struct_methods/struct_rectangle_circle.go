package main

import (
    "fmt"
    "math"
)

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
func (c Circle) Circumference() float64 {
    return 2 * math.Pi * c.Radius
}

func main() {
    // Create instances of the "Rectangle" and "Circle" structs
    rect := Rectangle{Width: 5, Height: 3}
    circle := Circle{Radius: 4.5}

    // Calculate and print properties of the rectangle
    rectArea := rect.Area()
    rectPerimeter := rect.Perimeter()

    fmt.Printf("Rectangle Area: %.2f square units\n", rectArea)
    fmt.Printf("Rectangle Perimeter: %.2f units\n", rectPerimeter)

    // Calculate and print properties of the circle
    circleArea := circle.Area()
    circleCircumference := circle.Circumference()

    fmt.Printf("Circle Area: %.2f square units\n", circleArea)
    fmt.Printf("Circle Circumference: %.2f units\n", circleCircumference)
}
