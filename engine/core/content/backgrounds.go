package content

import (
	"github.com/jonathonharrell/miri-ws-server/engine/core/game"
)

var Backgrounds = map[string]game.Background{
	"NOMADDESERT": game.Background{
		Name:        "Nomad - Kai Ara",
		ID:          "NOMADDESERT",
		Description: "You are a nomad, wandering the Ara Wastes.",
		AllowAll:    true,
	},
}
