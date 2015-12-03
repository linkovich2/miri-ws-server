package content

import (
	"github.com/jonathonharrell/miri-ws-server/app/core"
)

var Genders = map[string]core.Gender{
	"M": core.Gender{
		Name: "Male",
		ID:   "M",
	},
	"F": core.Gender{
		Name: "Female",
		ID:   "F",
	},
	"NA": core.Gender{
		Name: "N/A",
		ID:   "NA",
		Only: "AUTOMATON",
	},
}
