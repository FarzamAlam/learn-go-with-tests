package reflection

import "testing"

func TestWalk(t *testing.T) {
	t.Run("Reflection Suite : ", func(t *testing.T) {
		expected := "Chris"
		var got []string
		x := struct {
			Name string
		}{expected}
		walk(x, func(input string) {
			got = append(got, input)
		})
		if len(got) != 1 {
			t.Errorf("Wrong number of function calls, got %d while want %d", len(got), 1)
		}
		if got[0] != expected {
			t.Errorf("got %q, want %q", got[0], expected)
		}
	})

}
