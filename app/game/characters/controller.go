package characters

import (
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/app/content"
	"github.com/jonathonharrell/miri-ws-server/app/core"
	// "github.com/jonathonharrell/miri-ws-server/app/database"
	"github.com/jonathonharrell/miri-ws-server/app/game"
	"github.com/jonathonharrell/miri-ws-server/app/logger"

	// "gopkg.in/mgo.v2/bson"
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
	res, _ := json.Marshal(connection.Socket.User.Characters)
	connection.Socket.Send(res)
}

func (c *characterController) Delete(connection *game.Connection, game *game.Game, args *json.RawMessage) {
	// @todo stub
}

func (c *characterController) Select(connection *game.Connection, game *game.Game, args *json.RawMessage) {
	// @todo stub
}

func (c *characterController) Create(connection *game.Connection, game *game.Game, args *json.RawMessage) {
	character := core.Character{}
	err := json.Unmarshal(*args, &character)
	if err != nil {
		logger.Write.Error(err.Error()) // @todo handle json malformed or something like that
	}

	// @todo temp
	logger.Write.Info("Received a create character message for character \"%v, %v, %v: %v\"", character.Name, character.Race, character.Gender, character.Background)

	// @todo validate against existing characters (if 3 or greater and !connection.Socket.Admin, fail)
	// @todo validate the character itself
	// @todo save the character in the database
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
