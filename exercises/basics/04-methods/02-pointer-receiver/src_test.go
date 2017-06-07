package basics_test

import (
	"testing"

	. "github.com/aubm/golang-codelab/exercises/basics/04-methods/02-pointer-receiver"
)

func TestPersonGetOlder(t *testing.T) {
	for _, d := range []struct {
		originalAge, expected int
	}{
		{originalAge: 10, expected: 11},
		{originalAge: 21, expected: 22},
	} {
		person := &Person{Age: d.originalAge}
		person.GetOlder()
		if actual := person.Age; actual != d.expected {
			t.Errorf("when original age is %v, expected age to equal %v, but got %v", d.originalAge, d.expected, actual)
		}
	}
}
