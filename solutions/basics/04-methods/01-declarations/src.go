package basics

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

// In order for the tests to pass in this package to pass, you must add a new method to the type Person.
// The must be named GetFullName and return a string.
// The string is the full name of the person, for example if the first name is "Go" and the last name if "Lang",
// the full name is "Go Lang".

func (p Person) GetFullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}
