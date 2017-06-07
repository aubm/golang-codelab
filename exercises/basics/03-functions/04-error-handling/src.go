package basics

// PingURL must perform an HTTP GET request to the provided url.
// If the request fails, PingURL must return an error.
// If the request succeeds, PingURL must return nil in order to indicate that the url is alive.
// PingURL must not look into the detail of the response object, thus an url which responds
// with a HTTP status code 500 must not result as an error.
func PingURL(targetURL string) error {
}

// ReadURLContents must perform an HTTP GET request to the provided url.
// It must return the response bytes and a nil error if the request succeeds.
// It must return a nil slice of bytes and an error if one of the following happens:
// - the HTTP request fails
// - the server responds with a status code that is different from 200
// - the response body can't be read for some reason
func ReadURLContents(targetURL string) ([]byte, error) {
}
