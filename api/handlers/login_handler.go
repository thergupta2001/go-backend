package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/thergupta2001/go-backend.git/cmd/api"
	"github.com/thergupta2001/go-backend.git/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		defer r.Body.Close()
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var user interface{}
	var err error

	user = &models.Doctor{}
	err = api.DB.Where("email = ?", req.Email).First(user).Error

	if err == gorm.ErrRecordNotFound {
		user = &models.Receptionist{}
		err = api.DB.Where("email = ?", req.Email).First(user).Error
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error fetching user", http.StatusInternalServerError)
		}
		return
	}

	var hashedPassword string
	var userID uint
	var role string

	switch v := user.(type) {
	case *models.Doctor:
		hashedPassword = v.Password
		userID = v.ID
		role = "doctor"
	case *models.Receptionist:
		hashedPassword = v.Password
		userID = v.ID
		role = "receptionist"
	default:
		http.Error(w, "Invalid user type", http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(api.JWTSecret))
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}