package main

import ( 
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {
	buffer := bytes.Buffer{}
	spySleeper := &SpySleeper{}
	CountDown(&buffer,spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("got %q while want %q", got, want)
	}
	if spySleeper.Calls != 4{
		t.Errorf("got %d while want 4",spySleeper.Calls)
	}
}