package controllers

import (
	"github.com/jonathonharrell/miri-ws-server/engine/api/parameters"
	"github.com/jonathonharrell/miri-ws-server/engine/services"

	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestUser := new(parameters.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	responseStatus, token := services.Login(requestUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	requestUser := new(parameters.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	responseStatus, token := services.CreateUser(requestUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)
}

func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUser := new(parameters.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	w.Header().Set("Content-Type", "application/json")
	w.Write(services.RefreshToken(requestUser))
}

func Logout(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	err := services.Logout(r)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	requestUser := new(parameters.ForgotPassword)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	status, response := services.ForgotPassword(requestUser.Email)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	requestUser := new(parameters.ResetPassword)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	status, response := services.ResetPassword(requestUser.Password)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
