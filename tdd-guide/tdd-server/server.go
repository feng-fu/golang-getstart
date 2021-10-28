package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

type Player struct {
	Name string
	Wins int
}

const jsonContentType = "application/json"

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score, ok := s.scores[name]
	if !ok {
		return 0
	}
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store
	router := http.NewServeMux()

	router.Handle("/players/", http.HandlerFunc(p.handlePlayer))
	router.Handle("/league", http.HandlerFunc(p.handleLeague))
	p.Handler = router

	return p
}

func (p *PlayerServer) handlePlayer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) handleLeague(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(p.store.GetLeague())
	w.Header().Set("content-type", jsonContentType)
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	score, ok := i.store[name]
	if !ok {
		return 0
	}
	return score
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	// fmt.Println(name)
	_, ok := i.store[name]
	if !ok {
		i.store[name] = 0
	}
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}

// func main() {
// 	server := &PlayerServer{&InMemoryPlayerStore{}}

// 	if err := http.ListenAndServe(":5000", server); err != nil {
// 		log.Fatalf("could not listen on port 5000 %v", err)
// 	}
// }
