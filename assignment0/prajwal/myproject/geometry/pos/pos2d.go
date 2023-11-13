// geometry/pos2d.go

package pos

import (
	"fmt"
	"errors"
	"log"
    "math"
)

type Pos2D[T Num] struct {
    x T
    y T
}

func (Pos2D[T]) New(vals []T) interface{} {
    return Pos2D[T]{vals[0], vals[1]}
}

func (pos *Pos2D[T]) SetX(xval T) error {
    pos.x = xval
	return nil
}

func (pos *Pos2D[T]) SetY(yval T) error {
    pos.y = yval
	return nil
}

func (pos Pos2D[T]) X() T {
    return pos.x
}

func (pos Pos2D[T]) Y() T {
    return pos.y
}

func (pos1 Pos2D[T]) Distance(Ipos2 interface{}) T {
	fmt.Println("Inside 2D Distance")
	
	// assertion - assert that the Ipos2 parameter is of type Pos2D. 
	// If it is, the 'ok' variable will be true, and 'pos2' will hold the value as a Pos2D
	pos2, ok := Ipos2.(Pos2D[T])
	if !ok {
		log.Fatal("Invalid type for Ipos2.")
	}

	dx := float64(pos2.x - pos1.x)
	dy := float64(pos2.y - pos1.y)
	return T(math.Sqrt((dx * dx) + (dy * dy)))
}

// Implement SetProperty and GetProperty methods as needed

func (pos Pos2D[T]) SetProperty(val T, prop string) error {
	switch prop {
		case "X":
			return pos.SetX(val)
		case "Y":
			return pos.SetY(val)
	}
	return errors.New("invalid property")
}

func (pos Pos2D[T]) GetProperty(prop string) (T, error) {
	switch prop {
		case "X":
			return pos.x, nil
		case "Y":
			return pos.y, nil
	}
	return T(0), errors.New("invalid property")
}
