package middleware

import (
	"log"
	"net/http"
)

// TODO: Add function that will add all of the middleware to routes

func Logger(endpoint http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request from client '%v'\n", r.RemoteAddr)
		endpoint.ServeHTTP(w, r)
	})
}

// TODO: Implement auth middleware
func Auth(endpoint http.HandlerFunc) http.HandlerFunc {
	return endpoint
}
