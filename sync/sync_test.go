package sync

import "testing"

func TestCounter(t *testing.T) {
	t.Run("Running Counter 3 times : ", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		if counter.Value() != 3 {
			t.Errorf("got %d while want %d", 3, counter.Value())
		}
	})
}