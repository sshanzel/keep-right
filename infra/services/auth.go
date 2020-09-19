package services

import (
	"errors"

	"firebase.google.com/go/auth"
)

// VerifyToken checks if the token on the request is valid
func VerifyToken(token string) (*auth.Token, error) {
	ctx := GetContext()
	app, err := GetFirebaseApp()

	if err != nil {
		return nil, errors.New("Unable to initialize Firebase App")
	}

	client, err := app.Auth(ctx)

	if err != nil {
		return nil, errors.New("Error getting Auth Client")
	}

	return client.VerifyIDToken(ctx, token)
}
