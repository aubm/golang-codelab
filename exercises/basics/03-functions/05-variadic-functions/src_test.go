package basics_test

import (
	"bytes"
	"testing"

	. "github.com/aubm/golang-codelab/exercises/basics/03-functions/05-variadic-functions"
)

func TestComputeAverage(t *testing.T) {
	for _, d := range []struct {
		numbers  []float64
		expected float64
	}{
		{numbers: []float64{4, 5.1, 12, 54.12}, expected: 18.805},
		{numbers: []float64{14, 16.5, 18, 12.5}, expected: 15.25},
	} {
		if actual := ComputeAverage(d.numbers...); actual != d.expected {
			t.Errorf("when numbers are %v, expected average to be %v, but got %v", d.numbers, d.expected, actual)
		}
	}
}

func TestWriteEverythingInBufferWithOneLineOfCode(t *testing.T) {
	buf := new(bytes.Buffer)
	words := []interface{}{"My ", "favorite ", "programing ", "language ", "is ", "Golang"}
	expected := "My favorite programing language is Golang"
	WriteEverythingInBufferWithOneLineOfCode(buf, words)
	if actual := buf.String(); actual != expected {
		t.Errorf("expected buffer output to be '%v', but got '%v'", expected, actual)
	}
}
