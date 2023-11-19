package main

import "fmt"

// sum function accepts a variable number of integers and returns their sum.
func sum(numbers ...int) int { // this is an example of varadic function
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	result := sum(1, 2, 3, 4, 5)
	fmt.Printf("The sum is: %d\n", result)
}
