package main

/*
	struct
	association of method to struct

	generics - similar to templates

	interface - to implement polymorphic behaviour
*/

import (
	"errors"
	"fmt"
	"log"
	"math"
)

type Num interface {
	~float64 | ~float32 | ~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Pos[T Num] interface {
	Distance(Ipos2 interface{}) T
	// the method can accept a parameter of any type that satisfies the Distance interface, 
	// not just a specific type like Pos2D or Pos3D.
	SetProperty(val T, prop string) error
	GetProperty(prop string) (T, error)
}

type Pos2D[T Num] struct {
	x T
	y T
}

func (pos Pos2D[T]) SetX(xval T) error {
	pos.x = xval
	var err error = nil
	return err
}

func (pos Pos2D[T]) SetY(yval T) error {
	pos.y = yval
	var err error = nil
	return err
}

func (pos Pos2D[T]) SetProperty(val T, prop string) error {
	switch prop {
		case "X":
			return pos.SetX(val)
		case "Y":
			return pos.SetY(val)
	}
	return errors.New("Invalid property.")
}


func (pos Pos2D[T]) X() T {
	return pos.x
}

func (pos Pos2D[T]) Y() T {
	return pos.y
}

func (pos Pos2D[T]) GetProperty(prop string) (T, error) {
	switch prop {
		case "X":
			return pos.x, nil
		case "Y":
			return pos.y, nil
	}
	return 0, errors.New("Invalid property.")
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

type Pos3D[T Num] struct {
	x T
	y T
	z T
}

func (pos Pos3D[T]) SetX(xval T) error {
	pos.x = xval
	return nil
}

func (pos Pos3D[T]) SetY(yval T) error {
	pos.y = yval
	return nil
}

func (pos Pos3D[T]) SetZ(zval T) error {
	pos.z = zval
	return nil
}

func (pos Pos3D[T]) SetProperty(val T, prop string) error {
	switch prop {
		case "X":
			return pos.SetX(val)
		case "Y":
			return pos.SetY(val)
		case "Z":
			return pos.SetZ(val)
	}
	return errors.New("Invalid property.")
}

func (pos Pos3D[T]) X() T {return pos.x}
func (pos Pos3D[T]) Y() T {return pos.y}
func (pos Pos3D[T]) Z() T {return pos.z}

func (pos Pos3D[T]) GetProperty(prop string) (T, error) {
	switch prop {
	case "X":
		return pos.X(), nil
	case "Y":
		return pos.Y(), nil
	case "Z":
		return pos.Z(), nil
	}
	return 0, errors.New("Invalid property.")	
}

func (pos1 Pos3D[T]) Distance(Ipos2 interface{}) T {
	fmt.Println("Inside 3D Distance")
	
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

func main() {
	p1 := Pos2D[float64] {
		3.0,
		4.0,
	}

	var o Pos2D[float64]
	o.x = 0.0
	o.y = 0.0

	fmt.Println("Distance between", p1, "and", o, "is:", p1.Distance(o))

	n := 4
	pos_list := make([]Pos[float64], n)
	half := n / 2
	for i := 0; i < half; i++ {
		pos_list[i] = Pos2D[float64]{x: 3.0, y:4.0}
		pos_list[half + i] = Pos3D[float64]{x: 3.0, y:4.0, z:0.0}
	}

	for i := 0; i < half-1; i++ {
		fmt.Println("Distance between", pos_list[i], "and", pos_list[i+1], "is:", pos_list[i].Distance(pos_list[i+1]))
		fmt.Println("Distance between", pos_list[half+i], "and", pos_list[half+i+1], "is:", pos_list[half+i].Distance(pos_list[half+i+1]))
	}
	
}

