package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	port := "5000"
	log.Println("Starting server at :", port)
	http.ListenAndServe(":"+port, handler)
}
