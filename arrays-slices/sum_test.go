package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertion := func(got, want int, t *testing.T) {
		t.Helper()
		if got != want {
			t.Errorf("got %d while want %d", got, want)
		}
	}

	t.Run("variable size : ", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		assertion(got, want, t)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v while want %v", got, want)
	}
}
