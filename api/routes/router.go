package routes

import (
	"github.com/gorilla/mux"
	"github.com/thergupta2001/go-backend.git/api/handlers"
)

func SignUpRoute(router *mux.Router) {
	router.HandleFunc("/signup", handlers.SignUp).Methods("POST")
}