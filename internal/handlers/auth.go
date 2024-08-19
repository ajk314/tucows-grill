package handlers

import (
	"encoding/json"
	"net/http"
	"tucows-grill-service/internal/auth"
	"tucows-grill-service/internal/models"

	"golang.org/x/crypto/bcrypt"
)

// TODO: use real encrypted pw stored in aws or k8s or db, and decrypt
const rawPassword string = "root"
var storedHashedPassword string

func init() {
	// TODO: save this value in db for the user account when user makes an account
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	storedHashedPassword = string(hashedPassword)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(creds.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateJWT(creds.Username)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
