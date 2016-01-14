package game

import (
	"github.com/jonathonharrell/miri-ws-server/app/logger"
)

func cAddTarget(game *Game, c *Command) {
	id, err := c.GetInput()
	if err != nil {
		logger.Write.Error(err.Error())
		return
	}

	room := game.World.Realms[c.Character.Realm].Rooms[c.Character.Position]
	if room.ValidTarget(id) {
		c.Character.Targets = append(c.Character.Targets, id)
		game.defaultMessage(c.Connection, c.Character, []string{})
	} else {
		logger.Write.Error("Invalid target [%s] selected by character [%s]", id, c.Character.Name)
	}
}

func cRemoveTarget(game *Game, c *Command) {
	id, err := c.GetInput()
	if err != nil {
		logger.Write.Error(err.Error())
		return
	}

	for i, val := range c.Character.Targets {
		if val == id {
			c.Character.Targets = append(c.Character.Targets[:i], c.Character.Targets[i+1:]...)
			break
		}
	}

	game.defaultMessage(c.Connection, c.Character, []string{})
}
