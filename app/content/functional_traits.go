package content

import (
	"github.com/jonathonharrell/miri-ws-server/app/core"
)

var FunctionalTraits = map[string]core.FunctionalTraitCategory{
	"RELIGION": core.FunctionalTraitCategory{
		Name:   "Religion",
		ID:     "RELIGION",
		Unique: true,
		Traits: map[string]core.FunctionalTrait{
			"BAHAMUT": core.FunctionalTrait{
				Name:        "Follower of Bahamut",
				ID:          "BAHAMUT",
				Description: "The God of Light",
				Points:      "-1",
			},
			"MORRHIGAN": core.FunctionalTrait{
				Name:        "Worshipper of Morrhigan",
				ID:          "MORRHIGAN",
				Description: "The Goddess of Death",
				Points:      "-1",
			},
		},
	},
	"RACIAL": core.FunctionalTraitCategory{
		Name:            "Inherent Racial Ability",
		ID:              "RACIAL",
		DisallowedRaces: []string{"HUMAN"},
		Traits: map[string]core.FunctionalTrait{
			"DARKVISION": core.FunctionalTrait{
				Name:        "Darkvision",
				ID:          "DARKVISION",
				Description: "[PossessivePronoun] vision is fair even in dim light and partial darkness.",
				Points:      "0",
				Required:    true,
			},
			"FTFORESTS": core.FunctionalTrait{
				Name:        "Natural Attunement - Forests",
				ID:          "FTFORESTS",
				Only:        "ELF",
				Required:    true,
				Description: "[Pronoun] [HaveHas] an innate ability to navigate forests or woodland terrain.",
				Points:      "0",
			},
			"FTCAVERNS": core.FunctionalTrait{
				Name:        "Favored Terrain - Mountains",
				ID:          "FTCAVERNS",
				Only:        "DWARF",
				Required:    true,
				Description: "[Pronoun] [HaveHas] an innate ability to navigate mountains and other rocky terrain.",
				Points:      "0",
			},
		},
	},
}
