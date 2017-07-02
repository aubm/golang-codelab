package basics

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// PingURL must perform an HTTP GET request to the provided url.
// If the request fails, PingURL must return an error.
// If the request succeeds, PingURL must return nil in order to indicate that the url is alive.
// PingURL must not look into the detail of the response object, thus an url which responds
// with a HTTP status code 500 must not result as an error.
func PingURL(targetURL string) error {
	if _, err := http.Get(targetURL); err != nil {
		return err
	}
	return nil
}

// ReadURLContents must perform an HTTP GET request to the provided url.
// It must return the response bytes and a nil error if the request succeeds.
// It must return a nil slice of bytes and an error if one of the following happens:
// - the HTTP request fails
// - the server responds with a status code that is different from 200
// - the response body can't be read for some reason
// Hint: you can read the content of the HTTP response body with ioutil.ReadAll https://golang.org/pkg/io/ioutil/#ReadAll
func ReadURLContents(targetURL string) ([]byte, error) {
	resp, err := http.Get(targetURL)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code is %v", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body failed: %v", err)
	}

	return b, nil
}
