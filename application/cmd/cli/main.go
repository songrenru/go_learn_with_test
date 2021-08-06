package main

import (
	"fmt"
	"github.com/eason/application"
	"os"
	"log"
)

const dbFileName = "game.db.json"

func main() {
    store, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
    if err != nil {
        log.Fatal(err)
    }

	fmt.Println("Let's play poker")
	fmt.Println("Type {name} wins to record a win")

    poker.NewCLI(store, os.Stdin).PlayPoker()
}