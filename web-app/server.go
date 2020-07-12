package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(player string) int
}

type PlayerServer struct {
	store PlayerStore
}
type StubPlayerStore struct {
	scores map[string]int
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, p.store.GetPlayerScore(player))

}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	score := s.scores[player]
	return score
}
