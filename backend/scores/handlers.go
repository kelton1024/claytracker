package scores

import (
	"backend/middleware"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type ScoreHandler struct {
	service *ScoreService
}

func NewScoreHandler(mux *http.ServeMux, db *pgx.Conn) error {
	scoreSvc, err := NewScoreService(db)
	if err != nil {
		return err
	}
	scoreHandler := &ScoreHandler{service: scoreSvc}
	scoreHandler.registerScoreEndpoints(mux)
	return nil
}

func (sh *ScoreHandler) registerScoreEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("/add_score", middleware.Auth(middleware.Logger(http.HandlerFunc(sh.addScoreEndpoint))))
	mux.HandleFunc("/get_score", middleware.Auth(middleware.Logger(http.HandlerFunc(sh.queryEndpoint))))
}

// Handle HTTP related task here
func (sh *ScoreHandler) addScoreEndpoint(w http.ResponseWriter, r *http.Request) {
	var scoreData scoreAddRequest
	err := json.NewDecoder(r.Body).Decode(&scoreData)
	if err != nil {
		errMsg := fmt.Sprintf(`{"Decode error": "%v"}`, err.Error())
		w.Write([]byte(errMsg))
		return
	}

	ctx := context.Background()
	err = sh.service.AddScore(ctx, &scoreData)
	if err != nil {
		errMsg := fmt.Sprintf(`{"error":"failed to add score %v"}`, err.Error())
		w.Write([]byte(errMsg))
		return
	}

	w.Write([]byte(`{"message":"sucess"}`))

	err = r.Body.Close()
	if err != nil {
		log.Println("failed to close request body")
	}
}

// TODO: This needs updated to get the score from a query parameter instead of the request body
func (sh *ScoreHandler) queryEndpoint(w http.ResponseWriter, r *http.Request) {
	var scoreKey ScoreGetRequest
	err := json.NewDecoder(r.Body).Decode(&scoreKey)
	if err != nil {
		w.Write([]byte("failed to decode request body"))
		return
	}

	key := scoreKey.Key
	ctx := context.Background()
	response, err := sh.service.QueryScore(ctx, key)
	if err != nil {
		errMsg := fmt.Sprintf(`{"error":"failed to get score '%v' with error %v"}`, key, err)
		w.Write([]byte(errMsg))
		return
	}

	// TODO: Update response format
	w.Write(response)

	err = r.Body.Close()
	if err != nil {
		log.Println("failed to close request body")
	}
}
