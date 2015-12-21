package game

import (
	"encoding/json"
	"fmt"
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
)

type cMoveArgs struct {
	Direction string `json:"direction"`
}

func cMove(c *Command) {
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

	c.Character.Position = newPosition.ToString()
	defaultMessage(c.Connection, c.Character, []string{fmt.Sprintf("You make your way %s", params.Direction)})
}

func GetAvailableDirections(r *core.Room, realm string) map[string]bool {
	directions := make(map[string]bool)

	positions := r.Position.AdjacentPositions()

	for k, p := range positions {
		if _, roomExists := miri.Realms[realm].Rooms[p]; roomExists {
			directions[k] = true
		} else {
			directions[k] = false
		}
	}

	return directions
}
