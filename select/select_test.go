package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayServer(20 * time.Millisecond)
		fastServer := makeDelayServer(0 * time.Millisecond)

		defer slowServer.Close() //defer wykona te funkcje na koncu
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		// slowURL := "http://www.facebook.com"
		// fastURL := "http://www.quii.dev"

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayServer(25 * time.Second)
		// serverB := makeDelayServer(12 * time.Second)

		defer server.Close()
		// defer serverB.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("Expected an error but didnt get one")
		}
		// server := makeDelayedServer(25 * time.Millisecond)

		// defer server.Close()
		// _, err := ConfigurableRacer(server.URL, )
	})

}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
