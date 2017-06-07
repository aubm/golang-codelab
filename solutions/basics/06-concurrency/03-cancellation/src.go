package basics

import (
	"context"
	"io/ioutil"
	"net/http"
)

func FetchUrlContentOrCancelOnUserDemand(ctx context.Context, url string) []byte {
	select {
	case b := <-fetchUrlContent(url):
		return b
	case <-ctx.Done():
		return nil
	}
}

func fetchUrlContent(url string) chan []byte {
	c := make(chan []byte)
	go func() {
		resp, _ := http.Get(url)
		defer resp.Body.Close()
		b, _ := ioutil.ReadAll(resp.Body)
		c <- b
	}()
	return c
}
