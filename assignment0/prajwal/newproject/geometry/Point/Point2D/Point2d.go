package Point2D

import (
	"fmt"
	"math"
	"newproject/geometry/Point"
)

type Point2D[T Point.Num] struct {
	X, Y T
}

func (p *Point2D[T]) GetX() T {
	return p.X
}

func (p *Point2D[T]) GetY() T {
	return p.Y
}

func (p *Point2D[T]) Distance(other interface{}) T {
	if p2, ok := other.(*Point2D[T]); ok {
		dx := p2.X - p.X
		dy := p2.Y - p.Y
		return T(math.Sqrt(float64(dx*dx + dy*dy)))
	}
	return 0 // Handle the case where 'other' is not a Point2D[T Num]
}

func (p *Point2D[T]) String() string {
	return fmt.Sprintf("Point2D{X: %.2f, Y: %.2f}", p.X, p.Y)
}
