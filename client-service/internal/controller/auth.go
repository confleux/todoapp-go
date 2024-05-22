package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"client-service/internal/entities"
	"client-service/internal/service/auth"
	firebaseAuth "firebase.google.com/go/v4/auth"
)

type AuthController struct {
	authService *auth.AuthService
}

type SignUpResponse struct {
	Uid   string `json:"uid"`
	Email string `json:"email"`
}

func NewAuthController(authService *auth.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// SignUp godoc
//
//	@Summary Sign up user
//	@Description	Sign up user
//	@Tags         signup
//	@Accept       json
//	@Produce      json
//	@Param user body entities.User true "User credentials"
//	@Success		201	{object}	SignUpResponse
//	@Failure		400	{object}	nil "Bad request"
//	@Failure		500	{object}	nil "Internal Server Error"
//	@Router /signup [post]
func (ct *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	var user entities.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userRecord, err := ct.authService.SignUp(r.Context(), user.Email, user.Password)
	if err != nil {
		if firebaseAuth.IsEmailAlreadyExists(err) {
			http.Error(w, "Email already exists", http.StatusBadRequest)
			return
		} else {
			http.Error(w, fmt.Sprintf("Unable to sign up user: %v", err), http.StatusBadRequest)
			return
		}
	}

	err = ct.authService.CreateUser(r.Context(), userRecord.UID, userRecord.Email)
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
