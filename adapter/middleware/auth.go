package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/loan-service/internal/env"
)

type BearerToken = []byte

// AppToken is a middleware to check Authorization Bearer Header
// is a valid `env` Api Token
func (m Middleware) AppToken(w http.ResponseWriter, r *http.Request) (*http.Request, *Error) {
	token, err := GetBearerToken(r)

	if err != nil {
		return nil, &Error{err, http.StatusUnauthorized}
	}

	if env.AppToken() != string(token) {
		return nil, &Error{errors.New("Invalid AppToken"), http.StatusUnauthorized}
	}

	return r, nil
}

// InternalToolToken is a middleware to check Authorization Bearer Header
// is a valid `env` Api Token
func (m Middleware) InternalToolToken(w http.ResponseWriter, r *http.Request) (*http.Request, *Error) {
	token, err := GetBearerToken(r)

	if err != nil {
		return nil, &Error{err, http.StatusUnauthorized}
	}

	if env.InternalToolToken() != string(token) {
		return nil, &Error{errors.New("Invalid InternalToolToken"), http.StatusUnauthorized}
	}

	return r, nil
}

// GetBearerToken ...
func GetBearerToken(r *http.Request) (BearerToken, error) {
	authorizationHeader := r.Header.Get("Authorization")
	splitAuthorizationHeader := strings.Split(authorizationHeader, "Bearer")

	if len(splitAuthorizationHeader) != 2 {
		return nil, errors.New("invalid authorization bearer header")
	}

	token := strings.TrimSpace(splitAuthorizationHeader[1])

	return []byte(token), nil
}
