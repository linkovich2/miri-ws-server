package rooms

import (
	"github.com/jonathonharrell/miri-ws-server/app/core"
)

var Miri = map[string]*core.Room{
	"0:0:1": &core.Room{
		Position:    core.Position{0, 0, 1},
		Name:        "Kra Iree'a Outskirts",
		Description: "Description @todo heyyyyy",
		Size:        core.RoomSmall,
	},
	"1:1:1": &core.Room{
		Position:    core.Position{1, 1, 1},
		Name:        "Kra Iree'a",
		Description: "@todo Orange, blue and gold tents line the edge of this small traders outpost on the edge of the Tulinne Desert.",
		Size:        core.RoomLarge,
	},
	"1:2:1": &core.Room{
		Position:    core.Position{1, 2, 1},
		Name:        "Kra Iree'a Outskirts",
		Description: "Description stuff @todo.",
		Size:        core.RoomSmall,
	},
	"2:3:1": &core.Room{
		Position:    core.Position{2, 3, 1},
		Name:        "BIG TODO",
		Description: "Description stuff @todo.",
		Size:        core.RoomMedium,
	},
}
