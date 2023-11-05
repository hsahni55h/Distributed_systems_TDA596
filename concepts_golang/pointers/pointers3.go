package main

import "fmt"

type Point struct {
	x, y int
}

func main() {
	p := Point{2, 3}
	var pp *Point
	fmt.Printf("p = %d\n", p)
	pp = &p
	fmt.Printf("pp = %d\n", pp)
	fmt.Printf("*pp = %d\n", *pp)
	pp.x = 4
	fmt.Println(p) // p has been modified through pp
}
