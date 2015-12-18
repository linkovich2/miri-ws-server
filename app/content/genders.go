package content

import (
	"github.com/jonathonharrell/miri-ws-server/app/core"
)

var Genders = map[string]core.Gender{
	"M": core.Gender{
		Name:            "Male",
		ID:              "M",
		DisallowedRaces: []string{"AUTOMATON"},
		Human:           "man",
		Scientific:      "male",
		Pronoun:         "he",
		Possessive:      "his",
	},
	"F": core.Gender{
		Name:            "Female",
		ID:              "F",
		DisallowedRaces: []string{"AUTOMATON"},
		Human:           "woman",
		Scientific:      "female",
		Pronoun:         "she",
		Possessive:      "her",
	},
	"NA": core.Gender{
		Name:       "N/A",
		ID:         "NA",
		Only:       "AUTOMATON",
		Human:      "automaton",
		Scientific: "automaton",
		Pronoun:    "it",
		Possessive: "it's",
	},
}
