package stuff

import (
	"strconv"
)

var Name string = "Derek Banas"

func IntArrToStrArr(intArr []int) []string {
	var strArr []string
	for _, num := range intArr {
		strArr = append(strArr, strconv.Itoa(num))
	}
	return strArr
}

