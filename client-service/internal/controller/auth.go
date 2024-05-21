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

type SignUpResponse struct {
	Uid   string `json:"uid"`
	Email string `json:"email"`
}

func NewAuthController(authService *auth.AuthService, userDB *repository.UserRepository) *AuthController {
	return &AuthController{authService: authService, userDB: userDB}
}

func (ct *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	var user entities.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userRecord, err := ct.authService.SignUp(r.Context(), user.Email, user.Password)
	if err != nil {
		http.Error(w, "Unable to sign up user", http.StatusInternalServerError)
		return
	}

	_, err = ct.userDB.CreateUser(r.Context(), userRecord.UID, userRecord.Email)
	if err != nil {
		http.Error(w, "Unable to sign up user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err = json.NewEncoder(w).Encode(&SignUpResponse{Uid: userRecord.UID, Email: userRecord.Email}); err != nil {
		http.Error(w, "Unable to sign up user", http.StatusInternalServerError)
		return
	}
}
