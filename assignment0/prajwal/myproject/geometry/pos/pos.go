// geometry/pos/pos.go

package pos

type Num interface {
	~float64 | ~float32 | ~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Pos[T Num] interface {
	New(vals []T) interface{}
	Distance(Ipos2 interface{}) T
	// the method can accept a parameter of any type that satisfies the Distance interface, 
	// not just a specific type like Pos2D or Pos3D.
	SetProperty(val T, prop string) error
	GetProperty(prop string) (T, error)
}
