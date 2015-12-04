package rooms

import (
	"github.com/jonathonharrell/miri-ws-server/app/core"
)

var Miri = map[string]core.Room{
	"1:1:1": core.Room{
		Position:    core.Position{1, 1, 1},
		Name:        "Kra Iree'a",
		Description: "@todo Orange, blue and gold tents line the edge of this small traders outpost on the edge of the Tulinne Desert.",
	},
	"1:2:1": core.Room{
		Position:    core.Position{1, 2, 1},
		Name:        "Kra Iree'a Outskirts",
		Description: "Description stuff @todo.",
	},
}
