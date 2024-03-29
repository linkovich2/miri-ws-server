package game

import (
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"github.com/jonathonharrell/miri-ws-server/app/util"
	"strings"
	"time"
)

var (
	movementMessages = map[string]map[string][]string{
		"start": map[string][]string{
			"walk": []string{
				"You begin walking to the [Direction].",
				"You continue [Direction], walking through the [Detail].",
				"You look back on the [Detail] as you make your way [Direction].",
			},
		},
		"broadcastStart": map[string][]string{
			"walk": []string{
				"[Description] begins walking to the [Direction].",
			},
		},
		"arrive": map[string][]string{
			"walk": []string{
				"You arrive at [RoomName].",
			},
		},
		"broadcastArrive": map[string][]string{
			"walk": []string{
				"You notice [Description] entering the area.",
			},
		},
	}
)

func cMove(game *Game, c *Command) {
	if !c.Character.HasState(core.StateMoving) {
		move(game, c)
	} else {
		logger.Write.Error("Character [%s] from connection [%s] tried to move when they were already moving!", c.Character.Name, c.Connection.ID)
	}
}

func move(game *Game, c *Command) {
	d, err := c.GetInput()
	if err != nil {
		logger.Write.Error(err.Error())
		return
	}

	d = core.GetDirectionFromVariations(d)

	position, err := core.GetPosition(c.Character.Position)
	if err != nil {
		logger.Write.Error(err.Error())
		return
	}

	newPosition, err := position.Move(d)
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
		c.Character.AddState(core.StateMoving)
		room := game.World.Realms[c.Character.Realm].Rooms[c.Character.Position]
		game.broadcastToRoom(
			c.Connection,
			c.Character,
			getMovementMessage(c.Character, room, d, "broadcastStart"),
			getMovementMessage(c.Character, room, d, "start"),
			room,
		)

		// @todo as part of this sleep we may want to check for movement impeding stuff, such as being on a mountain
		// we may also want to check for exit traps or anything like that
		time.Sleep((time.Duration(c.Character.GetSpeed()) + time.Duration(room.GetSpeedMod())) * time.Second) // wait for character's move speed

		room.Remove(c.Connection.ID)

		c.Character.Position = newPosition.ToString()
		room = game.World.Realms[c.Character.Realm].Rooms[c.Character.Position]
		room.Add(c.Connection.ID, c.Character)

		c.Character.Targets = []string{}
		c.Character.RemoveState(core.StateMoving)
		game.broadcastToRoom(
			c.Connection,
			c.Character,
			getMovementMessage(c.Character, room, d, "broadcastArrive"),
			getMovementMessage(c.Character, room, d, "arrive"),
			room,
		)
	}()
}

func getMovementMessage(c *core.Character, room *core.Room, dir, pool string) string {
	res, _ := util.Sample(movementMessages[pool][c.GetMovementStyle()])

	desc := strings.Join([]string{"<strong>", strings.ToLower(c.ShortDescriptionWithName()), "</strong>"}, "")
	res = strings.Replace(res, "[Description]", desc, -1)

	locale := strings.Join([]string{"<em>", room.Name, "</em>"}, "")
	res = strings.Replace(res, "[RoomName]", locale, -1)

	res = strings.Replace(res, "[Direction]", dir, -1)

	if len(room.Details) > 0 {
		detail, _ := util.Sample(room.Details)
		res = strings.Replace(res, "[Detail]", detail, -1)
	}

	if pool == "broadcastStart" || pool == "broadcastArrive" {
		return strings.Join([]string{"<default>", res, "</default>"}, "")
	}

	return strings.Join([]string{"<movement>", res, "</movement>"}, "")
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
