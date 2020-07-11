package sync

import "testing"

func TestCounter(t *testing.T) {
	t.Run("Running Counter 3 times : ", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounter(t, counter, 3)
	})
}

func assertCounter(t *testing.T, got Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d while want %d", got.Value(), want)
	}
}
