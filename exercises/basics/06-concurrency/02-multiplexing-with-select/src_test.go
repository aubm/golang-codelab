package basics_test

import (
	"testing"
	"time"

	. "github.com/aubm/golang-codelab/exercises/basics/06-concurrency/02-multiplexing-with-select"
)

func TestStartSlate(t *testing.T) {
	var actual Slate
	go func() {
		actual = StartSlate()
	}()

	go func() {
		Biers <- 1
		Margarita <- 2
		Manhattan <- 3
		Wine <- 1
		Wine <- 1
		Manhattan <- 1
		Biers <- 2
		CanWeHaveTheCheckPlease <- true
	}()

	time.Sleep(2 * time.Second)

	expected := Slate{Biers: 3, Margarita: 2, Manhattan: 4, Wine: 2}
	if actual != expected {
		t.Errorf("expected slate to equal %v, but got %v", expected, actual)
	}
}

func TestAnotherStartSlate(t *testing.T) {
	var actual Slate
	go func() {
		actual = StartSlate()
	}()

	go func() {
		Manhattan <- 5
		Margarita <- 3
		Manhattan <- 1
		Wine <- 3
		Biers <- 5
		Wine <- 3
		Biers <- 1
		CanWeHaveTheCheckPlease <- true
	}()

	time.Sleep(2 * time.Second)

	expected := Slate{Biers: 6, Margarita: 3, Manhattan: 6, Wine: 6}
	if actual != expected {
		t.Errorf("expected slate to equal %v, but got %v", expected, actual)
	}
}
