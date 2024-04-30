package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayServer(time.Duration(20 * time.Millisecond))
		fastServer := makeDelayServer(time.Duration(0 * time.Millisecond))

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got, _ := Racer(slowServer.URL, fastServer.URL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayServer(time.Duration(7 * time.Millisecond))
		serverB := makeDelayServer(time.Duration(10 * time.Millisecond))

		defer serverA.Close()
		defer serverB.Close()

		_, err := ConfigurableRacer(serverB.URL, serverA.URL, time.Duration(5*time.Millisecond))

		if err == nil {
			t.Errorf("expected an error but didn`t get it :(")
		}
	})

}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
