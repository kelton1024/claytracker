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

func insert(station int, scores int) error {
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
