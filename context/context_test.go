package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}
func TestServer(t *testing.T) {
	t.Run("Cancel the req in 5 milli seconds :", func(t *testing.T) {
		data := "hello world"
		store := &SpyStore{response: data}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingContext, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingContext)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		if !store.cancelled {
			t.Errorf("store was not able to cancel")
		}
	})
	t.Run("match the fetch method : ", func(t *testing.T) {
		data := "hello world"
		server := Server(&SpyStore{response: data})

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		if response.Body.String() != data {
			t.Errorf("got %s while want %s", response.Body.String(), data)
		}
	})

}
