package basics

import (
	"context"
	"io/ioutil"
	"net/http"
)

// Some urls can be really slow, and again: life is too short.
// Using the context object provided as the first function argument, the user can decide to
// end the function execution early, returning a nil slice of bytes.
// But sometimes, the user can be patient. In that case, the function must wait for fetchUrlContent
// to publish a value in the returned channel, and return that value to the user.
func FetchUrlContentOrCancelOnUserDemand(ctx context.Context, url string) []byte {
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
