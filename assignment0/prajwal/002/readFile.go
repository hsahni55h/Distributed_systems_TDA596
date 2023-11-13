package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ReadStats(fname string) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err.Error()) // equivalent to Println() followed by exit()
	}
	defer file.Close()

	finfo, err := file.Stat()
	if err != nil {
		log.Fatal(err.Error()) // equivalent to Println() followed by exit()
	}

	fmt.Println("File Name:", finfo.Name())
	fmt.Println("Size:", finfo.Size())
	fmt.Println("Mode:", finfo.Mode())
	fmt.Println("Modification Time:", finfo.ModTime())
	fmt.Println("Sys??:", finfo.Sys())
}

func ReadWholeFile(fname string) {
	// data, err := ioutil.ReadFile(fname) // deprecated in Go 1.16 simply calls ReadFile()
	data, err := os.ReadFile(fname)
	if err != nil {
	}
	fmt.Println(string(data))
}

func ReadByLine(fname string) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func ReadByWord(fname string) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func ReadNBytesAtOnce(fname string, size int64) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// finfo, _ := file.Stat()
	// if finfo.Size() > size {
	// 	log.Fatal(fmt.Errorf("buffer size is too small."))
	// }

	buf := make([]byte, size)

	for {
		totalRead, err := file.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		fmt.Println(string(buf[:totalRead]))
	}

	// scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanBytes)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
}

/*
	-------------
	CONFIG FORMAT
	-------------
	key1=value1
	key2=value2
	key3=value3
	...
*/
func ReadConfig(fname string) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, "=")
		fmt.Println("key =", items[0], "\b, value =", items[1])
	}
}

func main() {
	fname := "data.txt"
	// ReadStats(fname)
	// ReadWholeFile(fname)
	ReadByLine(fname)
	// ReadByWord(fname)
	// ReadNBytesAtOnce(fname, 4)
	ReadConfig("myconfig.cfg")
}
