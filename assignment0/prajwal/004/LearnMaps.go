package main

import "fmt"

func main() {
	/*
		Syntax:
		var myMap map [keyType]valType
	*/
	var heros map[string]string
	
	heros = make(map[string]string)				// method 1
	villans := make(map[string]string)			// method 2

	heros["Batman"] = "Bruce Wayne"
	heros["Superman"] = "Clark Kent"
	heros["Spiderman"] = "Peter Parker"
	heros["Flash"] = "Barry Allen"

	villans["Lex Luther"] = "Lex Luther"

	health := map[string]int{					// method 3
		"Batman": 100,
		"Superman": 100,
		"Spiderman": 100,
		"Flash": 100,
		"Lex Luther": 500,
	}

	characters := heros
	for key, val := range villans {
		characters[key] = val
	} 
	
	for key, value := range characters {
		fmt.Println("Real name of", key, "is", value, "with health =", health[key])
	}

	_, ok := villans["Joker"]		// key not in the map returns false
	fmt.Println("Is Joker included this universe?:", ok)

	delete(heros, "Spiderman")
	fmt.Println("Spiderman removed since its not a part of DC unvierse")
	for key, value := range heros {
		fmt.Println("Real name of", key, "is", value, "with health =", health[key])
	}
}