package game

import (
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/util"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"strings"
	"time"
)

var (
	movementMessages = map[string]map[string][]string{
		"start": map[string][]string{
			"walk": []string{
				"You make your way to the [Direction]",
				"You continue [Direction], walking through the [Detail]",
			},
		},
		"broadcastStart": map[string][]string{
			"walk": []string{
				"[Description] begins walking to the [Direction]",
			},
		},
		"arrive": map[string][]string{
			"walk": []string{
				"You arrive at [RoomName]",
			},
		},
		"broadcastArrive": map[string][]string{
			"walk": []string{
				"You notice [Description] entering the area.",
			},
		},
	}
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
		c.Character.State = core.StateMoving
		room := game.World.Realms[c.Character.Realm].Rooms[c.Character.Position]
		game.broadcastToRoom(
			c.Connection,
			getMovementMessage(c.Character, room, params.Direction, "broadcastStart"),
			getMovementMessage(c.Character, room, params.Direction, "start"),
			room,
		)

		// @todo as part of this sleep we may want to check for movement impeding stuff, such as being on a mountain
		// we may also want to check for exit traps or anything like that
		time.Sleep((time.Duration(c.Character.GetSpeed()) + time.Duration(room.GetSpeedMod())) * time.Second) // wait for character's move speed

		room.Remove(c.Connection.ID)

		c.Character.Position = newPosition.ToString()
		room = game.World.Realms[c.Character.Realm].Rooms[c.Character.Position]
		room.Add(c.Connection.ID)

		c.Character.State = core.StateDefault
		game.defaultMessage(c.Connection, c.Character, []string{})
		game.broadcastToRoom(
			c.Connection,
			getMovementMessage(c.Character, room, params.Direction, "broadcastArrive"),
			getMovementMessage(c.Character, room, params.Direction, "arrive"),
			room,
		)
	}()
}

func getMovementMessage(c *core.Character, room *core.Room, dir, pool string) string {
	res, err := util.Sample(movementMessages[pool][c.GetMovementStyle()])
	if err != nil {
		logger.Write.Error(err.Error()) // wtf happened
		return ""
	}
	res = strings.Replace(res, "[Description]", ShortDescriptionForCharacter(c), -1)
	res = strings.Replace(res, "[RoomName]", room.Name, -1)
	res = strings.Replace(res, "[Direction]", dir, -1)
	return res
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
