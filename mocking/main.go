package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int;	
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type ConfigurableSleeper struct	{
	duration time.Duration
}

func (c *ConfigurableSleeper) Sleep() {
	time.Sleep(c.duration)
}

const finalWord = "Go!"
const countdownStart = 3

type CountdownOperationSpy struct {
	Calls []string
}

const write = "write"
const sleep = "sleep"

func (c *CountdownOperationSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	configurableSleeper := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, configurableSleeper)
}