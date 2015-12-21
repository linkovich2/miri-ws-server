package world

import (
	"github.com/jonathonharrell/miri-ws-server/app/content/world/rooms"
	"github.com/jonathonharrell/miri-ws-server/app/core"
)

var Miri = core.World{
	Name: "Miri",
	Realms: map[string]*core.Realm{
		"MIRI": &core.Realm{
			ID:    "MIRI",
			Name:  "The Miri",
			Rooms: rooms.Miri,
		},
		"SHADOW": &core.Realm{
			ID:    "SHADOW",
			Name:  "The Shadow",
			Rooms: rooms.Shadow,
		},
		"LIGHT": &core.Realm{
			ID:    "LIGHT",
			Name:  "The Light",
			Rooms: rooms.Light,
		},
		"CHAOS": &core.Realm{
			ID:    "CHAOS",
			Name:  "The Chaos",
			Rooms: rooms.Chaos,
		},
		"FAE": &core.Realm{
			ID:    "FAE",
			Name:  "The Fae",
			Rooms: rooms.Fae,
		},
	},
}
