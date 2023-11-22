package main

import "fmt"

func main() {
	// Creating and initializing maps
	// var name_of_the_map map [keyType]valyetype

	var heroes map[string]string
	heroes = make(map[string]string)

	villians := make(map[string]string)

	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clarke Kent"
	heroes["The Flash"] = "Barry Allen"

	villians["Lex Luther"] = "lex Luther"

	superPets := map[int]string{1: "krypto", 2: "bat Hound"}

	
	fmt.Printf("batman is %v\n", heroes["Batman"])
	fmt.Println("Chip : ", superPets[3])
	_, ok := superPets[3]
	fmt.Println("is there a 3rd pet :", ok)
	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}

}
