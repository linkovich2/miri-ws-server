package engine

import (
	// "github.com/jonathonharrell/miri-ws-server/engine/util"
	"github.com/jonathonharrell/miri-ws-server/engine/util/loader"
)

type (
	Background struct {
		Name          string                  `json:"name"`
		ID            string                  `json:"id"`
		Prerequisites BackgroundPrerequisites `json:"prerequisites"`
		AllowAll      bool                    `json:"allow_all"`
		Description   string                  `json:"description"`
	}

	BackgroundShort struct {
		Name          string                  `json:"name"`
		ID            string                  `json:"id"`
		Prerequisites BackgroundPrerequisites `json:"prerequisites"`
		AllowAll      bool                    `json:"allow_all"`
		Description   string                  `json:"description"`
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
	arr := []Background{}
	loader.Grab("backgrounds.json", &arr)
	for _, val := range arr {
		backgrounds[val.ID] = val
	}
}

func (bg *Background) Shorten() BackgroundShort {
	short := BackgroundShort{
		Name:          bg.Name,
		ID:            bg.ID,
		Prerequisites: bg.Prerequisites,
		AllowAll:      bg.AllowAll,
		Description:   bg.Description,
	}

	return short
}
