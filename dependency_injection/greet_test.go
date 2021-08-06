package greet

import (
	"testing"
	"bytes"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Eason")

	got := buffer.String()
	want := "Hello, Eason"

	if got != want {
		t.Errorf("got '%s' but want '%s'", got, want)	
	}
}