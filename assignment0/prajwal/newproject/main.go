package main

/*
	splitting into pakages and creating module
	use of interface for polymorphism
*/

import (
	"fmt"
	"newproject/geometry/Point/Point2D"
	"newproject/geometry/Point/Point3D"
)

type Num interface {
	~float64 | ~float32 | ~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Point[T Num] interface {
	GetX() T
	GetY() T
	Distance(other interface{}) T
	String() string
}

func main() {
	p2d1 := &Point2D.Point2D[float64]{X: 3.0, Y: 4.0}
	p2d2 := &Point2D.Point2D[float64]{X: 0.0, Y: 0.0}
	p3d := &Point3D.Point3D[float64]{X: 3.0, Y: 4.0, Z: 5.0}

	fmt.Println("Point2D 1:", p2d1.String())
	fmt.Println("Point2D 2:", p2d2.String())
	fmt.Println("Point3D:", p3d.String())

	// Create an array of type Point and store the created points
	points := []interface{}{p2d1, p2d2, p3d}

	// Calculate the distance between each point
	for i, p1 := range points {
		for j, p2 := range points {
			if i != j {
				distance := p2.(Point[float64]).Distance(p1)
				fmt.Printf("Distance between %s and %s is: %.2f\n", p2.(Point[float64]).String(), p1.(Point[float64]).String(), distance)
				// few values are incorrect but it is because of incorrect usage of Distance function on mixed data type which is not handled
			}
		}
	}
}
