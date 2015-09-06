package engine

import (
	"github.com/jonathonharrell/miri-ws-server/engine/util/loader"
)

type Race struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	LongDescription string `json:"long_description"`
	ID              string `json:"id, omitempty"`
}

var races []Race

func InitRaces() {
	loader.Grab("races.json", &races)
}
