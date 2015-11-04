package engine

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
