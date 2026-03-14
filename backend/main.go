package main

import (
	"backend/ranges"
	"backend/scores"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func main() {
	config, err := loadConfig("config.yaml")
	fmt.Println(config)

	log.Printf("Connecting to datatbase at %s:%s as %s", config.Host, config.DBPort, config.User)
	rootDatabaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.User, config.Password, config.Host, config.DBPort, config.Name)
	conn, err := pgx.Connect(context.Background(), rootDatabaseURL)
	if err != nil {
		log.Fatalf("failed to create DB connection %v", err)
	}

	// Register range related endpoints
	mux := http.NewServeMux()
	err = ranges.NewRangeHandler(mux, conn)
	if err != nil {
		log.Fatalf("failed to start the range service %v", err)
	}

	// Register score related endpoints
	err = scores.NewScoreHandler(mux, conn)
	if err != nil {
		log.Fatalf("failed to start the score service %v", err)
	}

	// TODO: Register user related endpoints

	log.Printf("Starting API on port %v", config.AppPort)
	http.ListenAndServe(fmt.Sprintf(":%s", config.AppPort), mux)
}
