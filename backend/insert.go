package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5"
)

func insert(station int, scores string) error {
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
	tx.Commit(ctx)
	return nil
}

func insertEndpoint(w http.ResponseWriter, r *http.Request) {
	requestBody := struct {
		Station int      `json:"course"`
		Scores  []string `json:"score"`
	}{}
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&requestBody)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"message": "failed to decode request body"}`))
		return
	}

	// TODO: Add validation before calling insert
	// TODO: make this less terrible lol
	var sb strings.Builder
	for _, score := range requestBody.Scores {
		sb.Write([]byte(score))
	}

	fmt.Println("sending request")
	err = insert(requestBody.Station, sb.String())
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Success!\n"))
}
