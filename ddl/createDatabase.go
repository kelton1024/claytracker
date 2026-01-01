package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func createDatabase(ddlSlice []string) error {
	conn, err := pgx.Connect(context.Background(), rootDatabaseURL)
	if err != nil {
		return err
	}

	// Drop database
	cmdTag, err := conn.Exec(context.Background(), ddlSlice[0])
	if err != nil {
		return err
	}
	fmt.Println(cmdTag)

	// Create database
	cmdTag, err = conn.Exec(context.Background(), ddlSlice[1])
	if err != nil {
		return err
	}
	fmt.Println(cmdTag)

	rangeDatabaseURL := fmt.Sprintf("postgres://postgres:mysecretpassword@localhost:5432/%v", databaseName)
	conn, err = pgx.Connect(context.Background(), rangeDatabaseURL)
	if err != nil {
		return err
	}

	// Create tables
	maxRetry := len(ddlSlice) * 2
	for i := 2; i < len(ddlSlice); i++ {
		cmdTag, err := conn.Exec(context.Background(), ddlSlice[i])
		if err != nil {
			// Retry logic in case a table that isn't created yet is referenced
			if i < maxRetry {
				ddlSlice = append(ddlSlice, ddlSlice[i])
				continue
			}
			return err
		}
		fmt.Println(cmdTag)
	}

	err = conn.Close(context.Background())
	if err != nil {
		return err
	}

	return nil
}
