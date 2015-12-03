package content

import (
	"github.com/jonathonharrell/miri-ws-server/app/core"
)

var Genders = map[string]core.Gender{
	"M": core.Gender{
		Name:            "Male",
		ID:              "M",
		DisallowedRaces: []string{"AUTOMATON"},
	},
	"F": core.Gender{
		Name:            "Female",
		ID:              "F",
		DisallowedRaces: []string{"AUTOMATON"},
	},
	"NA": core.Gender{
		Name: "N/A",
		ID:   "NA",
		Only: "AUTOMATON",
	},
}
