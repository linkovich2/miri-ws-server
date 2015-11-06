package characters

import (
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/game"
	//"github.com/jonathonharrell/miri-ws-server/app/content"
	"github.com/jonathonharrell/miri-ws-server/app/database"
	"github.com/jonathonharrell/miri-ws-server/app/logger"

	"gopkg.in/mgo.v2/bson"
)

type characterController struct{}

var Controller = characterController{}

func (c *characterController) List(connection *game.Connection, game *game.Game, args *json.RawMessage) {
	session, dbName := database.GetSession() // connect
	db := session.DB(dbName)
	defer session.Close()

	var characters []core.Character
	_ = db.C("characters").Find(bson.M{"user_id": connection.Socket.UserID}).All(&characters)

	logger.Write.Info("FOUND THIS SHIT %v", connection.Socket.UserID)
	logger.Write.Info("FOUND THESE SHITS %v", characters)
}
