package admin

import (
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/app/game"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
)

type adminController struct{}

var Controller = adminController{}

func (c *adminController) Test(connection *game.Connection, game *game.Game, args *json.RawMessage) {
	logger.Write.Info("Admin user called test method")
}
