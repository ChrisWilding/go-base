package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Heroku!")) //nolint:errcheck
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", http.HandlerFunc(hello)).Methods(http.MethodGet)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	err := http.ListenAndServe(addr, r)
	log.Fatal(err)
}
