package poker_test

import (
	"testing"
	"strings"
	"github.com/eason/application"
)

func TestCLI(t *testing.T) {
	t.Run("record Eason win from user input", func(t *testing.T) {
		in := strings.NewReader("Eason wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Eason")
	})

	t.Run("record Cleo  win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo  wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo ")
	})
}