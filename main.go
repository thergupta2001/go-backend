package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/thergupta2001/go-backend.git/cmd/api"
)

func main() {
	api.SetupDB()
	startAPIServer()
}

func startAPIServer() {
	port := "8080"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! on %s", port)
	})

	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}