package basics_test

import (
	"bytes"
	"testing"
	"time"

	. "github.com/aubm/golang-codelab/exercises/basics/06-concurrency/01-goroutines"
)

func TestExecute(t *testing.T) {
	var buf *bytes.Buffer

	func() {
		buffers := make(chan *bytes.Buffer)
		go func(done chan *bytes.Buffer) {
			done <- Execute(10)
		}(buffers)

		select {
		case b := <-buffers:
			buf = b
			return
		case <-time.After(5 * time.Second):
			return
		}
	}()

	if buf == nil {
		t.Errorf("Deadline exceeded, Execute ran for too long!")
		return
	}

	expected := "Done\nDone\nDone\nDone\nDone\nDone\nDone\nDone\nDone\nDone\n"
	if actual := buf.String(); actual != expected {
		t.Errorf(`expected buffer output to equal "%v", but got "%v"`, expected, actual)
	}
}
