package basics_test

import (
	"reflect"
	"testing"

	. "github.com/aubm/golang-codelab/exercises/basics/02-composite-types"
)

func TestMustReturnAnArrayOfThreeIntegers(t *testing.T) {
	v := MustReturnAnArrayOfThreeIntegers()
	expectedType := "[3]int"
	if actualType := reflect.TypeOf(v).String(); actualType != expectedType {
		t.Errorf("returned type should be %v, got %v", expectedType, actualType)
	}
}

func TestMustReturnTheFirstElementOfSlice(t *testing.T) {
	for _, d := range []struct {
		in       []int
		expected int
	}{
		{in: []int{1, 3, 5}, expected: 1},
		{in: []int{5, 3, 1}, expected: 5},
		{in: []int{3, 5, 1}, expected: 3},
	} {
		if actual := MustReturnTheFirstElementOfSlice(d.in); actual != d.expected {
			t.Errorf("when input is %v, expected output to equal %v, but got %v", d.in, d.expected, actual)
		}
	}
}

func TestMustReturnTheLastElementOfSlice(t *testing.T) {
	for _, d := range []struct {
		in       []int
		expected int
	}{
		{in: []int{1, 3, 5}, expected: 5},
		{in: []int{5, 3, 1}, expected: 1},
		{in: []int{3, 5, 1}, expected: 1},
	} {
		if actual := MustReturnTheLastElementOfSlice(d.in); actual != d.expected {
			t.Errorf("when input is %v, expected output to equal %v, but got %v", d.in, d.expected, actual)
		}
	}
}

func TestMustReturnTheGreatestElementOfSlice(t *testing.T) {
	for _, d := range []struct {
		in       []int
		expected int
	}{
		{in: []int{1, 3, 5}, expected: 5},
		{in: []int{5, 3, 1}, expected: 5},
		{in: []int{3, 5, 1}, expected: 5},
	} {
		if actual := MustReturnTheGreatestElementOfSlice(d.in); actual != d.expected {
			t.Errorf("when input is %v, expected output to equal %v, but got %v", d.in, d.expected, actual)
		}
	}
}

func TestMustReturnTheNumberOfWords(t *testing.T) {
	for _, d := range []struct {
		text     string
		expected int
	}{
		{text: "", expected: 0},
		{text: "Lorem", expected: 1},
		{text: "Lorem ipsum dolor sit amet", expected: 5},
		{text: "Lorem ipsum dolor sit amet consectetur adipiscing elit", expected: 8},
	} {
		if actual := MustReturnTheNumberOfWords(d.text); actual != d.expected {
			t.Errorf(`when input text is "%v", expected output to equal %v, but got %v`, d.text, d.expected, actual)
		}
	}
}

func TestMustAppendTheWordGolang(t *testing.T) {
	words := []string{"My", "favorite", "programing", "language", "is"}
	expected := []string{"My", "favorite", "programing", "language", "is", "Golang"}
	if actual := MustAppendTheWordGolang(words); !reflect.DeepEqual(actual, expected) {
		t.Errorf("when input is %v, expected output to equal %v, but got %v", words, expected, actual)
	}
}

func TestMustReturnTheLevelKeyOrOneByDefault(t *testing.T) {
	for _, d := range []struct {
		stats    map[string]int
		expected int
	}{
		{stats: map[string]int{"hits": 0}, expected: 1},
		{stats: map[string]int{"hits": 45, "level": 5}, expected: 5},
		{stats: map[string]int{"hits": 654, "level": 9}, expected: 9},
	} {
		if actual := MustReturnTheLevelKeyOrOneByDefault(d.stats); actual != d.expected {
			t.Errorf(`when stats are "%v", expected level to equal %v, but got %v`, d.stats, d.expected, actual)
		}
	}
}

func TestMustSortAndReturn(t *testing.T) {
	for _, d := range []struct {
		in, expected []int
	}{
		{in: []int{1, 3, 5}, expected: []int{1, 3, 5}},
		{in: []int{5, 3, 1}, expected: []int{1, 3, 5}},
		{in: []int{3, 5, 1}, expected: []int{1, 3, 5}},
	} {
		if actual := MustSortAndReturn(d.in); !reflect.DeepEqual(actual, d.expected) {
			t.Errorf("when input is %v, expected output to equal %v, but got %v", d.in, d.expected, actual)
		}
	}
}

func TestMustReturn21YearsOldJohnDoe(t *testing.T) {
	actual := MustReturn21YearsOldJohnDoe()
	expected := Person{FirstName: "John", LastName: "Doe", Age: 21}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, but got %v", expected, actual)
	}
}

func TestMustReturnAListOfPerson(t *testing.T) {
	jsonListOfPerson := []byte(`[{"first_name": "John", "last_name": "Doe", "age": 21}, {"first_name": "Jane", "last_name": "Doe", "age": 19}]`)
	expected := []Person{{FirstName: "John", LastName: "Doe", Age: 21}, {FirstName: "Jane", LastName: "Doe", Age: 19}}
	if actual := MustReturnAListOfPerson(jsonListOfPerson); !reflect.DeepEqual(actual, expected) {
		t.Errorf("when json is %s, expected %v, but got %v", jsonListOfPerson, expected, actual)
	}

	defer func() {
		recover()
	}()
	invalidJSON := []byte(`{"first_name"`)
	MustReturnAListOfPerson(invalidJSON)
	t.Errorf("when json is invalid, expected function to panic, but it did not")
}
