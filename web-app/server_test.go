package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayers(t *testing.T) {
	t.Run("Returns Pepper's Score : ", func(t *testing.T) {
		request, _ := newRequest("Pepper")
		response := httptest.NewRecorder()

		PlayerServer(response, request)
		got := response.Body.String()
		want := "20"
		assertResponseBody(t, got, want)
	})
	t.Run("Returns Floyd's Score : ", func(t *testing.T) {
		request, _ := newRequest("Floyd")
		response := httptest.NewRecorder()

		PlayerServer(response, request)
		got := response.Body.String()
		want := "10"
		assertResponseBody(t, got, want)
	})
}

func assertResponseBody(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got %q while want %q ", got, want)
	}
}

func newRequest(name string) (*http.Request, error) {
	return http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
}
