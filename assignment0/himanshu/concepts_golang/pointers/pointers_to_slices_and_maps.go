package main

import "fmt"

func main() {
	var s1 []int
	s2 := make([]int, 5)
	fmt.Println(s1)
	fmt.Println(s2)
	var p *[]int

	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1)
	p = &s1
	(*p)[1] = 99
	fmt.Println(s1)

	s2[2] = 42
	fmt.Println(s2)
	p = &s2
	(*p)[4] = 88
	fmt.Println(s2)

	fmt.Println(s1) // [1 99 3]
	fmt.Println(s2) // [0 0 42 0 88]
}
