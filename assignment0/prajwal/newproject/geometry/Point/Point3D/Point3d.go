package Point3D

import (
	"fmt"
	"math"
	"newproject/geometry/Point"
)

type Point3D[T Point.Num] struct {
	X, Y, Z T
}

func (p *Point3D[T]) GetX() T {
	return p.X
}

func (p *Point3D[T]) GetY() T {
	return p.Y
}

func (p *Point3D[T]) Distance(other interface{}) T {
	if p3, ok := other.(*Point3D[T]); ok {
		dx := p3.X - p.X
		dy := p3.Y - p.Y
		dz := p3.Z - p.Z
		return T(math.Sqrt(float64(dx*dx + dy*dy + dz*dz)))
	}
	return 0 // Handle the case where 'other' is not a Point3D[T Num]
}

func (p *Point3D[T]) String() string {
	return fmt.Sprintf("Point3D{X: %.2f, Y: %.2f, Z: %.2f}", p.X, p.Y, p.Z)
}
