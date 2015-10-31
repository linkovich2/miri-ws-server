package services

import (
  "net/http"

  "github.com/jonathonharrell/miri-ws-server/engine/api/parameters"
  "github.com/jonathonharrell/miri-ws-server/engine/models"
  "github.com/jonathonharrell/miri-ws-server/engine/core/database"

  "gopkg.in/mgo.v2/bson"
)

func CreateCharacter(character *parameters.Character, userId string) (status int, body []byte) {
  // @todo stub
  return http.StatusCreated, []byte("Whatever")
}

func ListCharacters(userId string) (status int, body []byte) {
  u := models.User{}
	err := database.GetDB().C("users").Find(bson.M{"_id": bson.ObjectIdHex(userId)}).One(&u)

	if err != nil { // no existing user
		return http.StatusUnauthorized, nil
	}

  // @todo list characters

  return http.StatusOK, []byte("Whatever")
}
