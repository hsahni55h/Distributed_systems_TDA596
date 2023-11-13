
package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        defer fmt.Println(i) // Deferred print statements will execute in reverse order
    }
    fmt.Println("Loop completed.")
}
