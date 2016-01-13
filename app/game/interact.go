package game

import (
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
)

type cInteractArgs struct {
	Target string `json:"target"`
	Action string `json:"action"`
}

func cInteract(game *Game, c *Command) {
	params := cInteractArgs{}
	err := json.Unmarshal(*c.Args, &params)
	if err != nil {
		logger.Write.Error(err.Error())
		return
	}

	room := game.World.Realms[c.Character.Realm].Rooms[c.Character.Position]
	room.Interact(c.Character, params.Target, params.Action, game.World.GetSendCallback())
}
