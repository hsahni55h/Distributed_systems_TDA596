/*
O_RDONLY : open the file read only
O_WRONLY : open he file write only
O_RDWR   : open the file read-write

Exactly one of these three must be specified

O_APPEND  : append data to file when writing
O_CREATE  : create a new file if none exists
O_EXCL    : used with O_CREATE, file must not exist
O_SYNC    : open for synchronous I/O
O_TRUNC   : runcate regular writeable file when opened

*/


package main

import (
	"fmt"
	"log"
	"os"
)

var pl = fmt.Println

func main() {

	_, err := os.Stat("data.txt")
	if os.IsNotExist(err) {
		pl("File doesn't exist")
	} else {
		f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		if _, err := f.WriteString("13\n"); err != nil {
			log.Fatal(err)
		}
	}

}
