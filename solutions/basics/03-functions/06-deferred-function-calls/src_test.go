package basics_test

import (
	"testing"

	. "github.com/aubm/golang-codelab/solutions/basics/03-functions/06-deferred-function-calls"
)

func TestAlwaysInvokeCallbackBeforeExitButOnlyWriteItOnce(t *testing.T) {
	values := []int{-1, 9, 49, 99}
	for _, v := range values {
		changeMe := false
		cb := func() {
			changeMe = true
		}
		AlwaysInvokeCallbackBeforeExitButOnlyWriteItOnce(v, cb)
		if !changeMe {
			t.Errorf("when someNumber is %v, expected callback to have been called, but it was not", v)
		}
	}
}
