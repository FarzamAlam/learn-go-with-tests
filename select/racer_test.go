package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compare speed of servers between two urls :", func(t *testing.T) {
		slowServer := makeDelayedServer(2 * time.Second)
		fastServer := makeDelayedServer(1 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("Error : %v ", err)
		}
		if got != want {
			t.Errorf("got %s while want %s", got, want)
		}
	})
	t.Run("Check default timeout : ", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)
		defer server.Close()
		_, err := ConfigurableRacer(server.URL, server.URL, time.Microsecond)
		if err == nil {
			t.Errorf("expected nil but got %v : ", err)
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
