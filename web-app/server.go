package main

import (
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	if player == "Pepper" {
		w.Write([]byte("20"))
		return
	}
	if player == "Floyd" {
		w.Write([]byte("10"))
		return
	}
}
