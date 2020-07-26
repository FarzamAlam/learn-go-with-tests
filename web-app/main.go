package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{}
	port := "5000"
	log.Println("Starting server at :", port)
	http.ListenAndServe(":"+port, server)
}
