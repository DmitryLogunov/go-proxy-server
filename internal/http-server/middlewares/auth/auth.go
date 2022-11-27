package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/DmitryLogunov/go-proxy-server/internal/configuration"
	jwt "github.com/dgrijalva/jwt-go"
)

// Exception ...
type Exception struct {
	Message string `json:"message"`
}

// ValidateJWT checks if JWT token exists and is correct
func ValidateJWT(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")

		if authorizationHeader == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
			return
		}

		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) != 2 {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization header"})
		}

		token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte("secret"), nil
		})

		if error != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: error.Error()})
			return
		}

		if token.Valid {
			h.ServeHTTP(w, req)
		} else {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
		}
	})
}

// ValidateToken checks if token (as a secret string from environment variable) exists and is correct
func ValidateToken(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		config := *(&configuration.Configuration{}).GetInstance()

		authorizationHeader := req.Header.Get("authorization")

		if authorizationHeader == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
			return
		}

		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) != 2 {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization header"})
		}

		if strings.Compare(bearerToken[1], config.AuthToken) == 0 {
			h.ServeHTTP(w, req)
		} else {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
		}
	})
}

// NoneAuthenticate handles HTTP request without authentication
func NoneAuthenticate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		h.ServeHTTP(w, req)
	})
}
