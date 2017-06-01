package basics

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func PingURL(targetURL string) error {
	if _, err := http.Get(targetURL); err != nil {
		return err
	}
	return nil
}

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
