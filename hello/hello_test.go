package main

import "testing"

func TestHello(t *testing.T) {
     assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got '%q' want '%q'", got, want)
        }
    }

    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris", "Spanish")
        want := "Hola, Chris"
        assertCorrectMessage(t, got, want)
    })

    t.Run("empty string defaults to 'world'", func(t *testing.T) {
        got := Hello("", "French")
        want := "Bonjour, World"
        assertCorrectMessage(t, got, want)
    })
}