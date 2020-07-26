package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := &PlayerServer{&store}
	t.Run("Floyd Request : ", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"
		assertResponseCode(t, response.Code, http.StatusOK)
		assertResponse(t, got, want)
	})
	t.Run("Pepper Request : ", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"
		assertResponseCode(t, response.Code, http.StatusOK)
		assertResponse(t, got, want)
	})
	t.Run("Apollo 404 Request : ", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound
		assertResponseCode(t, got, want)
	})
}

func assertResponseCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d while want %d", got, want)
	}
}

func assertResponse(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q while want %q", got, want)
	}
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)
	return request
}
