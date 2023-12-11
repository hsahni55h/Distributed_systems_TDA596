/* A rune is a built-in type that represents a Unicode code point. 
It is the equivalent of a character in other programming languages, but it is more general because it can represent any Unicode character, not just ASCII characters.
*/

package main

import (
	
	"fmt"
	"time"
)

var pl = fmt.Println

func main() {
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(),now.Second())
}