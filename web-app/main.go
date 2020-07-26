package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}
func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	port := "5000"
	log.Println("Starting server at :", port)
	http.ListenAndServe(":"+port, server)
}
