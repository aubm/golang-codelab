package basics_test

import (
	"reflect"
	"testing"

	. "github.com/aubm/golang-codelab/exercises/basics/04-methods/03-encapsulation"
)

func TestSetAge(t *testing.T) {
	person := NewPerson(10)
	typeOfPerson := reflect.TypeOf(*person)
	_, fieldExists := typeOfPerson.FieldByName("Age")
	if fieldExists {
		t.Error(`Type Person has a property Age that is settable.
		The age of a person should only be determined once when calling NewPerson.
		Or mutated when calling Person.GetOlder().`)
	}
}
