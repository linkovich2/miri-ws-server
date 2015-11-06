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
}
