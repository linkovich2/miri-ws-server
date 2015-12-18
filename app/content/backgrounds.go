package content

import (
	"github.com/jonathonharrell/miri-ws-server/app/core"
)

var Backgrounds = map[string]core.Background{
	"NOMADKAIARA": core.Background{
		Name:          "Nomad - Tulinne Desert",
		ID:            "NOMADKAIARA",
		Description:   "[Pronoun] [IsAre] a nomad, currently traveling through the Tulinne Desert.",
		AllowAll:      true,
		StartRealm:    "MIRI",
		StartPosition: "1:1:1",
	},
	"NOMADELDOREI": core.Background{
		Name:          "Nomad - Sapphire Coast",
		ID:            "NOMADELDOREI",
		Description:   "[INCOMPLETE] [Pronoun] [IsAre] a nomad, currently traveling through the Sapphire Coast.",
		AllowAll:      true,
		StartRealm:    "MIRI",
		StartPosition: "1:1:1",
	},
	"NOMADFOXEAR": core.Background{
		Name:          "Nomad - Soro Fields",
		ID:            "NOMADFOXEAR",
		Description:   "[INCOMPLETE] [Pronoun] [IsAre] a nomad, currently traveling through the Soro Fields.",
		AllowAll:      true,
		StartRealm:    "MIRI",
		StartPosition: "1:1:1",
	},
	"NOMADBRIREE": core.Background{
		Name:          "Nomad - Skyshroud Forest",
		ID:            "NOMADBRIREE",
		Description:   "[INCOMPLETE] [Pronoun] [IsAre] a nomad, currently traveling through the Skyshround Forest.",
		AllowAll:      true,
		StartRealm:    "MIRI",
		StartPosition: "1:1:1",
	},
	"BAHAMUTKNIGHT": core.Background{
		Name:        "Temple of Bahamut - Understudy",
		ID:          "BAHAMUTKNIGHT",
		Description: "[INCOMPLETE] [Pronoun] [IsAre] an understudy at the temple of Bahamut, perhaps in training to be a knight or cleric.",
		Prerequisites: core.BackgroundPrerequisites{
			FunctionalTraits: []string{"BAHAMUT"},
		},
		StartRealm:    "MIRI",
		StartPosition: "1:1:1",
	},
}
