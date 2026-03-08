package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func query(key string) (string, error) {
	return "", nil
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
