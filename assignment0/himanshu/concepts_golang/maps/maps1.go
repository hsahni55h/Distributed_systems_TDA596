package main

import "fmt"

func main() {
	// Creating and initializing maps
	// var name_of_the_map map [keyType]valyetype
	
	fruitsColor := map[string]string{
		"apple":  "red",
		"banana": "yellow",
		"cherry": "red",
	}

	// Accessing and modifying map elements
	appleColor := fruitsColor["apple"]
	fmt.Println("Color of an apple:", appleColor)
	fruitsColor["banana"] = "green"
	fmt.Println("Updated Color of a banana:", fruitsColor["banana"])

	// Checking if a key exists in a map
	color, exists := fruitsColor["mango"]
	if exists {
		fmt.Println("Color of a mango:", color)
	} else {
		fmt.Println("Mango not found in the map")
	}

	// Deleting elements from the map
	delete(fruitsColor, "banana")
	// Attempting to delete a non-existent element
	delete(fruitsColor, "mango")

	// Iterating over map elements
	fmt.Println("Iterating over map elements:")
	for fruit, color := range fruitsColor {
		fmt.Printf("%s is %s\n", fruit, color)
	}

	// Maps with different value types
	studentCourses := map[string][]string{
		"Alice":   {"Math", "Science"},
		"Bob":     {"History"},
		"Charlie": {"English", "Art"},
	}

	userDetails := map[string]map[string]string{
		"Alice": {"age": "25", "city": "New York"},
		"Bob":   {"age": "30", "city": "Chicago"},
	}

	// Accessing and modifying map values
	studentCourses["Alice"] = append(studentCourses["Alice"], "Music")
	userDetails["Bob"]["city"] = "Los Angeles"

	fmt.Println("Student Courses:", studentCourses)
	fmt.Println("User Details:", userDetails)

}
