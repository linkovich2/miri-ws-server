package services

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/jonathonharrell/miri-ws-server/engine/api/parameters"
	"github.com/jonathonharrell/miri-ws-server/engine/core/content"
	"github.com/jonathonharrell/miri-ws-server/engine/core/database"
	"github.com/jonathonharrell/miri-ws-server/engine/models"
)

func CreateCharacter(character *parameters.Character, userId string) (status int, body []byte) {
	// valid, errors := validate(character)
	// @todo stub
	return http.StatusCreated, []byte("Whatever")
}

func CharacterOptions(request *http.Request) (status int, body []byte) {
	options := getOptions(request.URL.Query().Get("for"))
	if options == nil {
		return http.StatusBadRequest, []byte{}
	}

	res, _ := json.Marshal(options)
	return http.StatusOK, res
}

func ListCharacters(userId string) (status int, body []byte) {
	session, dbName := database.GetSession() // connect
	db := session.DB(dbName)
	defer session.Close()

	u := models.User{}
	err := db.C("users").Find(bson.M{"_id": bson.ObjectIdHex(userId)}).One(&u)

	if err != nil { // no existing user
		return http.StatusUnauthorized, nil
	}

	res, _ := json.Marshal(u.Characters)
	return http.StatusOK, res
}

func DeleteCharacter() {
	// @todo stub
}

func getOptions(query string) interface{} {
	switch query {
	case "races":
		return content.Races
	case "genders":
		return content.Genders
	case "aesthetic_traits":
		return content.AestheticTraits
	case "functional_traits":
		return content.FunctionalTraits
	case "backgrounds":
		return content.Backgrounds
	default:
		return nil
	}
}
