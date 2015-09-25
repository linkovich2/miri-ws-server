package engine

import (
	"github.com/jonathonharrell/miri-ws-server/engine/util/loader"
)

type Gender struct {
	Name            string   `json:"name"`
	ID              string   `json:"id"`
	DisallowedRaces []string `json:"disallowed_races"`
	Only            string   `json:"only"`
}

type GenderShort struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

var genders = make(map[string]Gender)

func InitGenders() {
	arr := []Gender{}
	loader.Grab("genders.json", &arr)
	for _, val := range arr {
		genders[val.ID] = val
	}
}
