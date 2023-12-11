package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random integer between 0 and 99
	randomInt := rand.Intn(100)
	fmt.Println("Random Integer:", randomInt)

	// Generate a random float64 between 0.0 and 1.0
	randomFloat := rand.Float64()
	fmt.Println("Random Float:", randomFloat)
}
