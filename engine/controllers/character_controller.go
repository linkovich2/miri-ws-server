package controllers

import (
	"github.com/jonathonharrell/miri-ws-server/engine/api/parameters"
	"github.com/jonathonharrell/miri-ws-server/engine/services"
	"github.com/jonathonharrell/miri-ws-server/engine/logger"
	"github.com/jonathonharrell/miri-ws-server/engine/core/authentication"
  jwt "github.com/dgrijalva/jwt-go"
	"encoding/json"
	"net/http"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

func CreateCharacter(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	character := new(parameters.Character)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&character)

	userId, err := getUserId(r)
	if err != nil {
		logger.Write.Error(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	responseStatus, responseBody := services.CreateCharacter(character, userId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(responseBody)
}

func ListCharacters(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	userId, err := getUserId(r)
	if err != nil {
		logger.Write.Error(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	responseStatus, list := services.ListCharacters(userId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(list)
}

func getUserId(r *http.Request) (string, error) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		return authBackend.Key, nil
	})
	if err != nil {
		return "", err
	}

	if !bson.IsObjectIdHex(token.Claims["sub"].(string)) {
		return "", errors.New("Invalid hex value for user ID found. Hacking attempt?")
	}

	return token.Claims["sub"].(string), nil
}
