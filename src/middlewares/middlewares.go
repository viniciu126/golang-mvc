package middlewares

import (
	"api/src/auth"
	"api/src/responses"
	"log"
	"net/http"
)

// Logger log request info
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf(" %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Auth verify if user is autenticated
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			responses.Err(w, http.StatusUnauthorized, err)
			return
		}

		next(w, r)
	}
}
