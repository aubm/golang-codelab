package basics

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

var mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "My favorite language is Golang!")
}))

func TestReadFirstBytesFromHttpResponseBody(t *testing.T) {
	response, _ := http.Get(mockServer.URL)
	expected := []byte("My favorite")
	if actual := ReadFirstBytesFromHttpResponseBody(response); !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected first bytes to equal %v, but got %v", expected, actual)
	}
}

func TestReadFirstBytesFromFile(t *testing.T) {
	file, _ := os.Open("test_data/some_file.txt")
	expected := []byte("My favorite")
	if actual := ReadFirstBytesFromFile(file); !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected first bytes to equal %v, but got %v", expected, actual)
	}
}

func TestReadFirstBytesFromBuffer(t *testing.T) {
	buf := bytes.NewBuffer([]byte("My favorite language is Golang!"))
	expected := []byte("My favorite")
	if actual := ReadFirstBytesFromBuffer(buf); !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected first bytes to equal %v, but got %v", expected, actual)
	}
}

func TestReadFirstBytes(t *testing.T) {
	expected := []byte("My favorite")

	response, _ := http.Get(mockServer.URL)
	if actual := ReadFirstBytes(response.Body); !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected first bytes to equal %v, but got %v", expected, actual)
	}

	file, _ := os.Open("test_data/some_file.txt")
	if actual := ReadFirstBytes(file); !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected first bytes to equal %v, but got %v", expected, actual)
	}

	buf := bytes.NewBuffer([]byte("My favorite language is Golang!"))
	if actual := ReadFirstBytes(buf); !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected first bytes to equal %v, but got %v", expected, actual)
	}
}
