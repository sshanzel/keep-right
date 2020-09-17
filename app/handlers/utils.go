package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/net/context"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
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
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

var _app *firebase.App
var _ctx = context.Background()
var _path = "C:/Users/hsolevilla/Downloads/jovial-trilogy-268111-firebase-adminsdk-foile-7b0c883342.json"

// GetFirebaseApp returns firebase.App instance
func GetFirebaseApp() (*firebase.App, error) {
	if _app != nil {
		return _app, nil
	}

	opt := option.WithCredentialsFile(_path)
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	_app = app

	return app, nil
}

// GetContext returns a Singleton Context
func GetContext() context.Context {
	return _ctx
}
