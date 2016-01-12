package content

import (
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/app/core"
)

var entities map[string]core.ComponentBag

type readEntity struct {
	Name       string     `json:"name"`
	Behaviors  []string   `json:"behaviors"`
	Properties [][]string `json:"properties"`
	NotVisible bool       `json:"not_visible"`
}

func init() {
	if len(entities) <= 0 {
		files, err := AssetDir("json/entities")
		if err != nil {
			panic(err)
		}

		for _, f := range files {
			data, _ := Asset("json/entities/" + f)
			a := map[string]readEntity{}
			err := json.Unmarshal(data, &a)
			if err != nil {
				panic(err)
			}

			// make the real componentBag from the readEntities
		}
	}
}
