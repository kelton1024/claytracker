package ranges

import (
	"backend/middleware"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type RangeHandler struct {
	service *RangeService
}

func NewRangeHandler(mux *http.ServeMux, db *pgx.Conn) error {
	rangeSvc, err := NewRangeService(db)
	if err != nil {
		return err
	}
	rangeHandler := &RangeHandler{service: rangeSvc}
	rangeHandler.registerRangeEndpoints(mux)
	return nil
}

func (rh *RangeHandler) registerRangeEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("/add_range", middleware.Auth(middleware.Logger(http.HandlerFunc(rh.addRangeEndpoint))))
	mux.HandleFunc("/update_range", middleware.Auth(middleware.Logger(http.HandlerFunc(rh.updateRangeEndpoint))))
	mux.HandleFunc("/delete_range", middleware.Auth(middleware.Logger(http.HandlerFunc(rh.deleteRangeEndpoint))))
}

// Handle HTTP related task here
func (rh *RangeHandler) addRangeEndpoint(w http.ResponseWriter, r *http.Request) {
	var rangeData rangeAddRequest
	err := json.NewDecoder(r.Body).Decode(&rangeData)
	if err != nil {
		errMsg := fmt.Sprintf(`{"Decode error": "%v"}`, err.Error())
		w.Write([]byte(errMsg))
		return
	}

	ctx := context.Background()
	err = rh.service.AddRange(ctx, &rangeData)
	if err != nil {
		errMsg := fmt.Sprintf(`{"error":"failed to add range %v"}`, err.Error())
		w.Write([]byte(errMsg))
		return
	}

	w.Write([]byte(`{"message":"sucess"}`))

	err = r.Body.Close()
	if err != nil {
		log.Println("failed to close request body")
	}
}

func (rh *RangeHandler) updateRangeEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message":"Implement"}`))
}

func (rh *RangeHandler) deleteRangeEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message":"Implement"}`))
}
