// geometry/pos3d.go

package pos

import (
    "errors"
	"fmt"
	"log"
    "math"
)

type Pos3D[T Num] struct {
    x T
    y T
    z T
}

func (Pos3D[T]) New(vals []T) Pos3D[T] {
    return Pos3D[T]{vals[0], vals[1], vals[2]}
}

func (pos *Pos3D[T]) SetX(xval T) error {
    pos.x = xval
    return nil
}

func (pos *Pos3D[T]) SetY(yval T) error {
    pos.y = yval
    return nil
}

func (pos *Pos3D[T]) SetZ(zval T) error {
    pos.z = zval
    return nil
}

func (pos Pos3D[T]) X() T {
    return pos.x
}

func (pos Pos3D[T]) Y() T {
    return pos.y
}

func (pos Pos3D[T]) Z() T {
    return pos.z
}

func (pos1 Pos3D[T]) Distance(Ipos2 interface{}) T {
	fmt.Println("Inside 2D Distance")
	
	// assertion - assert that the Ipos2 parameter is of type Pos3D. 
	// If it is, the 'ok' variable will be true, and 'pos2' will hold the value as a Pos3D
	pos2, ok := Ipos2.(Pos3D[T])
	if !ok {
		log.Fatal("Invalid type for Ipos2.")
	}

	dx := float64(pos2.x - pos1.x)
	dy := float64(pos2.y - pos1.y)
    dz := float64(pos2.z - pos1.z)
	return T(math.Sqrt((dx * dx) + (dy * dy) + (dz * dz)))
}

// Implement SetProperty and GetProperty methods as needed

func (pos Pos3D[T]) SetProperty(val T, prop string) error {
	switch prop {
		case "X":
			return pos.SetX(val)
		case "Y":
			return pos.SetY(val)
        case "Z":
			return pos.SetZ(val)
	}
	return errors.New("invalid property")
}

func (pos Pos3D[T]) GetProperty(prop string) (T, error) {
	switch prop {
		case "X":
			return pos.x, nil
		case "Y":
			return pos.y, nil
        case "Z":
			return pos.z, nil
	}
	return T(0), errors.New("invalid property")
}
