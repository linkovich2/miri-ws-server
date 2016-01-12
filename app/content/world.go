package content

import (
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"strings"
)

type readRoom struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Size        int           `json:"size"`
	Details     []string      `json:"details"`
	Entities    []string      `json:"entities"`
	Position    core.Position `json:"position"`
}

func World() *core.World {
	data := MustAsset("json/world/world.json")
	world := &core.World{}
	json.Unmarshal(data, world)

	for k, r := range world.Realms {
		data := MustAsset("json/world/" + strings.ToLower(k) + ".json")
		a := map[string]readRoom{}
		err := json.Unmarshal(data, &a)
		if err != nil {
			panic(err)
		}

		r.Rooms = make(map[string]*core.Room)

		for positionString, tmp := range a {
			// @todo build entities out before attaching to room
			// @todo persistance layer probably has something to say about this
			r.Rooms[positionString] = &core.Room{
				Name:        tmp.Name,
				Description: tmp.Description,
				Size:        tmp.Size,
				Details:     tmp.Details,
				Position:    tmp.Position,
			}
		}
	}

	return world
}
