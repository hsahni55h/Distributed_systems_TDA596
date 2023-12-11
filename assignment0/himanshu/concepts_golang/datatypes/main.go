package main

import (
	"fmt"
)

func main() {
	// Numeric Types
	var intVar int = 42
	var uintVar uint = 64
	var floatVar float64 = 3.14
	var complexVar complex128 = 2 + 3i

	// String Type
	var strVar string = "Hello, Go!"

	// Boolean Type
	var boolVar bool = true

	// Derived Types
	var arrVar [3]int = [3]int{1, 2, 3}
	var sliceVar []int = []int{4, 5, 6}
	var mapVar map[string]int = map[string]int{"a": 7, "b": 8}
	var structVar struct {
		Name string
		Age  int
	} = struct {
		Name string
		Age  int
	}{"John", 30}
	var ptrVar *int = &intVar

	// Interface Type
	type Shape interface {
		Area() float64
	}
	var shapeVar Shape

	// Channel Type
	var chanVar chan int = make(chan int)

	// Byte Type
	var byteVar byte = 'A'

	// Rune Type
	var runeVar rune = '😊'

	// Print values
	fmt.Println("Numeric Types:")
	fmt.Println("intVar:", intVar)
	fmt.Println("uintVar:", uintVar)
	fmt.Println("floatVar:", floatVar)
	fmt.Println("complexVar:", complexVar)

	fmt.Println("\nString Type:")
	fmt.Println("strVar:", strVar)

	fmt.Println("\nBoolean Type:")
	fmt.Println("boolVar:", boolVar)

	fmt.Println("\nDerived Types:")
	fmt.Println("arrVar:", arrVar)
	fmt.Println("sliceVar:", sliceVar)
	fmt.Println("mapVar:", mapVar)
	fmt.Println("structVar:", structVar)
	fmt.Println("ptrVar:", *ptrVar)

	fmt.Println("\nInterface Type:")
	fmt.Println("shapeVar:", shapeVar)

	fmt.Println("\nChannel Type:")
	fmt.Println("chanVar:", chanVar)

	fmt.Println("\nByte Type:")
	fmt.Println("byteVar:", byteVar)

	fmt.Println("\nRune Type:")
	fmt.Println("runeVar:", runeVar)
}
