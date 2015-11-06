package content

import (
	"github.com/jonathonharrell/miri-ws-server/app/core"
)

var Backgrounds = map[string]core.Background{
	"NOMADDESERT": core.Background{
		Name:        "Nomad - Kai Ara",
		ID:          "NOMADDESERT",
		Description: "You are a nomad, wandering the Ara Wastes.",
		AllowAll:    true,
	},
}
