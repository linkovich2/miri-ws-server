package game

import (
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"strings"
)

type cBasicChatArgs struct {
	Input string `json:"input"`
}

func cSay(game *Game, c *Command) {
	params := cBasicChatArgs{}
	err := json.Unmarshal(*c.Args, &params)
	if err != nil {
		logger.Write.Error(err.Error())
		return
	}

	room := game.World.Realms[c.Character.Realm].Rooms[c.Character.Position]
	desc := strings.Join([]string{"<strong>", strings.ToLower(ShortDescriptionForCharacter(c.Character)), "</strong>"}, "")
	game.broadcastToRoom(
		c.Connection,
		strings.Join([]string{desc, " says, \"<say>", params.Input, "</say>\""}, ""),
		strings.Join([]string{"You say, \"<say>", params.Input, "</say>\""}, ""),
		room,
	)
}

func cYell(game *Game, c *Command) {
	params := cBasicChatArgs{}
	err := json.Unmarshal(*c.Args, &params)
	if err != nil {
		logger.Write.Error(err.Error())
		return
	}

	pos, _ := core.GetPosition(c.Character.Position)
	room := game.World.Realms[c.Character.Realm].Rooms[c.Character.Position]
	desc := strings.Join([]string{"<strong>", strings.ToLower(ShortDescriptionForCharacter(c.Character)), "</strong>"}, "")
	game.broadcastToRoom(
		c.Connection,
		strings.Join([]string{desc, " yells, \"<yell>", params.Input, "</yell>\""}, ""),
		strings.Join([]string{"You yell, \"<yell>", params.Input, "</yell>\""}, ""),
		room,
	)

	oppositeDirections := map[string]string{
		"north":"south",
		"south":"north",
		"east":"west",
		"west":"east",
		"northeast":"southwest",
		"northwest":"southeast",
		"southeast":"northwest",
		"southwest":"northeast",
	}

	for direction, value := range pos.AdjacentPositions() {
		if room, exists := game.World.Realms[c.Character.Realm].Rooms[value]; exists {
			room.Broadcast(
				strings.Join([]string{desc, " yells from the ", oppositeDirections[direction], ", \"<yell>", params.Input, "</yell>\""}, ""),
				game.World.GetSendCallback(),
			)
		}
	}
}
