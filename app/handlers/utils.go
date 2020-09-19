package handlers

import (
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/sshanzel/keep-right/infra/services"
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

// GetFirebaseToken return FirebaseToken
func GetFirebaseToken(r *http.Request) (*auth.Token, error) {
	token := ExtractToken(r)

	return services.VerifyToken(token)
}
