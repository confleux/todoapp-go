package auth

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"

	"client-service/internal/repository"
)

type AuthService struct {
	app    *firebase.App
	userDB *repository.UserRepository
}

func NewAuthService(serviceAccountKeyPath string, userDB *repository.UserRepository) (*AuthService, error) {
	opt := option.WithCredentialsFile(serviceAccountKeyPath)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase app: %w", err)
	}

	return &AuthService{app: app, userDB: userDB}, nil
}

func (s *AuthService) SignUp(ctx context.Context, email string, password string) (*auth.UserRecord, error) {
	client, err := s.app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("error loading firebase client: %w", err)
	}

	params := (&auth.UserToCreate{}).Email(email).EmailVerified(false).Password(password)

	createdUser, err := client.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *AuthService) VerifyToken(ctx context.Context, idToken string) (*auth.Token, error) {
	client, err := s.app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("error loading firebase client: %w", err)
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, fmt.Errorf("unable to authorize token: %w", err)
	}

	return token, nil
}

func (s *AuthService) CreateUser(ctx context.Context, uid string, email string) error {
	_, err := s.userDB.CreateUser(ctx, uid, email)
	if err != nil {
		return err
	}

	return nil
}
