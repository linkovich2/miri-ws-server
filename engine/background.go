package engine

import (
// "github.com/jonathonharrell/miri-ws-server/engine/util"
// "github.com/jonathonharrell/miri-ws-server/engine/util/loader"
)

type (
	Background struct {
		Name          string                  `json:"name"`
		ID            string                  `json:"id"`
		Prerequisites BackgroundPrerequisites `json:"prerequisites"`
		AllowAll      bool                    `json:"allow_all"`
	}

	BackgroundPrerequisites struct {
		Races            []string
		Genders          []string
		AestheticTraits  []string
		FunctionalTraits []string
	}
)

var backgrounds = make(map[string]Background)

func InitBackgrounds() {
	// arr := []AestheticTraitCategory{}
	// loader.Grab("aesthetic_traits.json", &arr)
	// for _, val := range arr {
	// 	aestheticTraitsCategorized[val.ID] = val
	// 	for _, t := range val.Traits {
	// 		t.Category = val.ID
	// 		aestheticTraits[t.ID] = t
	// 	}
	// }
}
