package middlewares

import (
	"api/src/authentication"
	"api/src/response"
	"log"
	"net/http"
)

// Authenticate do authentication
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if error := authentication.ValidateToken(r); error != nil {
			response.Error(w, http.StatusUnauthorized, error)
			return
		}
		next(w, r)
	}
}

// Logger write request information at the terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}
