package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		PORT string
	}
	Database struct {
		HOST     string
		PORT     string
		NAME     string
		USER     string
		PASSWORD string
	}
}

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
	// Set up viper to read config file
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// Read config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Get values that match config struct from file
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	log.Printf("Connecting to datatbase at %s:%s as %s", config.Database.HOST, config.Database.PORT, config.Database.USER)
	rootDatabaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.Database.USER, config.Database.PASSWORD, config.Database.HOST, config.Database.PORT, config.Database.NAME)

	dbConn, err = NewDBConnection(rootDatabaseURL)
	if err != nil {
		log.Fatalf("failed to create DB connection %v", err)
	}

	log.Printf("Starting API on port %v", config.App.PORT)
	mux := registerEndpoints()
	http.ListenAndServe(fmt.Sprintf(":%s", config.App.PORT), mux)
}
