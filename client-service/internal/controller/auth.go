package controller

import (
	"client-service/internal/entities"
	"client-service/internal/repository"
	"client-service/internal/service/auth"
	"encoding/json"
	"net/http"
)

type AuthController struct {
	authService *auth.AuthService
	userDB      *repository.UserRepository
}

func NewAuthController(authService *auth.AuthService, userDB *repository.UserRepository) *AuthController {
	return &AuthController{authService: authService, userDB: userDB}
}

func (ac *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	var user entities.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userRecord, err := ac.authService.SignUp(r.Context(), user.Email, user.Password)
	if err != nil {
		http.Error(w, "Unable to sign up user", http.StatusInternalServerError)
		return
	}

	uid, err := ac.userDB.CreateUser(r.Context(), userRecord.UID, userRecord.Email)
	if err != nil {
		http.Error(w, "Unable to sign up user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(uid))
}
