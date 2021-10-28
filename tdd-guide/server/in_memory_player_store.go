package main

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

func (i *InMemoryPlayerStore) GetLeague() League {
	var league League
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}
