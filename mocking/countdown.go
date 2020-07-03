package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countDownStart = 3
const finalWord = "Go!"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct{}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func CountDown(w io.Writer, s Sleeper) {
	for i := countDownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(w, i)
	}
	s.Sleep()
	fmt.Fprint(w, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	CountDown(os.Stdout, sleeper)
}
