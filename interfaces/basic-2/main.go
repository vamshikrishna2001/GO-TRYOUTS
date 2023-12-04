package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

// Implementing the String() method for the Person type
func (p Person) String() string {
	return fmt.Sprintf("%s camila mendes %s", p.FirstName, p.LastName)
}

func main() {
	// Creating a Person
	person := Person{"John", "Doe"}

	// Using the String() method implicitly through fmt.Println
	fmt.Println(person) // Output: John Doe
}
