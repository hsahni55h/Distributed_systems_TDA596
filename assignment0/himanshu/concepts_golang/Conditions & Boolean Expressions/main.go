package main

import "fmt"

func main() {
	x := 5
	y := 6.5
	str1 := "Hello"
	str2 := "World"
	str3 := "Hello"

	// Equal to (==)
	equal := x == 5
	fmt.Printf("x == 5: %t\n", equal)

	// Not equal to (!=)
	notEqual := x != 6
	fmt.Printf("x != 6: %t\n", notEqual)

	// Greater than (>)
	greaterThan := y > 6.0
	fmt.Printf("y > 6.0: %t\n", greaterThan)

	// Less than (<)
	lessThan := x < 10
	fmt.Printf("x < 10: %t\n", lessThan)

	// Greater than or equal to (>=)
	greaterThanOrEqual := x >= 5
	fmt.Printf("x >= 5: %t\n", greaterThanOrEqual)

	// Less than or equal to (<=)
	lessThanOrEqual := y <= 6.5
	fmt.Printf("y <= 6.5: %t\n", lessThanOrEqual)

	// String Equality
	stringEqual := str1 == str2
	fmt.Printf("str1 == str2: %t\n", stringEqual)

	// String Inequality
	stringNotEqual := str1 != str3
	fmt.Printf("str1 != str3: %t\n", stringNotEqual)

	// Combining Conditions (AND &&)
	condition1 := x > 4
	condition2 := y > 6.0
	combinedCondition := condition1 && condition2
	fmt.Printf("x > 4 && y > 6.0: %t\n", combinedCondition)

	// Combining Conditions (OR ||)
	condition3 := x < 3
	condition4 := y < 6.0
	combinedCondition2 := condition3 || condition4
	fmt.Printf("x < 3 || y < 6.0: %t\n", combinedCondition2)

	// Negating a Condition (NOT !)
	negatedCondition := !(x > 4)
	fmt.Printf("!(x > 4): %t\n", negatedCondition)
}
