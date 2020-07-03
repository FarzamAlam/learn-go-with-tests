package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countDownStart = 3
const finalWord = "Go!"
const write = "write"
const sleep = "sleep"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct{}

type CountDownOperationsSpy struct {
	Calls []string
}

func (s *CountDownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountDownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

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
