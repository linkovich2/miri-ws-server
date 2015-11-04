package content

import (
	"github.com/jonathonharrell/miri-ws-server/engine/core/game"
)

var FunctionalTraits = map[string]game.FunctionalTraitCategory{
	"RELIGION": game.FunctionalTraitCategory{
		Name:   "Religion",
		ID:     "RELIGION",
		Unique: true,
		Traits: map[string]game.FunctionalTrait{
			"NORELIGION": game.FunctionalTrait{
				Name:        "N/A",
				ID:          "NORELIGION",
				Description: "You are not a particularly religious individual",
				Points:      "0",
			},
			"BAHAMUT": game.FunctionalTrait{
				Name:        "Follower of Bahamut",
				ID:          "BAHAMUT",
				Description: "The God of Light",
				Points:      "-1",
			},
			"MORRHIGAN": game.FunctionalTrait{
				Name:        "Worshipper of Morrhigan",
				ID:          "MORRHIGAN",
				Description: "The Goddess of Death",
				Points:      "-1",
			},
		},
	},
	"RACIAL": game.FunctionalTraitCategory{
		Name:   "Inherent Racial Ability",
		ID:     "RACIAL",
		DisallowedRaces: []string{"HUMAN"},
		Traits: map[string]game.FunctionalTrait{
			"DARKVISION": game.FunctionalTrait{
				Name:        "Darkvision",
				ID:          "DARKVISION",
				Description: "Your vision is fair even in dim light and partial darkness.",
				Points:      "0",
				Required:    true,
			},
			"FTFORESTS": game.FunctionalTrait{
				Name:        "Natural Attunement - Forests",
				ID:          "FTFORESTS",
				Only:        "ELF",
				Required:    true,
				Description: "You have a natural ability to navigate forests.",
				Points:      "0",
			},
			"FTCAVERNS": game.FunctionalTrait{
				Name:        "Favored Terrain - Mountains",
				ID:          "FTCAVERNS",
				Only:        "DWARF",
				Required:    true,
				Description: "You have a natural ability to navigate caverns, mountains, and other rocky terrain.",
				Points:      "0",
			},
		},
	},
}
