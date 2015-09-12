package engine

import (
	"github.com/jonathonharrell/miri-ws-server/engine/util/loader"
)

type Gender struct {
	Name string `json:"name"`
	ID   string `json:"id, omitempty"`
}

var genders []Gender

func InitGenders() {
	loader.Grab("genders.json", &genders)
}
