package main

import (
	stuff "example/project/my_package"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello", stuff.Name)
	intArr := []int{1, 2, 3, 4, 5}
	strArr := stuff.IntArrToStrArr(intArr)
	fmt.Println(strArr)
	fmt.Println(reflect.TypeOf(strArr))
}