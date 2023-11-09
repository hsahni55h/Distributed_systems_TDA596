package main

import (
	"fmt"
	"reflect"	
)

var pl = fmt.Println
func main() {

	// variables are mutable as long as the data type remains the same
	var name string
	name = "Himanshu"

	var number uint16 = 260

	fmt.Println(name)
	fmt.Println(number)
	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.14))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("Hello"))
}
