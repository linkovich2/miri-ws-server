package characters

import (
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/app/game"
	// "github.com/jonathonharrell/miri-ws-server/app/database"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
)

type characterController struct{}

var Controller = characterController{}

func (c *characterController) List(connection *game.Connection, game *game.Game, args *json.RawMessage) {
	// session, dbName := database.GetSession() // connect
	// db := session.DB(dbName)
	// defer session.Close()

	logger.Write.Info("FOUND THIS SHIT %v", connection.Socket.UserID)
}
