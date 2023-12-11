/* A rune is a built-in type that represents a Unicode code point. 
It is the equivalent of a character in other programming languages, but it is more general because it can represent any Unicode character, not just ASCII characters.
*/

package main

import (
	
	"fmt"
	"unicode/utf8"
)

var pl = fmt.Println

func main() {
	rStr := "abcdefg"
	pl("Rune Count:", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr{
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}
}