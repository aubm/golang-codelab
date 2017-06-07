package basics_test

import (
	"reflect"
	"testing"

	. "github.com/aubm/golang-codelab/exercises/basics/03-functions/07-pointers"
)

func TestGiveMeAPersonAndIWillRenameItToJohnDoe(t *testing.T) {
	actual := Person{FirstName: "Some", LastName: "Guy"}
	valueOfFunc := reflect.ValueOf(GiveMeAPersonAndIWillRenameItToJohnDoe)
	typeOfFunc := reflect.TypeOf(GiveMeAPersonAndIWillRenameItToJohnDoe)
	personArg := typeOfFunc.In(0)
	in := make([]reflect.Value, 0)
	if personArg.Kind() == reflect.Ptr {
		in = append(in, reflect.ValueOf(&actual))
	} else {
		in = append(in, reflect.ValueOf(actual))
	}
	valueOfFunc.Call(in)

	expected := Person{FirstName: "John", LastName: "Doe"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v to have been renamed to %v, but has not", actual, expected)
	}
}
