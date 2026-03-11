package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func insertEndpoint(w http.ResponseWriter, r *http.Request) {
	requestBody := struct {
		Station int `json:"course"`
		Scores  int `json:"score"`
	}{}
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&requestBody)
	if err != nil {
		errMsg := fmt.Sprintf(`{"message": "%v"}`, err.Error())
		w.Write([]byte(errMsg))
		return
	}

	// TODO: Add validation before calling insert

	fmt.Println("sending request")
	err = insert(requestBody.Station, requestBody.Scores)
	if err != nil {
		errMsg := fmt.Sprintf(`{"message": "%v"}`, err.Error())
		w.Write([]byte(errMsg))
		return
	}
	w.Write([]byte(`{"message": "success"}`))
}

func queryEndpoint(w http.ResponseWriter, r *http.Request) {
	requestBody := struct {
		Key string `json:"key"`
	}{}
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&requestBody)
	if err != nil {
		w.Write([]byte("failed to decode request body"))
		return
	}

	result, err := query(requestBody.Key)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(result + "\n"))

	err = r.Body.Close()
	if err != nil {
		log.Println("failed to close request body")
	}
}
