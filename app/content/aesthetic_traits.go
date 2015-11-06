package content

import (
	"github.com/jonathonharrell/miri-ws-server/app/core"
)

var AestheticTraits = map[string]core.AestheticTraitCategory{
	"HAIRCOLOR": core.AestheticTraitCategory{
		ID:      "HAIRCOLOR",
		Name:    "Hair Color",
		Unique:  true,
		Minimum: 1,
		Traits: map[string]core.AestheticTrait{
			"BROWN": core.AestheticTrait{
				Name:  "Brown",
				ID:    "BROWN",
				Image: 1,
			},
			"BLONDE": core.AestheticTrait{
				Name:  "Blonde",
				ID:    "BLONDE",
				Image: 2,
			},
			"GREY": core.AestheticTrait{
				Name:  "Grey",
				ID:    "GREY",
				Image: 3,
			},
			"BLACK": core.AestheticTrait{
				Name:  "Black",
				ID:    "BLACK",
				Image: 4,
			},
		},
	},
	"HAIRSTYLE": core.AestheticTraitCategory{
		ID:      "HAIRSTYLE",
		Name:    "Hair Style",
		Unique:  true,
		Minimum: 1,
		Traits: map[string]core.AestheticTrait{
			"MESSYBUN": core.AestheticTrait{
				Name:              "Messy Bun",
				ID:                "MESSYBUN",
				Image:             0,
				DisallowedGenders: []string{"M"},
				Description:       "Your hair is loosely arranged in a bun above your head, strands falling about messily.",
			},
			"LONGSTRAIGHT": core.AestheticTrait{
				Name:        "Long, Straight",
				ID:          "LONGSTRAIGHT",
				Image:       0,
				Description: "Your hair is long and straight, reaching just below shoulder-length.",
			},
			"SHORT": core.AestheticTrait{
				Name:        "Short",
				ID:          "SHORT",
				Image:       0,
				Description: "Your hair is kept neatly close to your head, less then half an inch long.",
			},
			"BALD": core.AestheticTrait{
				Name:        "Bald",
				ID:          "BALD",
				Image:       0,
				Description: "You are bald, but at least you're aerodynamic.",
			},
		},
	},
	"FACIALHAIR": core.AestheticTraitCategory{
		ID:                "FACIALHAIR",
		Name:              "Facial Hair",
		Unique:            true,
		DisallowedRaces:   []string{"ELF"},
		DisallowedGenders: []string{"F"},
		Traits: map[string]core.AestheticTrait{
			"LONGBRAIDEDBEARD": core.AestheticTrait{
				Name:  "Long, braided beard",
				ID:    "LONGBRAIDEDBEARD",
				Image: 0,
			},
			"STUBBLE": core.AestheticTrait{
				Name:  "Stubble",
				ID:    "STUBBLE",
				Image: 0,
			},
		},
	},
	"ELFEARS": core.AestheticTraitCategory{
		ID:      "ELFEARS",
		Name:    "Ear Style",
		Unique:  true,
		Only:    "ELF",
		Minimum: 1,
		Traits: map[string]core.AestheticTrait{
			"LONGUP": core.AestheticTrait{
				Name:        "Long, Pointed Up",
				ID:          "LONGUP",
				Image:       0,
				Description: "Your ears are long and point upwards.",
			},
			"SHORTUP": core.AestheticTrait{
				Name:        "Short, Pointed Up",
				ID:          "SHORTUP",
				Image:       0,
				Description: "Your ears are short and point upwards.",
			},
		},
	},
	"MISC": core.AestheticTraitCategory{
		ID:     "MISC",
		Name:   "Miscellaneous",
		Unique: false,
		Traits: map[string]core.AestheticTrait{
			"FRECKLES": core.AestheticTrait{
				Name:        "Freckles",
				ID:          "FRECKLES",
				Image:       0,
				Description: "Freckles and sun spots dot your skin.",
			},
			"REDNOSED": core.AestheticTrait{
				Name:        "Red-Nosed",
				ID:          "REDNOSED",
				Only:        "DWARF",
				Image:       0,
				Description: "Your nose is red, probably from drinking too much.",
			},
			"SCARREDFACE": core.AestheticTrait{
				Name:        "Facial Scars",
				ID:          "SCARREDFACE",
				Image:       0,
				Description: "Your face is scarred. Whether from battle or clumsiness, only you know.",
			},
			"SCARREDBODY": core.AestheticTrait{
				Name:        "Scarred Body",
				ID:          "SCARREDBODY",
				Image:       0,
				Description: "Your body is scarred. I'm sure there's a story behind each mark.",
			},
		},
	},
}
