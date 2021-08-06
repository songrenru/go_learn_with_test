package poker

import (
	"testing"
	// "strings"
	// "io"
	"io/ioutil"
	"os"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("/league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"name": "Celo", "wins": 78},
			{"name": "Chris", "wins": 99}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemStore(database)
		assertNoError(t, err)

		got := store.GetLeague()
		want := []Player{
			{"Chris", 99},
			{"Celo", 78},
		}

		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"name": "Celo", "wins": 78},
			{"name": "Chris", "wins": 99}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemStore(database)
		assertNoError(t, err)

		got := store.GetPlayerScore("Chris")
		want := 99

		if got != want {
	        t.Errorf("got %d want %d", got, want)
	    }
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"name": "Celo", "wins": 78},
			{"name": "Chris", "wins": 99}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemStore(database)
		assertNoError(t, err)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 100

		if got != want {
	        t.Errorf("got %d want %d", got, want)
	    }
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"name": "Celo", "wins": 78},
			{"name": "Chris", "wins": 99}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemStore(database)
		assertNoError(t, err)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1

		if got != want {
	        t.Errorf("got %d want %d", got, want)
	    }
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemStore(database)
		assertNoError(t, err)
	})

	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"name": "Celo", "wins": 78},
			{"name": "Chris", "wins": 99}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemStore(database)
		assertNoError(t, err)

		got := store.GetLeague()
		want := []Player{
			{"Chris", 99},
			{"Celo", 78},
		}

		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})
}

func assertScoreEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
        t.Errorf("got %d want %d", got, want)
    }
}

func assertNoError(t *testing.T, err error) {
    t.Helper()
    if err != nil {
        t.Fatalf("didnt expect an error but got one, %v", err)
    }
}

// func createTempFile(t *testing.T, initialData string) (io.ReadWriteSeeker, func()) {
func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpFile.Write([]byte(initialData))

	removeFile := func() {
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}