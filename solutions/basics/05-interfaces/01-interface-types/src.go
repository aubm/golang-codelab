package basics

import (
	"bytes"
	"io"
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

func ReadFirstBytes(r io.Reader) []byte {
	b := make([]byte, bytesToRead)
	r.Read(b)
	return b
}
