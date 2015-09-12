package engine

import (
	"github.com/jonathonharrell/miri-ws-server/engine/util/loader"
)

type Gender struct {
	Name            string   `json:"name"`
	ID              string   `json:"id"`
	DisallowedRaces []string `json:"disallowed_races"`
}

type GenderShort struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

var genders []Gender

func InitGenders() {
	loader.Grab("genders.json", &genders)
}
