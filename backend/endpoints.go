package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/alexedwards/argon2id"
	"runtime"
)

var encParams = &argon2id.Params {
	Memory:      64 * 1024,
	Iterations:  1,
	Parallelism: uint8(runtime.NumCPU()),
	SaltLength:  16,
	KeyLength:   32,
}

type EndpointData interface {
	Validate() (EndpointData, error)

	HandleDatabase() error
}

//   -----  RANGE  -----
type RangeData struct {
		Name string `json:"name"`
		Address1 string `json:"address1"`
		Address2 string `json:"address2"`
		City string `json:"city"`
		State string `json:"state"`
		Zipcode string `json:"zipcode"`
		Lat float64 `json:"lat"`
		Long float64 `json:"long"`
}

//TODO: validate Range input
func (r RangeData) Validate() (EndpointData, error) {
	return r,nil
}
func (r RangeData) HandleDatabase() error {
	return addRange(
		r.Name,
		r.Address1,
		r.Address2,
		r.City,
		r.State,
		r.Zipcode,
		r.Lat,
		r.Long)
}

func addRangeEndpoint(w http.ResponseWriter, r *http.Request){
	addEndpoint(&RangeData{},w,r)
}
// ----- END RANGE -----

// ----- USER -----
type UserData struct {
	Username string `json:"username"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City string `json:"city"`
	State string `json:"state"`
	Zipcode string `json:"zipcode"`
	Password string `json:"password"`
}
//TODO: validate user input
func (u UserData) Validate() (EndpointData, error){
	return u, nil
}

func (u UserData) HandleDatabase() error{
	passHash,err := argon2id.CreateHash(u.Password, encParams)
	if err != nil {
		return err
	}
	//TODO: this can be removed after verification that its working
	match,err := argon2id.ComparePasswordAndHash(u.Password,passHash)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("Passowrd verificatin error")
	}

	return registerUser(
	u.Username,
	u.FirstName,
	u.LastName,
	u.Email,
	u.Address1,
	u.Address2,
	u.City,
	u.State,
	u.Zipcode,
	passHash)
}
func registerUserEndpoint(w http.ResponseWriter, r *http.Request) {
	addEndpoint(&UserData{},w,r)
}
// ----- END USER -----



//  -----  SCORE  -----
type ScoreData struct {
	// TODO: fix score struct to reflect data model
	Station int `json:"course"`
	Scores  int `json:"score"`
}
//TODO: validate Score input
func (s ScoreData) Validate() (EndpointData, error) {
	return s,nil
}
func (s ScoreData) HandleDatabase() error{
	return addScore(s.Station,s.Scores)
}

func addScoreEndpoint(w http.ResponseWriter, r *http.Request) {
	addEndpoint(&ScoreData{},w,r)
}
// ----- END SCORE -----

func addEndpoint(e EndpointData, w http.ResponseWriter, r *http.Request){
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(e)
	if err != nil {
		errMsg := fmt.Sprintf(`{"Decode error": "%v"}`, err.Error())
		w.Write([]byte(errMsg))
		return
	}
	_,err = e.Validate()
	if err != nil {
		errMsg := fmt.Sprintf(`{"Validation error": "%v"}`, err.Error())
		w.Write([]byte(errMsg))
		return
	}
	fmt.Println("sending request")
	err= e.HandleDatabase()
	if err != nil {
		errMsg := fmt.Sprintf(`{"Database error": "%v"}`, err.Error())
		w.Write([]byte(errMsg))
		return
	}
	w.Write([]byte(`{"message":"sucess"}`))
}


// -----  TODO: NEED SIMILAR METHOD AS ABOVE FOR RETREIVING DATA -----

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
