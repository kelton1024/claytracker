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
	"net/http"
)

func insert(key string, value string) error {
	cKey := C.CString(key)
	cValue := C.CString(value)

	// Make call to our glue code
	success := C.insert(cKey, cValue)
	if int(success) != 0 {
		return fmt.Errorf("failed to insert key '%v' and value '%v' into the database", key, value)
	}
	return nil
}

func insertEndpoint(w http.ResponseWriter, r *http.Request) {
	requestBody := struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}{}
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&requestBody)
	if err != nil {
		w.Write([]byte("failed to decode request body"))
		return
	}

	// TODO: Add validation before calling insert
	err = insert(requestBody.Key, requestBody.Value)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Success!\n"))
}
