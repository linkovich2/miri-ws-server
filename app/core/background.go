package core

type (
	Background struct {
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
