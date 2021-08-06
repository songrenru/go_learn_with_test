package main

import (
	"net/http"
	"log"
	// "os"
	"github.com/eason/application"
)

// type InMemoryPlayerStore struct{}

// func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
//     return 123
// }

// func (i *InMemoryPlayerStore) RecordWin(name string) {}

const dbFileName = "game.db.json"

func main() {
	// db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// 	log.Fatalf("problem opening %s %v", dbFileName, err)
	// }
	// // store := &FileSystemStore{db}
	// store, err := poker.NewFileSystemStore(db)
	// if err != nil {
	//     log.Fatalf("problem creating file system player store, %v ", err)
	// }
    store, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
    if err != nil {
        log.Fatal(err)
    }

	// handler := http.HandlerFunc(PlayerServer)
	// server := &PlayerServer{&InMemoryPlayerStore{}}
	// server := &PlayerServer{NewInMemoryPlayerStore()}
	// server := NewPlayerServer(NewInMemoryPlayerStore())
	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}