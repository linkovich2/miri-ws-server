package game

import (
	"encoding/json"
	"fmt"
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"strings"
	"time"
)

type cMoveArgs struct {
	Direction string `json:"direction"`
}

func cMove(game *Game, c *Command) {
	if c.Character.State != core.StateMoving {
		move(game, c)
	} else {
		logger.Write.Error("Character [%s] from connection [%s] tried to move when they were already moving!", c.Character.Name, c.Connection.ID)
	}
}

func move(game *Game, c *Command) {
	params := cMoveArgs{}
	err := json.Unmarshal(*c.Args, &params)
	if err != nil {
		logger.Write.Error(err.Error())
		return
	}

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

	if _, exists := game.World.Realms[c.Character.Realm].Rooms[newPosition.ToString()]; !exists {
		logger.Write.Error(
			"Character [%s] connection [%s] tried to move to a room that doesn't exist!",
			c.Character.Name,
			c.Connection.ID,
		)
		return
	}

	go func() {
		startMovingMessage := response{Messages: []string{"You begin moving"}}
		res, _ := json.Marshal(&startMovingMessage)
		c.Connection.Send(res)

		c.Character.State = core.StateMoving
		room := game.World.Realms[c.Character.Realm].Rooms[c.Character.Position]
		// @todo as part of this sleep we may want to check for movement impeding stuff, such as being on a mountain
		// we may also want to check for exit traps or anything like that
		time.Sleep((time.Duration(c.Character.GetSpeed()) + time.Duration(room.GetSpeedMod())) * time.Second) // wait for character's move speed

		room.Remove(c.Connection.ID)

		c.Character.Position = newPosition.ToString()
		room = game.World.Realms[c.Character.Realm].Rooms[c.Character.Position]
		room.Add(c.Connection.ID)

		c.Character.State = core.StateDefault
		game.defaultMessage(c.Connection, c.Character, []string{})
		game.broadcastToRoom(c.Connection, getMovementMessage(c.Character, params.Direction), getMovementMessageBroadcast(c.Character, params.Direction), room)
	}()
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
