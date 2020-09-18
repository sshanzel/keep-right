package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/api/option"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

// ExtractToken gets token embedded on the request
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

// VerifyToken checks if the token on the request is valid
func VerifyToken(r *http.Request) (*auth.Token, error) {
	ctx := GetContext()
	app, err := GetFirebaseApp()

	if err != nil {
		return nil, errors.New("Unable to initialize Firebase App")
	}

	client, err := app.Auth(ctx)

	if err != nil {
		return nil, errors.New("Error getting Auth Client")
	}

	tokenID := ExtractToken(r)

	return client.VerifyIDToken(ctx, tokenID)
}

var _app *firebase.App
var _ctx = context.Background()
var _env = os.Getenv("GAE_ENV")
var _path = "C:/Users/hsolevilla/Downloads/keepright-firebase-adminsdk-nc3ok-7dfcf9e1a4.json"

// GetFirebaseApp returns firebase.App instance
func GetFirebaseApp() (*firebase.App, error) {
	if _app != nil {
		return _app, nil
	}

	var err error
	if _env == "" {
		opt := option.WithCredentialsFile(_path)
		_app, err = firebase.NewApp(context.Background(), nil, opt)
	} else {
		_app, err = firebase.NewApp(context.Background(), nil)
	}

	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	return _app, nil
}

// GetContext returns a Singleton Context
func GetContext() context.Context {
	return _ctx
}
