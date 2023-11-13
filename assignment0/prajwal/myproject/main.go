// main.go
package main

import (
	"fmt"
	"myproject/geometry/pos"
)


func main() {
	var p pos.Pos[float64]
	fmt.Println(p)

	var p1 pos.Pos2D[float64]
	fmt.Println(p1)

	p1 = p1.New([]float64{3.0, 4.0}).(pos.Pos2D[float64])
	fmt.Println(p1)

	p = p1
	fmt.Println(p)
}
