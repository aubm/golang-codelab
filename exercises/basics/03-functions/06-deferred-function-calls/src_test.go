package basics_test

import (
	"testing"

	. "github.com/aubm/golang-codelab/exercises/basics/03-functions/06-deferred-function-calls"
)

func TestAlwaysInvokeCallbackBeforeExitButOnlyWriteItOnce(t *testing.T) {
	for _, d := range []struct {
		in, expected int
	}{
		{in: -1, expected: 0},
		{in: 9, expected: 10},
		{in: 49, expected: 50},
		{in: 99, expected: 100},
	} {
		changeMe := 200
		cb := func(v int) {
			changeMe = v
		}
		AlwaysInvokeCallbackBeforeExitButOnlyWriteItOnce(d.in, cb)
		if changeMe != d.expected {
			t.Errorf("when someNumber is %v, expected callback to have been given %v, but got %v", d.in, d.expected, changeMe)
		}
	}
}
