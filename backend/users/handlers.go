package users

import (
	"backend/middleware"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(mux *http.ServeMux, db *pgx.Conn) error {
	userSvc, err := NewUserService(db)
	if err != nil {
		return err
	}
	userHandler := &UserHandler{service: userSvc}
	userHandler.registerUserEndpoints(mux)
	return nil
}

func (uh *UserHandler) registerUserEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("/register", middleware.Auth(middleware.Logger(http.HandlerFunc(uh.addUserEndpoint))))
	mux.HandleFunc("/login", middleware.Auth(middleware.Logger(http.HandlerFunc(uh.loginEndpoint))))
	mux.HandleFunc("/update_users_name", middleware.Auth(middleware.Logger(http.HandlerFunc(uh.updateUsersNameEndpoint))))
	mux.HandleFunc("/update_username", middleware.Auth(middleware.Logger(http.HandlerFunc(uh.updateUsernameEndpoint))))
	mux.HandleFunc("/update_users_password", middleware.Auth(middleware.Logger(http.HandlerFunc(uh.updateUserPasswordEndpoint))))
	mux.HandleFunc("/update_users_address", middleware.Auth(middleware.Logger(http.HandlerFunc(uh.updateUserAddressEndpoint))))
	mux.HandleFunc("/delete_user", middleware.Auth(middleware.Logger(http.HandlerFunc(uh.deleteUserEndpoint))))
}

// Handle HTTP related task here
func (uh *UserHandler) addUserEndpoint(w http.ResponseWriter, r *http.Request) {
	var userData userAddRequest
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		errMsg := fmt.Sprintf(`{"Decode error": "%v"}`, err.Error())
		w.Write([]byte(errMsg))
		return
	}

	ctx := context.Background()
	err = uh.service.AddUser(ctx, &userData)
	if err != nil {
		errMsg := fmt.Sprintf(`{"error":"failed to add user %v"}`, err.Error())
		w.Write([]byte(errMsg))
		return
	}

	w.Write([]byte(`{"message":"sucess"}`))

	err = r.Body.Close()
	if err != nil {
		log.Println("failed to close request body")
	}
}

//TODO: Everything below here
func (rh *UserHandler) loginEndpoint(w http.ResponseWriter, r *http.Request) {
	
	w.Write([]byte(`{"message":"Implement"}`))
}

func (rh *UserHandler) updateUsersNameEndpoint(w http.ResponseWriter, r *http.Request) {
	
	w.Write([]byte(`{"message":"Implement"}`))
}
	
func (rh *UserHandler) updateUsernameEndpoint(w http.ResponseWriter, r *http.Request) {
	
	w.Write([]byte(`{"message":"Implement"}`))
}

func (rh *UserHandler) updateUserPasswordEndpoint(w http.ResponseWriter, r *http.Request) {
	
	w.Write([]byte(`{"message":"Implement"}`))
}

func (rh *UserHandler) updateUserAddressEndpoint(w http.ResponseWriter, r *http.Request) {
	
	w.Write([]byte(`{"message":"Implement"}`))
}

func (rh *UserHandler) deleteUserEndpoint(w http.ResponseWriter, r *http.Request) {
	
	w.Write([]byte(`{"message":"Implement"}`))
}
