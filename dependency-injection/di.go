package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyHandlerFunc))
}

func MyHandlerFunc(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Farzam")
}
