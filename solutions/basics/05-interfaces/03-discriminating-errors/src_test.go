package basics_test

import (
	"errors"
	"testing"

	. "github.com/aubm/golang-codelab/solutions/basics/05-interfaces/03-discriminating-errors"
)

type MockProductFinder struct{}

func (m MockProductFinder) FindOneByID(id string) (*Product, error) {
	switch id {
	case "is invalid":
		return nil, InvalidIDError{ID: id}
	case "is the id of an existing product":
		return &Product{}, nil
	case "does not exist":
		return nil, ProductNotFoundError{ID: id}
	default:
		return nil, errors.New("something bad happened")
	}
}

func TestDoesProductExist(t *testing.T) {
	finder := MockProductFinder{}
	for _, d := range []struct {
		id, expected string
	}{
		{id: "is invalid", expected: "invalid id"},
		{id: "is the id of an existing product", expected: "yes"},
		{id: "does not exist", expected: "no"},
		{id: "is ok but some random error happened", expected: "cannot tell"},
	} {
		if actual := DoesProductExist(finder, d.id); actual != d.expected {
			t.Errorf(`when id %s, expected output to equal "%s", but got "%s"`, d.id, d.expected, actual)
		}
	}
}
