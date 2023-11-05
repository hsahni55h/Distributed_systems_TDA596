package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{4, 2, 7, 1, 9}
	pointers := make([]*int, len(numbers))
	fmt.Println(numbers)
	fmt.Println(pointers)

	// for i, v := range numbers {
	//     pointers[i] = &v
	//     fmt.Printf("%d ", *pointers[i])
	// }

	for i, v := range numbers {
		value := v // Create a new variable to capture the current value of v
		pointers[i] = &value
		fmt.Printf("%d \n", *pointers[i])
	}

	sort.Slice(pointers, func(i, j int) bool {
		return *pointers[i] < *pointers[j]
	})

	for _, p := range pointers {
		fmt.Printf("%d ", *p)
	}
	fmt.Println()
}
