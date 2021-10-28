package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("/league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store := FileSystemStore{database}
		got := store.GetLeague()

		want := League{
			{"Cleo", 10},
			{"Chris", 33},
		}
		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)

		defer cleanDatabase()

		store := FileSystemStore{database}
		got := store.GetPlayerScore("Chris")

		assertScoreEqual(t, got, 33)
	})

	t.Run("store wins for existing players", func(t *testing.T) {

		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store := FileSystemStore{database}
		store.RecordWin("Chris")
		got := store.GetPlayerScore("Chris")
		assertScoreEqual(t, got, 34)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()
		store := FileSystemStore{database}

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		assertScoreEqual(t, got, 1)
	})
}

func assertScoreEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("want '%d', got '%d'", want, got)
	}
}

func createTempFile(t *testing.T, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
