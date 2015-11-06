package content

import (
	"github.com/jonathonharrell/miri-ws-server/app/core/game"
)

var Genders = map[string]game.Gender{
	"M": game.Gender{
		Name: "Male",
		ID:   "M",
	},
	"F": game.Gender{
		Name: "Female",
		ID:   "F",
	},
}
