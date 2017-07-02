package basics_test

import (
	"reflect"
	"testing"

	. "github.com/aubm/golang-codelab/exercises/basics/01-basic-data-types"
)

func TestMustReturnAnInteger(t *testing.T) {
	v := MustReturnAnInteger()
	if typeOfV := reflect.TypeOf(v); typeOfV.Name() != "int" {
		t.Errorf("returned type should be an int, got %v", typeOfV.Name())
	}
}

func TestMustReturnTheSum(t *testing.T) {
	for _, d := range []struct {
		a, b, res int
	}{
		{a: 0, b: 0, res: 0},
		{a: 1, b: 0, res: 1},
		{a: 3, b: 3, res: 6},
		{a: 3, b: 4, res: 7},
	} {
		if v := MustReturnTheSum(d.a, d.b); v != d.res {
			t.Errorf("when a=%v and b=%v, expected %v, got %v", d.a, d.b, d.res, v)
		}
	}
}

func TestMustReturnAFloat(t *testing.T) {
	v := MustReturnAFloat()
	if typeOfV := reflect.TypeOf(v); typeOfV.Name() != "float64" {
		t.Errorf("returned type should be an float64, got %v", typeOfV.Name())
	}
}

func TestMustBeTrue(t *testing.T) {
	if v := MustBeTrue(); v != true {
		t.Errorf("should be true, got %v", v)
	}
}

func TestMustBeFalse(t *testing.T) {
	if v := MustBeFalse(); v != false {
		t.Errorf("should be false, got %v", v)
	}
}

func TestMustBeTrueIfThereAreTenEggsOrMore(t *testing.T) {
	for _, d := range []struct {
		in       int
		expected bool
	}{
		{0, false},
		{2, false},
		{5, false},
		{9, false},
		{10, true},
		{11, true},
		{100, true},
	} {
		if v := MustBeTrueIfThereAreTenEggsOrMore(d.in); v != d.expected {
			t.Errorf("when there are %v eggs, expected %v, got %v", d.in, d.expected, v)
		}
	}
}

func TestMustReturnHello(t *testing.T) {
	expected := "Hello"
	if v := MustReturnHello(); v != expected {
		t.Errorf(`expected "%v", got "%v"`, expected, v)
	}
}

func TestMustReturnWorld(t *testing.T) {
	expected := "world"
	if v := MustReturnWorld(); v != expected {
		t.Errorf(`expected "%v", got "%v"`, expected, v)
	}
}

func TestMustSetToTrueAndReturn(t *testing.T) {
	if v := MustSetToTrueAndReturn(); v != true {
		t.Errorf("should be true, got %v", v)
	}
}
