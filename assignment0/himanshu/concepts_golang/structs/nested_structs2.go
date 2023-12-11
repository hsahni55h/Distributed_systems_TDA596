package main

import "fmt"

type contact struct {
	fName string
	lName string
	phone string
}

type business struct {
	name    string
	address string
	contact
}

func (b business) info() {
	fmt.Printf("Contact at %s is %s %s", b.name, b.contact.fName, b.contact.lName)
}

func main() {
	con1 := contact{
		"James",
		"Wang",
		"555-1212",
	}

	bus1 := business{
		"ABC Plumbing",
		"234 North St",
		con1,
	}
	bus1.info()
}
