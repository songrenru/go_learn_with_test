package poker

import "testing"

type StubPlayerStore struct {
	scores map[string]int
	winCalls []string
	league []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) < 1 {
        t.Fatal("expected a win call but didn't get any")
    }
    if store.winCalls[0] != winner {
        t.Errorf("didn't record correct winner, got '%s', want '%s'", store.winCalls[0], winner)
    }
}