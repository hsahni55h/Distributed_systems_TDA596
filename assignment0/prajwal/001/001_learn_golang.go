package main

import (
	"fmt"
	"log"
	"os"
)

var pl = fmt.Println

func main() {
	fname := "data.txt"
	_, err := os.Stat(fname)
	if err != nil {
		// if command line args exist then write to file else return err
		if len(os.Args) > 1 {
			// create a file
			file, err := os.Create(fname)
			if err != nil {
				log.Fatal(err)
				return
			}
			defer file.Close()

			for _, data := range os.Args[1:] {
				file.WriteString(data + "\n")
				// idata, _ := strconv.Atoi(data)
				// pl(idata)
				// file.Write([]byte(strconv.Itoa(idata)))
			}
		} else {
			err := fmt.Errorf("Please enter few integer as args.")
			log.Fatal(err)
		}
		return
	}

	// open the file in reading mode
	file, err := os.OpenFile(fname, os.O_RDONLY, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	// read the file
	data := make([]byte, 4*10)
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
		return
	}

	// output the data
	pl(data)
	pl(string(data))
}
