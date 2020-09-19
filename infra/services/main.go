package services

import (
	"fmt"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/api/option"

	firebase "firebase.google.com/go"
)

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
