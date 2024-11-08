package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thergupta2001/go-backend.git/api/routes"
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

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	routes.SignUpRoute(router)

	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}