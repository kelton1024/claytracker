package main

/*
#cgo CXXFLAGS: -std=c++20 -I../glue
#cgo LDFLAGS: -L..//glue -lglue -static-libstdc++ -lpthread
#include "../glue/glue.h"
*/
import (
	"C"
)
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func query(key string) (string, error) {
	cKey := C.CString(key)

	// Make call to our glue code
	cResults := C.query(cKey)
	results := C.GoString(cResults)

	var emptyString string
	if results == emptyString {
		return emptyString, fmt.Errorf("failed to query database using key '%v'", key)
	}
	return results, nil
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
