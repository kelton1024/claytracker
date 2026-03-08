package main

import (
	"log"
	"net/http"
)

func loggerMiddleware(endpoint http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request from client '%v'\n", r.RemoteAddr)
		endpoint.ServeHTTP(w, r)
	})
}
