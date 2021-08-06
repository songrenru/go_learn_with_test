package main

import (
	"testing"
	"bytes"
	"reflect"
)

func TestCountdown(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}
		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}

		if spySleeper.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
		}
	})

	t.Run("sleep after every print", func(t *testing.T) {
		spySleeper := &CountdownOperationSpy{}
		Countdown(spySleeper, spySleeper)

		got := spySleeper.Calls
		want := []string{
	        write,
	        sleep,
	        write,
	        sleep,
	        write,
	        sleep,
	        write,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	})
}