package basics_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	. "github.com/aubm/golang-codelab/exercises/basics/03-functions/04-error-handling"
)

var operationalServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Go!"))
}))

var errorServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Hello error!"))
}))

func TestPingURL(t *testing.T) {
	for _, d := range []struct {
		targetURL     string
		errorExpected bool
	}{
		{targetURL: operationalServer.URL, errorExpected: false},
		{targetURL: errorServer.URL, errorExpected: false},
		{targetURL: "some invalid url", errorExpected: true},
	} {
		err := PingURL(d.targetURL)
		if err != nil && d.errorExpected == false {
			t.Errorf("when url is %v, expected no errors, but got %v", d.targetURL, err)
		}
		if err == nil && d.errorExpected == true {
			t.Errorf("when url is %v, expected an error, but got none", d.targetURL)
		}
	}
}

func TestReadURLContents(t *testing.T) {
	for _, d := range []struct {
		targetURL     string
		errorExpected bool
		expectedBytes []byte
	}{
		{targetURL: operationalServer.URL, errorExpected: false, expectedBytes: []byte("Hello Go!")},
		{targetURL: errorServer.URL, errorExpected: true, expectedBytes: nil},
		{targetURL: "some invalid url", errorExpected: true, expectedBytes: nil},
	} {
		b, err := ReadURLContents(d.targetURL)
		if err != nil && d.errorExpected == false {
			t.Errorf("when url is %v, expected no errors, but got %v", d.targetURL, err)
		}
		if err == nil && d.errorExpected == true {
			t.Errorf("when url is %v, expected an error, but got none", d.targetURL)
		}
		if !reflect.DeepEqual(b, d.expectedBytes) {
			t.Errorf("when url is %v, expected returned value to equal %s, but go %s", d.targetURL, d.expectedBytes, b)
		}
	}
}
