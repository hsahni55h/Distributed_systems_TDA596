package main

import (
	"fmt"
	"math"
)

type Num interface {
	~float64 | ~float32 | ~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Point[T Num] interface {
	GetX() T
	GetY() T
	Distance(other Point[T]) T
	String() string
}

type Point2D[T Num] struct {
	X, Y T
}

type Point3D[T Num] struct {
	X, Y, Z T
}

func (p *Point2D[T]) GetX() T {
	return p.X
}

func (p *Point2D[T]) GetY() T {
	return p.Y
}

func (p *Point2D[T]) Distance(other Point[T]) T {
	if p2, ok := other.(*Point2D[T]); ok {
		dx := float64(p2.X - p.X)
		dy := float64(p2.Y - p.Y)
		return T(math.Sqrt(dx*dx + dy*dy))
	}
	return 0 // Handle the case where 'other' is not a Point2D[T Num]
}

func (p *Point2D[T]) String() string {
	return fmt.Sprintf("Point2D{X: %.2f, Y: %.2f}", p.X, p.Y)
}

func (p *Point3D[T]) GetX() T {
	return p.X
}

func (p *Point3D[T]) GetY() T {
	return p.Y
}

func (p *Point3D[T]) Distance(other Point[T]) T {
	if p3, ok := other.(*Point3D[T]); ok {
		dx := float64(p3.X - p.X)
		dy := float64(p3.Y - p.Y)
		dz := float64(p3.Z - p.Z)
		return T(math.Sqrt(dx*dx + dy*dy + dz*dz))
	}
	return 0 // Handle the case where 'other' is not a Point3D[T Num]
}

func (p *Point3D[T]) String() string {
	return fmt.Sprintf("Point3D{X: %.2f, Y: %.2f, Z: %.2f}", p.X, p.Y, p.Z)
}

func main() {
	p2d1 := &Point2D[float64]{X: 3.0, Y: 4.0}
	p2d2 := &Point2D[float64]{X: 0.0, Y: 0.0}
	p3d := &Point3D[float64]{X: 3.0, Y: 4.0, Z: 5.0}

	fmt.Println("Point2D 1:", p2d1.String())
	fmt.Println("Point2D 2:", p2d2.String())
	fmt.Println("Point3D:", p3d.String())

	// Create an array of type Point and store the created points
	points := []Point[float64]{p2d1, p2d2, p3d}

	// Calculate the distance between each point
	for i, p1 := range points {
		for j, p2 := range points {
			if i != j {
				distance := p1.Distance(p2)
				fmt.Printf("Distance between %s and %s: %f\n", p1, p2, distance)
			}
		}
	}
}
