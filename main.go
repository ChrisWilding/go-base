package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", http.HandlerFunc(hello)).Methods(http.MethodGet)
	http.ListenAndServe(":8080", r)
}
