package main

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Database struct {
}

func NewDBConnection(connString string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func query(key string) (string, error) {
	return "", nil
}

func addScore(station int, scores int) error {
	ctx := context.Background()
	tx, err := dbConn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	sql := `INSERT INTO scores_tracking (score, station_number) VALUES ($1, $2);`
	_, err = tx.Exec(ctx, sql, scores, station)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func addRange(name string, address1 string, address2 string, city string, state string, zipcode string, lat float64, long float64) error {
	ctx := context.Background()
	tx, err := dbConn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	sql := `INSERT INTO ranges(name, address1, address2, city, state_id, zipcode, lat, lng) VALUES ($1, $2, $3, $4, (select state_id from states where name=UPPER($5)), $6, $7, $8);`
	_, err = tx.Exec(ctx, sql, name, address1, address2, city, state, zipcode, lat, long)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func registerUser(username string, firstname string, lastname string, email string, address1 string, address2 string, city string, state string, zipcode string, passHash string) error{
	ctx := context.Background()
	tx, err := dbConn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	sql := `INSERT INTO users(username, first_name, last_name, email, address1, address2, city, state_id, zipcode, password_hash) VALUES ($1, $2, $3, $4, $5, $6, $7, (select state_id from states where name=UPPER($8)), $9, $10);`
	_, err = tx.Exec(ctx, sql, username, firstname, lastname, email, address1, address2, city, state, zipcode, passHash)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
