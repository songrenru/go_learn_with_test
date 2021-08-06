package poker

import (
	"encoding/json"
	// "io"
	"fmt"
	"os"
)

func NewLeague(file *os.File) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(file).Decode(&league)
	if err != nil {
		fmt.Errorf("problem parsing league, %v", err)
	}
	return league, err
}

type League []Player

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}