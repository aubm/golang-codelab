package basics

import (
	"bytes"
	"net/http"
	"os"
)

const bytesToRead = 11

func ReadFirstBytesFromHttpResponseBody(response *http.Response) []byte {
	b := make([]byte, bytesToRead)
	response.Body.Read(b)
	return b
}

func ReadFirstBytesFromFile(file *os.File) []byte {
	b := make([]byte, bytesToRead)
	file.Read(b)
	return b
}

func ReadFirstBytesFromBuffer(buffer *bytes.Buffer) []byte {
	b := make([]byte, bytesToRead)
	buffer.Read(b)
	return b
}

// The three functions above execute exactly the same logic for three different objects:
// response.Body, file and buffer.
// Wouldn't it be great if we had a function named ReadFirstBytes, that could take any of those objects as an input?
// Let's do that!
