package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	log.Println("Starting Server at : 5000")
	log.Fatal(http.ListenAndServe(":5000", server))
}
