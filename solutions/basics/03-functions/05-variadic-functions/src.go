package basics

import (
	"bytes"
	"fmt"
)

func ComputeAverage(numbers ...float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func WriteEverythingInBufferWithOneLineOfCode(buffer *bytes.Buffer, wordsToPrint []interface{}) {
	// This function must be one line long, not one more
	buffer.WriteString(fmt.Sprint(wordsToPrint...))
}
