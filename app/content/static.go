package content

import (
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/app/core"
)

var aestheticTraits map[string]core.AestheticTraitCategory

func AestheticTraits() map[string]core.AestheticTraitCategory {
	if len(aestheticTraits) <= 0 {
		data, err := Asset("json/aesthetic_traits.json")
		if err != nil {
			panic(err)
		}

		a := map[string]core.AestheticTraitCategory{}
		err = json.Unmarshal(data, &a)
		if err != nil {
			panic(err)
		}
		aestheticTraits = a
	}

	return aestheticTraits
}
