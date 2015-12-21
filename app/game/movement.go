package game

import (
	"encoding/json"
	"fmt"
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"strings"
)

type cMoveArgs struct {
	Direction string `json:"direction"`
}

func cMove(game *Game, c *Command) {
	params := cMoveArgs{}
	err := json.Unmarshal(*c.Args, &params)
	if err != nil {
		logger.Write.Error(err.Error())
		return
	}

	// @todo this should be timed out based on the character's move speed

	position, err := core.GetPosition(c.Character.Position)
	if err != nil {
		logger.Write.Error(err.Error())
		return
	}

	newPosition, err := position.Move(params.Direction)
	if err != nil {
		logger.Write.Error(err.Error())
		return
	}

	room := game.World.Realms[c.Character.Realm].Rooms[c.Character.Position]
	room.Remove(c.Connection.ID)

	c.Character.Position = newPosition.ToString()
	room = game.World.Realms[c.Character.Realm].Rooms[c.Character.Position]
	room.Add(c.Connection.ID)

	game.defaultMessage(c.Connection, c.Character, []string{})
	game.broadcastToRoom(c.Connection, getMovementMessage(c.Character, params.Direction), getMovementMessageBroadcast(c.Character, params.Direction), room)
}

func getMovementMessage(c *core.Character, dir string) string {
	// @todo this is mostly a stub for putting together a movement string for a character
	// ex. if the player is on horseback and we want to illustrate that
	return strings.Join([]string{ShortDescriptionForCharacter(c), " makes their way into the area."}, "")
}

func getMovementMessageBroadcast(c *core.Character, dir string) string {
	// @todod see above method
	return fmt.Sprintf("You make your way %s", dir)
}

func (game *Game) getAvailableDirections(r *core.Room, realm string) map[string]bool {
	directions := make(map[string]bool)

	positions := r.Position.AdjacentPositions()

	for k, p := range positions {
		if _, roomExists := game.World.Realms[realm].Rooms[p]; roomExists {
			directions[k] = true
		} else {
			directions[k] = false
		}
	}

	return directions
}
