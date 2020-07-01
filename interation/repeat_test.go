package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("repeated is %q while expected is %q", repeated, expected)
	}
}

func BenchmarkRepeat(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Repeat("a", 100)
	}
}

func ExampleRepeat() {
	res := Repeat("a", 4)
	fmt.Println(res)
	//Output: aaaa
}
