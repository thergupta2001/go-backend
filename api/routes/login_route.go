package routes

import (
	"github.com/gorilla/mux"
	"github.com/thergupta2001/go-backend.git/api/handlers"
)

func LoginRoute(router *mux.Router) {
	router.HandleFunc("/login", handlers.Login).Methods("POST")
}