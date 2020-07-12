package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
	}
	server := &PlayerServer{&store}
	t.Run("It returns accepted on POST : ", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertResponseCode(t, response.Code, http.StatusAccepted)
	})
}

func TestGetPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := &PlayerServer{&store}
	t.Run("Returns Pepper's Score : ", func(t *testing.T) {
		request, _ := newRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		got := response.Body.String()
		want := "20"
		assertResponseBody(t, got, want)
	})
	t.Run("Returns Floyd's Score : ", func(t *testing.T) {
		request, _ := newRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		got := response.Body.String()
		want := "10"
		assertResponseBody(t, got, want)
	})
	t.Run("returns 404 on missing players ", func(t *testing.T) {
		request, _ := newRequest("Apollo")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := response.Code
		want := 404
		assertResponseCode(t, got, want)
	})
}
func assertResponseCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d while want %d", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q while want %q ", got, want)
	}
}

func newRequest(name string) (*http.Request, error) {
	return http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
}
