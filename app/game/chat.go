package game

import (
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"github.com/jonathonharrell/miri-ws-server/app/util"
	"strings"
)

func cSpacialChat(game *Game, c *Command) {
	input, err := c.GetInput()
	if err != nil {
		logger.Write.Error(err.Error())
		return
	}

	input = util.Capitalize(input)

	broadcast := []string{"yell", "shout"}
	actions := map[string]string{
		"say":   "say",
		"yell":  "yell",
		"shout": "shout",
	}

	if _, exists := actions[c.Value]; !exists {
		logger.Write.Error("Got chat command [%s] that does not exist in the action list!", c.Value)
		return
	}

	action := actions[c.Value]
	pos, _ := core.GetPosition(c.Character.Position)
	room := game.World.Realms[c.Character.Realm].Rooms[c.Character.Position]
	desc := strings.Join([]string{"<strong>", strings.ToLower(c.Character.ShortDescription()), "</strong>"}, "")
	game.broadcastToRoom(
		c.Connection,
		c.Character,
		strings.Join([]string{desc, " ", action, "s, \"<", action, ">", input, "</", action, ">\""}, ""),
		strings.Join([]string{"You ", action, ", \"<", action, ">", input, "</", action, ">\""}, ""),
		room,
	)

	if ok, _ := util.InArray(c.Value, broadcast); ok {
		for direction, value := range pos.AdjacentPositions() {
			if room, exists := game.World.Realms[c.Character.Realm].Rooms[value]; exists && room.Active {
				d, err := core.GetOppositeDirection(direction)
				if err != nil {
					logger.Write.Error(err.Error())
					continue
				}

				room.Broadcast(
					strings.Join([]string{desc, " ", action, "s from the ", d, ", \"<", action, ">", input, "</", action, ">\""}, ""),
					game.World.GetSendCallback(),
				)
			}
		}
	}
}
