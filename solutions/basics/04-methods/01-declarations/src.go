package basics

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) GetFullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}
