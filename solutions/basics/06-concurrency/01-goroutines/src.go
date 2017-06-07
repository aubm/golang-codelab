package basics

import (
	"bytes"
	"sync"
	"time"
)

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
	buf.WriteString("Done\n")
}
