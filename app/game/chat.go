package game

import (
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"strings"
)

type cSayArgs struct {
	Input string `json:"input"`
}

func cSay(game *Game, c *Command) {
	params := cSayArgs{}
	err := json.Unmarshal(*c.Args, &params)
	if err != nil {
		logger.Write.Error(err.Error())
		return
	}

	room := game.World.Realms[c.Character.Realm].Rooms[c.Character.Position]
	desc := strings.Join([]string{"<strong>", strings.ToLower(ShortDescriptionForCharacter(c.Character)), "</strong>"}, "")
	game.broadcastToRoom(
		c.Connection,
		strings.Join([]string{desc, " says, \"<say>", params.Input, "</say>\""}, ""),
		strings.Join([]string{"You say, \"<say>", params.Input, "</say>\""}, ""),
		room,
	)
}

func cYell(game *Game, c *Command) {
	// @todo stub
	logger.Write.Info("Called yell")
}
