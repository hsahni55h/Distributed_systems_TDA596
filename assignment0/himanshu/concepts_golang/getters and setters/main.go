/* An identifier (variable, type, function, etc.) in Go is exported if its name starts with an uppercase letter. Exported identifiers are accessible from other packages.

An identifier is unexported if its name starts with a lowercase letter. Unexported identifiers are only accessible within the same package.
*/
package main

import "fmt"

// Person is a simple struct representing a person.
type Person struct {
    firstName string // unexported field
    lastName  string // unexported field
}

// NewPerson is a constructor function for creating a new Person.
func NewPerson(firstName, lastName string) *Person {
    return &Person{firstName: firstName, lastName: lastName}
}

// Getter method for getting the first name.
func (p *Person) GetFirstName() string {
    return p.firstName
}

// Setter method for setting the first name.
func (p *Person) SetFirstName(newFirstName string) {
    p.firstName = newFirstName
}

// Getter method for getting the last name.
func (p *Person) GetLastName() string {
    return p.lastName
}

// Setter method for setting the last name.
func (p *Person) SetLastName(newLastName string) {
    p.lastName = newLastName
}

func main() {
    // Create a new Person
    person := NewPerson("John", "Doe")

    // Use getters
    fmt.Println("First Name:", person.GetFirstName())
    fmt.Println("Last Name:", person.GetLastName())

    // Use setter
    person.SetFirstName("Jane")

    // Verify the change
    fmt.Println("Updated First Name:", person.GetFirstName())
}
