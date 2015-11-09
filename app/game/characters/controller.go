package characters

import (
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/app/content"
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/database"
	"github.com/jonathonharrell/miri-ws-server/app/game"
	"github.com/jonathonharrell/miri-ws-server/app/logger"

	"gopkg.in/mgo.v2/bson"
)

type (
	characterController struct{}
	optionsParams       struct {
		Get string `json:"get"`
	}
	deleteParams struct {
		Id string `json:"id"`
	}
	selectParams   deleteParams
	validateParams core.Character
)

var Controller = characterController{}

func (c *characterController) List(connection *game.Connection, game *game.Game, args *json.RawMessage) {
	session, dbName := database.GetSession() // connect
	db := session.DB(dbName)
	defer session.Close()

	var characters []core.Character
	_ = db.C("characters").Find(bson.M{"user_id": connection.Socket.UserID}).All(&characters)

	res, _ := json.Marshal(characters)
	connection.Socket.Send(res)
}

func (c *characterController) Delete(connection *game.Connection, game *game.Game, args *json.RawMessage) {
	// @todo stub
}

func (c *characterController) Select(connection *game.Connection, game *game.Game, args *json.RawMessage) {
	// @todo stub
}

func (c *characterController) Validate(connection *game.Connection, game *game.Game, args *json.RawMessage) {
	// @todo stub
}

func (c *characterController) Options(connection *game.Connection, game *game.Game, args *json.RawMessage) {
	var body interface{}
	params := optionsParams{}
	err := json.Unmarshal(*args, &params)
	if err != nil {
		logger.Write.Error(err.Error()) // @todo handle json malformed or something like that
	}

	switch params.Get {
	case "races":
		body = content.Races
	case "genders":
		body = content.Genders
	case "aesthetic_traits":
		body = content.AestheticTraits
	case "functional_traits":
		body = content.FunctionalTraits
	case "backgrounds":
		body = content.Backgrounds
	default:
		body = content.Races
	}

	res, _ := json.Marshal(body)
	connection.Socket.Send(res)
}
