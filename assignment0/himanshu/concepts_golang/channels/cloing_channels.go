package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()

	// Receive values until the channel is closed
	for value := range ch {
		fmt.Println("Received:", value)
	}
}
