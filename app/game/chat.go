package game

import (
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"strings"
)

func cSay(game *Game, c *Command) {
	input, err := c.GetInput()
	if err != nil {
		logger.Write.Error(err.Error())
		return
	}

	room := game.World.Realms[c.Character.Realm].Rooms[c.Character.Position]
	desc := strings.Join([]string{"<strong>", strings.ToLower(ShortDescriptionForCharacter(c.Character)), "</strong>"}, "")
	game.broadcastToRoom(
		c.Connection,
		c.Character,
		strings.Join([]string{desc, " says, \"<say>", input, "</say>\""}, ""),
		strings.Join([]string{"You say, \"<say>", input, "</say>\""}, ""),
		room,
	)
}

func cYell(game *Game, c *Command) {
	input, err := c.GetInput()
	if err != nil {
		logger.Write.Error(err.Error())
		return
	}

	pos, _ := core.GetPosition(c.Character.Position)
	room := game.World.Realms[c.Character.Realm].Rooms[c.Character.Position]
	desc := strings.Join([]string{"<strong>", strings.ToLower(ShortDescriptionForCharacter(c.Character)), "</strong>"}, "")
	game.broadcastToRoom(
		c.Connection,
		c.Character,
		strings.Join([]string{desc, " yells, \"<yell>", input, "</yell>\""}, ""),
		strings.Join([]string{"You yell, \"<yell>", input, "</yell>\""}, ""),
		room,
	)

	for direction, value := range pos.AdjacentPositions() {
		if room, exists := game.World.Realms[c.Character.Realm].Rooms[value]; exists {
			d, err := core.GetOppositeDirection(direction)
			if err != nil {
				logger.Write.Error(err.Error())
				continue
			}

			room.Broadcast(
				strings.Join([]string{desc, " yells from the ", d, ", \"<yell>", input, "</yell>\""}, ""),
				game.World.GetSendCallback(),
			)
		}
	}
}
