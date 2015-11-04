package services

import (
	"encoding/json"
	"net/http"

	"github.com/jonathonharrell/miri-ws-server/engine/api/parameters"
	"github.com/jonathonharrell/miri-ws-server/engine/core/content"
	"github.com/jonathonharrell/miri-ws-server/engine/core/database"
	"github.com/jonathonharrell/miri-ws-server/engine/models"

	"gopkg.in/mgo.v2/bson"
)

func CreateCharacter(character *parameters.Character, userId string) (status int, body []byte) {
	// @todo stub
	return http.StatusCreated, []byte("Whatever")
}

func CharacterOptions() (status int, body []byte) {
	options := make(map[string]interface{})
	options["races"] = content.Races
	options["genders"] = content.Genders
	options["aesthetic_traits"] = content.AestheticTraits
	options["functional_traits"] = content.FunctionalTraits

	res, _ := json.Marshal(options)

	return http.StatusOK, res
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
