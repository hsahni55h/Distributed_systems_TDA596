package main

import (
	"fmt"
)

var pl = fmt.Println

func add(op1 int32, op2 int32) int32 {
	return op1 + op2
}

func div(op1 float32, op2 float32) (ans float32, err error) {
	if op2 == 0 {
		return 0, fmt.Errorf("Can't divide by zero")
	}
	return op1 / op2, nil
}

func getSum2(ops ...int32) int32 {
	var sum int32 = 0
	for _, op := range ops {
		sum += op
	}
	return sum
}

func getArraySum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func changeValue2(myPtr *int) *int {
	*myPtr = 12
	return myPtr
}

func avg(nums []float32) float32 {
	var sum float32 = 0.
	for _, num := range nums {
		sum += num
	}
	return sum / float32(len(nums))
}

func main() {
	pl("Hello World")
	pl("2 + 3 =", add(2, 3), "\n3 + 2 =", add(3, 2))

	op1, err := div(3, 2)
	pl("3 / 2 =", op1, "and err =", err)

	op2, err := div(3, 0)
	pl("3 / 0 =", op2, "and err =", err)

	pl("1 + 2 + 3 =", getSum2(1, 2, 3))
	pl("1 + 2 + 3 + 4 + 5 =", getSum2(1, 2, 3, 4, 5))

	var arr1 = []int{1, 2, 3}
	pl("1 + 2 + 3 =", getArraySum(arr1))

	var arr2 = []int{1, 2, 3, 4, 5}
	pl("1 + 2 + 3 + 4 + 5 =", getArraySum(arr2))

	pl("old value is:", arr1[0])
	pl("new value is:", *changeValue2(&arr1[0]))

	var sl_as_arrf = []float32{1., 2., 3., 4., 5.}
	sl_as_arrf = append(sl_as_arrf, 6.)
	pl("avg of", sl_as_arrf, "is:", avg(sl_as_arrf))

	slf := make([]float32, 5)
	slf = append(slf, 6)
	pl("avg of", slf, "is:", avg(slf))

	// since avg() is written for slices
	// var arrf = [5]float32{1., 2., 3., 4., 5.}
	// pl("avg of", arrf, "is:", avg(arrf))

}
