package content

import (
	"github.com/jonathonharrell/miri-ws-server/engine/core/game"
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
