package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!")) //nolint:errcheck
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", http.HandlerFunc(hello)).Methods(http.MethodGet)
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}
