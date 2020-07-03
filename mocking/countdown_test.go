package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountDown(t *testing.T) {
	t.Run("prints 3 to Go!", func (t *testing.T){

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
		
	})
	t.Run("sleep before every print", func(t *testing.T){
		spySleepPrinter := &CountDownOperationsSpy{}
		CountDown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, spySleepPrinter.Calls){
			t.Errorf("wanted calls %v got %v",want, spySleepPrinter.Calls)
		}
	})
}