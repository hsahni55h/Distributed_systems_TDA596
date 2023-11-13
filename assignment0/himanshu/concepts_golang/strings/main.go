package main

import (
	"fmt"
	"strings"
	
)

var pl = fmt.Println

func main(){
	sv1 := "A Word"
	pl(sv1)
	replacer := strings.NewReplacer("A", "Another")
	sv2 := replacer.Replace(sv1)
	pl(sv2)
	pl("Length: ", len(sv2))
	pl("Contains Another: ", strings.Contains(sv2, "Another"))
	pl("O index: ", strings.Index(sv2, "o"))
	pl("Replace: ", strings.Replace(sv2, "o", "0", -1))
	sv3 := "\nSome Words\n"
	sv3 = strings.TrimSpace(sv3)
	pl(sv3)
	pl("Split: ", strings.Split("a-b-c-d", "-"))
	pl("Lower: ", strings.ToLower(sv2))
	pl("Upper: ", strings.ToUpper(sv2))
	pl("Prefix: ", strings.HasPrefix("tacocat", "taco"))
	pl("Suffix: ", strings.HasSuffix("tacocat", "cat"))

}