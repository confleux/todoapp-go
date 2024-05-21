package auth

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"google.golang.org/api/option"
)

type AuthService struct {
	app *firebase.App
}

func NewAuthService(serviceAccountKeyPath string) (*AuthService, error) {
	opt := option.WithCredentialsFile(serviceAccountKeyPath)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase app: %w", err)
	}

	return &AuthService{app: app}, nil
}

func (as *AuthService) SignUp(ctx context.Context, email string, password string) (*auth.UserRecord, error) {
	client, err := as.app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("error loading firebase client: %w", err)
	}

	params := (&auth.UserToCreate{}).Email(email).EmailVerified(false).Password(password)

	createdUser, err := client.CreateUser(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	return createdUser, nil
}

func (as *AuthService) VerifyToken(ctx context.Context, idToken string) (*auth.Token, error) {
	client, err := as.app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("error loading firebase client: %w", err)
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, fmt.Errorf("unable to authorize token: %w", err)
	}

	return token, nil
}
