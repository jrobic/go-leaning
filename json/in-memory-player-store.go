package main

import "sync"

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}, sync.RWMutex{}}
}

type InMemoryPlayerStore struct {
	scores map[string]int
	lock   sync.RWMutex
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.lock.Lock()
	defer i.lock.Unlock()

	return i.scores[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.lock.Lock()
	defer i.lock.Unlock()

	i.scores[name]++
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player

	for name, wins := range i.scores {
		league = append(league, Player{name, wins})
	}

	return league
}
