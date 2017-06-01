package basics_test

import (
	"testing"

	. "github.com/aubm/golang-codelab/solutions/basics/03-functions/01-simple"
)

func TestRed(t *testing.T) {
	expected := "FF0000"
	actual := Red()
	if actual != expected {
		t.Errorf("expected %v, but got %v", expected, actual)
	}
}

func TestGreen(t *testing.T) {
	expected := "00FF00"
	actual := Green()
	if actual != expected {
		t.Errorf("expected %v, but got %v", expected, actual)
	}
}

func TestBlue(t *testing.T) {
	expected := "0000FF"
	actual := Blue()
	if actual != expected {
		t.Errorf("expected %v, but got %v", expected, actual)
	}
}
