package engine

import (
	"github.com/jonathonharrell/miri-ws-server/engine/util/loader"
)

type Race struct {
	Name        string
	Description string
	ID          string
}

var races []Race

func InitRaces() {
	loader.Grab("races.json", &races)
}
