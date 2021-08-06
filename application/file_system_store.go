package poker

import (
	// "io"
	"encoding/json"
	"os"
	"fmt"
	"sort"
)

type FileSystemStore struct {
	// database io.ReadWriteSeeker
	// database io.Writer
	// database tape
	database *json.Encoder
	league League
}

// func NewFileSystemStore(database io.ReadWriteSeeker) *FileSystemStore {
func NewFileSystemStore(file *os.File) (*FileSystemStore, error) {
	err := initialPlayerDBFile(file)

	if err != nil {
        return nil, fmt.Errorf("problem initialising player db file, %v", err)
    }

	league, err := NewLeague(file)

	if err != nil {
        return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
    }

	return &FileSystemStore{
		// database: database,
		// database: tape{database},
		database: json.NewEncoder(&tape{file}),
		league: league,
	}, nil
}

func FileSystemPlayerStoreFromFile(path string) (*FileSystemStore, error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
        return nil, fmt.Errorf("problem opening %s %v", path, err)
    }

    store, err := NewFileSystemStore(db)
    if err != nil {
        return nil, fmt.Errorf("problem creating file system player store, %v ", err)
    }

    return store, nil
}

func initialPlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()
	if err != nil {
        return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
    }

    if info.Size() == 0 {
    	file.Write([]byte("[]"))
    	file.Seek(0, 0)
    }

    return nil
}

func (f *FileSystemStore) GetLeague() League {
	// f.database.Seek(0, 0)
	// league, _ := NewLeague(f.database)
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func (f *FileSystemStore) GetPlayerScore(name string) int {
	// player := f.GetLeague().Find(name)
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemStore) RecordWin(name string) {
	// league := f.GetLeague()
	// player := league.Find(name)
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	// f.database.Seek(0, 0)
	// json.NewEncoder(f.database).Encode(f.league)
	f.database.Encode(f.league)
}
