package basics_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	. "github.com/aubm/golang-codelab/solutions/basics/06-concurrency/03-cancellation"
)

var (
	sleepTimeBeforeServerRespond time.Duration
	serverResponse               = []byte("Hello Golang!")
)

var mockHttpServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	time.Sleep(sleepTimeBeforeServerRespond)
	w.Write(serverResponse)
}))

func TestFetchUrlContentOrCancelOnUserDemandWhenTheServerRespondsRightAway(t *testing.T) {
	sleepTimeBeforeServerRespond = 0
	ctx := context.Background()
	actual := FetchUrlContentOrCancelOnUserDemand(ctx, mockHttpServer.URL)
	if !reflect.DeepEqual(actual, serverResponse) {
		t.Errorf(`when the server responds right away, expected returned bytes to equal "%s", but got "%s"`, serverResponse, actual)
	}
}

func TestFetchUrlContentOrCancelOnUserDemandWhenTheUserIsPatient(t *testing.T) {
	sleepTimeBeforeServerRespond = 2 * time.Second
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	actual := FetchUrlContentOrCancelOnUserDemand(ctx, mockHttpServer.URL)
	if !reflect.DeepEqual(actual, serverResponse) {
		t.Errorf(`when the user is patient, expected returned bytes to equal "%s", but got "%s"`, serverResponse, actual)
	}
}

func TestFetchUrlContentOrCancelOnUserDemandWhenTheUserIsVeryPatient(t *testing.T) {
	sleepTimeBeforeServerRespond = 5 * time.Second
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(7*time.Second))
	actual := FetchUrlContentOrCancelOnUserDemand(ctx, mockHttpServer.URL)
	if !reflect.DeepEqual(actual, serverResponse) {
		t.Errorf(`when the user is very patient, expected returned bytes to equal "%s", but got "%s"`, serverResponse, actual)
	}
}

func TestFetchUrlContentOrCancelOnUserDemandWhenTheUserIsNotPatient(t *testing.T) {
	sleepTimeBeforeServerRespond = 30 * time.Second
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	actual := FetchUrlContentOrCancelOnUserDemand(ctx, mockHttpServer.URL)
	if actual != nil {
		t.Errorf(`when the user is not patient, expected returned bytes to be nil, but got "%s"`, actual)
	}
}
