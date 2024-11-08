package handlers

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"strings"

	"github.com/thergupta2001/go-backend.git/cmd/api"
	"github.com/thergupta2001/go-backend.git/models"
)

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var req SignUpRequest

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		defer r.Body.Close()
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.Password == "" || req.Role == "" || req.Email == "" {
		http.Error(w, "All fields are required!", http.StatusBadRequest)
		return
	}

	req.Role = strings.ToLower(req.Role)
	if !models.ValidRoles[req.Role] {
		http.Error(w, "Invalid role", http.StatusBadRequest)
		return
	}

	// Signup
	var err error
	switch req.Role {
	case models.DoctorRole:
		err = api.DB.Create(&models.Doctor{Name: req.Name, Email: req.Email, Password: req.Password, Role: req.Role}).Error
	case models.ReceptionistRole:
		err = api.DB.Create(&models.Receptionist{Name: req.Name, Email: req.Email, Password: req.Password, Role: req.Role}).Error
	}

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			http.Error(w, "User already exists", http.StatusConflict)
		} else {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}
