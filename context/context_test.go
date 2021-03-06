package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return -1, nil
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	t.Run("Cancel the req in 5 milli seconds :", func(t *testing.T) {
		t.Helper()
		data := "hello world"
		store := &SpyStore{response: data, t: t}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingContext, cancel := context.WithCancel(request.Context())

		time.AfterFunc(10*time.Millisecond, cancel)

		request = request.WithContext(cancellingContext)
		response := &SpyResponseWriter{}

		server.ServeHTTP(response, request)

		if response.written {
			t.Errorf("a response has not been written.")
		}

	})
	t.Run("match the fetch method : ", func(t *testing.T) {
		data := "hello world"
		store := &SpyStore{response: data, t: t}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %s while want %s", response.Body.String(), data)
		}

	})

}
