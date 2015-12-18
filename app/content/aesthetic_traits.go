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
				Name:        "Brown",
				ID:          "BROWN",
				Description: "[PossessivePronoun] hair is brown.",
			},
			"BLONDE": core.AestheticTrait{
				Name:        "Blonde",
				ID:          "BLONDE",
				Description: "[PossessivePronoun] hair is bright blonde.",
			},
			"GREY": core.AestheticTrait{
				Name:        "Grey",
				ID:          "GREY",
				Description: "[PossessivePronoun] hair is grey and sullen.",
			},
			"BLACK": core.AestheticTrait{
				Name:        "Black",
				ID:          "BLACK",
				Description: "[PossessivePronoun] hair is jet black.",
			},
			"RED": core.AestheticTrait{
				Name:        "Red",
				ID:          "RED",
				Description: "[Pronoun] has ginger-red hair.",
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
				DisallowedGenders: []string{"M"},
				Description:       "[PossessivePronoun] hair is loosely arranged in a bun above [PossessivePronoun] head, strands falling about messily.",
			},
			"LONGSTRAIGHT": core.AestheticTrait{
				Name:        "Long, Straight",
				ID:          "LONGSTRAIGHT",
				Description: "[PossessivePronoun] hair is long and straight, reaching just below [PossessivePronoun] shoulders.",
			},
			"SHORT": core.AestheticTrait{
				Name:        "Short",
				ID:          "SHORT",
				Description: "[PossessivePronoun] hair is kept neatly close to [PossessivePronoun] head, less then half an inch long.",
			},
			"BALD": core.AestheticTrait{
				Name:        "Bald",
				ID:          "BALD",
				Description: "[PossessivePronoun] are bald, but at least [Pronoun]'s aerodynamic.",
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
				Name: "Long, braided beard",
				ID:   "LONGBRAIDEDBEARD",
			},
			"STUBBLE": core.AestheticTrait{
				Name: "Stubble",
				ID:   "STUBBLE",
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
				Description: "[PossessivePronoun] ears are long and point upwards.",
			},
			"SHORTUP": core.AestheticTrait{
				Name:        "Short, Pointed Up",
				ID:          "SHORTUP",
				Description: "[PossessivePronoun] ears are short and point upwards.",
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
				Description: "Freckles and sun spots dot [PossessivePronoun] skin.",
			},
			"REDNOSED": core.AestheticTrait{
				Name:        "Red-Nosed",
				ID:          "REDNOSED",
				Only:        "DWARF",
				Description: "[PossessivePronoun] nose is red, probably from drinking too much.",
			},
			"SCARREDFACE": core.AestheticTrait{
				Name:        "Facial Scars",
				ID:          "SCARREDFACE",
				Description: "[PossessivePronoun] face is scarred. Whether from battle or clumsiness, only [Pronoun] knows.",
			},
			"SCARREDBODY": core.AestheticTrait{
				Name:        "Scarred Body",
				ID:          "SCARREDBODY",
				Description: "[PossessivePronoun] body is scarred. There's a story behind each mark.",
			},
		},
	},
}
