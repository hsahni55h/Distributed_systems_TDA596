package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {

	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	prime_numbers_array := []int {2,3,5,7,11,13}
	var string_prime_numbers [] string
	for _, i := range prime_numbers_array {
		string_prime_numbers = append(string_prime_numbers, strconv.Itoa(i))
	}

	for _, num := range string_prime_numbers {
		_, err = f.WriteString(num + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scan1 := bufio.NewScanner(f)
	for scan1.Scan(){
		pl("Prime : ", scan1.Text())
	}

	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}

}