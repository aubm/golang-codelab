package basics_test

import (
	"testing"

	. "github.com/aubm/golang-codelab/exercises/basics/04-methods/01-declarations"
)

func TestPersonGetFullName(t *testing.T) {
	for _, d := range []struct {
		person   Person
		expected string
	}{
		{person: Person{FirstName: "John", LastName: "Doe"}, expected: "John Doe"},
		{person: Person{FirstName: "Jane", LastName: "Doe"}, expected: "Jane Doe"},
		{person: Person{FirstName: "Apple", LastName: "Seed"}, expected: "Apple Seed"},
	} {
		if actual := d.person.GetFullName(); actual != d.expected {
			t.Errorf("when person is %v, expected full name to equal %v, but got %v", d.person, d.expected, actual)
		}
	}
}
