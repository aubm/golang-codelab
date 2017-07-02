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

// This function uses a for/range loop, which is unfortunate because the same thing could be written with only
// one line of code. Can you see how?
// Hint: the type for wordsToPrint is a slice of interface{}. Interfaces are addressed in the further exercises.
// For now, consider interface{} to be equivalent to "any type". So []interface{} means "a slice of any type", either
// strings, integers, booleans, etc...
func WriteEverythingInBufferWithOneLineOfCode(buffer *bytes.Buffer, wordsToPrint []interface{}) {
	buffer.WriteString(fmt.Sprint(wordsToPrint...))
}
