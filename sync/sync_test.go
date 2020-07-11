package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Running Counter n times : ", func(t *testing.T) {
		counter := NewCounter()
		want := 1000
		var wg sync.WaitGroup
		wg.Add(want)
		for i := 0; i < want; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()
		assertCounter(t, counter, want)
	})
}

func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d while want %d", got.Value(), want)
	}
}
