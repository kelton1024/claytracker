package main

import (
	"log"
	"os"
)

const (
	outputFileName  = "gen_range_tracker_relational_model.sql"
	jsonDirectory   = "../json/"
	databaseName    = "range_tracker"
	rootDatabaseURL = "postgres://postgres:mysecretpassword@localhost:5432"
)

func main() {
	args := os.Args[1:]
	action := args[0]

	switch action {
	case "generate":
		outputFile, err := os.Create(outputFileName)
		if err != nil {
			log.Fatalf("failed to create output file with the following error %v", err)
		}

		ddlSlice, err := generateDDL()
		if err != nil {
			log.Fatalf("failed to create DDL with the following error %v", err)
		}

		for _, sql := range ddlSlice {
			outputFile.Write([]byte(sql))
		}

	case "create":
		ddlSlice, err := generateDDL()
		if err != nil {
			log.Fatalf("failed to create DDL with the following error %v", err)
		}

		err = createDatabase(ddlSlice)
		if err != nil {
			log.Fatalf("failed to create database with the following error %v", err)
		}
	default:
		log.Fatalf("invalid option was provided")
	}
}
