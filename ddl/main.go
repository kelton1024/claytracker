package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

const (
	outputFileName   = "gen_range_tracker_relational_model.sql"
	jsonDirectory    = "../json/"
	databaseName     = "range_tracker"
	rootDatabaseURL  = "postgres://postgres:mysecretpassword@localhost:5432"
	createDBFile     = "createDatabase.sql"
	createTablesFile = "createTables.sql"
)

func main() {
	conn, err := pgx.Connect(context.Background(), rootDatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to PG %v", err)
	}

	// Drop database
	dropDDL := fmt.Sprintf(`DROP DATABASE IF EXISTS %v`, databaseName)
	cmdTag, err := conn.Exec(context.Background(), dropDDL)
	if err != nil {
		log.Fatalf("failed to connect to create database %v", err)
	}
	fmt.Println(cmdTag)

	// Create database
	createDDL := fmt.Sprintf(`CREATE DATABASE %v;`, databaseName)
	cmdTag, err = conn.Exec(context.Background(), createDDL)
	if err != nil {
		log.Fatalf("failed to connect to create database %v", err)
	}
	fmt.Println(cmdTag)

	// Reconnect to the range database
	rangeDatabaseURL := fmt.Sprintf("postgres://postgres:mysecretpassword@localhost:5432/%v", databaseName)
	conn, err = pgx.Connect(context.Background(), rangeDatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database '%v' %v", databaseName, err)
	}

	// Create all of the tables
	createTableSql, err := os.ReadFile(createTablesFile)
	cmdTag, err = conn.Exec(context.Background(), string(createTableSql))
	if err != nil {
		log.Fatalf("failed to connect to create tables %v", err)
	}
	fmt.Println(cmdTag)

	err = conn.Close(context.Background())
	if err != nil {
		log.Fatalf("failed to close database connection %v", err)
	}
}
