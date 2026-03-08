package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func insert(station int, scores string) error {
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
	fmt.Println(requestBody)

	// TODO: Add validation before calling insert
	// err = insert(requestBody.Key, requestBody.Value)
	// if err != nil {
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }
	w.Write([]byte("Success!\n"))
}
