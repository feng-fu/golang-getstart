package main

import (
	"encoding/json"
	"io"
)

type FileSystemStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemStore) GetLeague() League {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemStore) RecordWin(name string) {
	league := f.GetLeague()
	winner := league.Find(name)
	if winner != nil {
		winner.Wins++
	} else {
		league = append(league, Player{name, 1})
	}
	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}

func (f *FileSystemStore) GetPlayerScore(name string) (wins int) {
	league := f.GetLeague()
	player := league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}
