package main

import (
	"log"
	"net/http"
	"rediscache/internal/db"
	"rediscache/internal/handlers"
	"rediscache/internal/storage"
)

const (
	host = "localhost:8080"
)

func main() {
	storage.Conn()
	db.Conn()
	log.Printf("Server up at host: %v", host)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{id}", handlers.GetHandler)
	mux.HandleFunc("GET /", handlers.GetAllHandler)
	mux.HandleFunc("POST /", handlers.PostHandler)
	mux.HandleFunc("PUT /{id}", handlers.PutHandler)
	mux.HandleFunc("DELETE /{id}", handlers.DeleteHandler)

	if err := http.ListenAndServe(host, mux); err != nil {
		log.Fatal(err)
	}
}
