package Point

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
