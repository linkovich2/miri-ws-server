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

func (b *Background) IsAllowedForCharacter(character *Character) bool {
	if b.AllowAll {
		return true
	}

	matchesRace, matchesGender, matchesAestheticTraits, matchesFunctionalTraits := true, true, true, true

	if len(b.Prerequisites.Races) > 0 {
		matchesRace = false
		for _, r := range b.Prerequisites.Races {
			if character.Race == r {
				matchesRace = true
				break
			}
		}
	}

	if len(b.Prerequisites.Genders) > 0 {
		matchesGender = false
		for _, g := range b.Prerequisites.Genders {
			if character.Gender == g {
				matchesGender = true
				break
			}
		}
	}

	// @todo aesthetic traits
	// @todo functional traits

	return (matchesRace && matchesGender && matchesAestheticTraits && matchesFunctionalTraits)
}
