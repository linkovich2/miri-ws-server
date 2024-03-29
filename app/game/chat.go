package game

import (
	"bytes"
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"github.com/jonathonharrell/miri-ws-server/app/util"
	"strings"
	"time"
	"math/rand"
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
	desc := strings.Join([]string{"<strong>", strings.ToLower(c.Character.ShortDescriptionWithName()), "</strong>"}, "")
	target := ""

	if len(c.Character.Targets) > 0 {
		buffer := bytes.NewBuffer([]byte{})
		var index int
		for _, t := range c.Character.Targets {
			index = index + 1
			e, err := room.GetTarget(t) // grab entity from room
			if err != nil {
				continue
			}

			if len(e.Properties.Matching(c.Value + "Response")) > 0 || len(e.Properties.Matching("chatResponse")) > 0 {
				go func() {
					rand.Seed(time.Now().UTC().UnixNano())
					time.Sleep(time.Duration(3) * time.Second)

					options   := append(e.Properties.Matching(c.Value + "Response"), e.Properties.Matching("chatResponse")...)
					selection := options[rand.Intn(len(options))]

					game.broadcastToRoom(
						c.Connection,
						c.Character,
						strings.Replace(selection.Value, "[RACE]", strings.ToLower(c.Character.Race), -1),
						strings.Replace(selection.Value, "[RACE]", strings.ToLower(c.Character.Race), -1),
						room,
					)
				}()
			}

			if index > 1 {
				buffer.Write([]byte("and"))
			}
			buffer.Write([]byte(" to the "))
			buffer.Write([]byte(e.Name))
			buffer.Write([]byte(" "))
		}
		target = buffer.String()
	}

	game.broadcastToRoom(
		c.Connection,
		c.Character,
		strings.Join([]string{desc, " ", action, "s", target, ", \"<", action, ">", input, "</", action, ">\""}, ""),
		strings.Join([]string{"You ", action, target, ", \"<", action, ">", input, "</", action, ">\""}, ""),
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
					strings.Join([]string{desc, " ", action, "s", target, " from the ", d, ", \"<", action, ">", input, "</", action, ">\""}, ""),
					game.World.GetSendCallback(),
				)
			}
		}
	}
}
