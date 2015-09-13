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

var races = make(map[string]Race)

func InitRaces() {
	arr := []Race{}
	loader.Grab("races.json", &arr)
	for _, val := range arr {
		races[val.ID] = val
	}
}
