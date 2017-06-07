package basics

import (
	"bytes"
	"time"
)

// Execute really takes too much time, and life is too short!
// Maybe if we ran all the long processes concurrently, it would be much faster.
// Let's do that!
// Beware, all the processes must end before that Execute returns.
func Execute(nbProcess int) *bytes.Buffer {
	buf := new(bytes.Buffer)
	for i := 0; i < nbProcess; i++ {
		someReallyLongProcess(buf)
	}
	return buf
}

func someReallyLongProcess(buf *bytes.Buffer) {
	time.Sleep(1 * time.Second)
	buf.WriteString("Done\n")
}
