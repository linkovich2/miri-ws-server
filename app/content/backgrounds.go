package content

import (
	"github.com/jonathonharrell/miri-ws-server/app/core"
)

var Backgrounds = map[string]core.Background{
	"NOMADKAIARA": core.Background{
		Name:        "Nomad - Tulin Desert",
		ID:          "NOMADKAIARA",
		Description: "You are a nomad, currently traveling through the Tulin Desert.",
		AllowAll:    true,
	},
	"NOMADELDOREI": core.Background{
		Name:        "Nomad - Sapphire Coast",
		ID:          "NOMADELDOREI",
		Description: "You are a nomad, currently traveling through the Sapphire Coast.",
		AllowAll:    true,
	},
	"NOMADFOXEAR": core.Background{
		Name:        "Nomad - Soro Fields",
		ID:          "NOMADFOXEAR",
		Description: "You are a nomad, currently traveling through the Soro Fields.",
		AllowAll:    true,
	},
	"NOMADBRIREE": core.Background{
		Name:        "Nomad - Skyshroud Forest",
		ID:          "NOMADBRIREE",
		Description: "You are a nomad, currently traveling through the Skyshround Forest.",
		AllowAll:    true,
	},
}
