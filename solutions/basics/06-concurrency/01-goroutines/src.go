package basics

import (
	"bytes"
	"sync"
	"time"
)

// Execute really takes too much time, and life is too short!
// Maybe if we ran all the long processes concurrently, it would be much faster.
// Let's do that!
// Beware, all the processes must end before that Execute returns.

var m sync.Mutex

func Execute(nbProcess int) *bytes.Buffer {
	buf := new(bytes.Buffer)
	wg := &sync.WaitGroup{}
	wg.Add(nbProcess)
	for i := 0; i < nbProcess; i++ {
		go func() {
			someReallyLongProcess(buf)
			wg.Done()
		}()
	}
	wg.Wait()
	return buf
}

func someReallyLongProcess(buf *bytes.Buffer) {
	time.Sleep(1 * time.Second)
	m.Lock()
	defer m.Unlock()
	buf.WriteString("Done\n")
}
