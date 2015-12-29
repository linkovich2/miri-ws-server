package game

import (
	"github.com/jonathonharrell/miri-ws-server/app/core"
)

type (
	response struct {
		Room       *core.Room      `json:"room"`
		Messages   []string        `json:"messages"`
		Directions map[string]bool `json:"directions"`
		State      []string        `json:"state"`
	}

	miniResponse struct {
		Messages []string `json:"messages"`
	}
)
