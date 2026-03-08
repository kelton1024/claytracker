// Ref: https://go.dev/wiki/cgo

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
	"log"
	"net/http"
)

// TODO: Read environment variables/config file
const (
	address = ":8080"
	dbPath  = "/tmp/rocksdb_database"
)

func init() {
	db := C.CString(dbPath)
	cValue := C.init(db)
	if int(cValue) != 0 {
		panic("failed to initialize database")
	}
}

func loggerMiddleware(endpoint http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request from client '%v'\n", r.RemoteAddr)
		endpoint.ServeHTTP(w, r)
	})
}

func main() {
	log.Println("Starting API...")
	mux := http.NewServeMux()
	// TODO: Add update/delete endpoints
	mux.HandleFunc("/query", loggerMiddleware(http.HandlerFunc(queryEndpoint)))
	mux.HandleFunc("/insert", loggerMiddleware(http.HandlerFunc(insertEndpoint)))
	http.ListenAndServe(address, mux)
}
