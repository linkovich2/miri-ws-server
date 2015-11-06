package controllers

import (
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/engine/api/parameters"
	"github.com/jonathonharrell/miri-ws-server/engine/services"
	"net/http"
)

func CreateCharacter(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	character := new(parameters.Character)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&character)

	userId, err := getUserId(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	responseStatus, responseBody := services.CreateCharacter(character, userId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(responseBody)
}

func CharacterOptions(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	responseStatus, options := services.CharacterOptions(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(options)
}

func ListCharacters(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	userId, err := getUserId(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	responseStatus, list := services.ListCharacters(userId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(list)
}
