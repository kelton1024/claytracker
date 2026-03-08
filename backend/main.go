package main

import (
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

// TODO: read these in from conf/env var using viper
const (
	address         = ":8080"
	db_name         = "range_tracker"
	rootDatabaseURL = "postgres://postgres:mysecretpassword@localhost:5432"
)

// TODO: Create a database struct to hold the connections and that can be used by
// the various endpoints
var dbConn *pgx.Conn

// TODO: Add update/delete endpoints and define them
func registerEndpoints() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/query", loggerMiddleware(http.HandlerFunc(queryEndpoint)))
	mux.HandleFunc("/insert", loggerMiddleware(http.HandlerFunc(insertEndpoint)))
	mux.HandleFunc("/login", loggerMiddleware(http.HandlerFunc(queryEndpoint)))
	mux.HandleFunc("/regsiter", loggerMiddleware(http.HandlerFunc(queryEndpoint)))
	return mux
}

func main() {
	log.Printf("Starting API on port %v", address)

	var err error
	dbConn, err = NewDBConnection(rootDatabaseURL)
	if err != nil {
		log.Fatalf("failed to create DB connection %v", err)
	}

	mux := registerEndpoints()
	http.ListenAndServe(address, mux)
}
